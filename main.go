package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/abbot/go-http-auth"
	"github.com/carlescere/scheduler"
	"github.com/dethi/torus/torrent"
	"github.com/dethi/torus/util"
)

type Record struct {
	BeginTime       time.Time
	EndDownloadTime time.Time
	EndTime         time.Time
	RequestedBy     string
	Pathname        string

	torrent.Torrent

	err error
}

var db *Database

func dispatcher(mailer *Mailer, newMail <-chan Message, newJob chan<- Record,
	endJob <-chan Record, quit <-chan os.Signal) {

	db = NewDatabase(cfg.DatabasePath)
	defer db.Close()

	scheduler.Every(1).Hours().Run(func() {
		db.DeleteRecords(2 * 24 * time.Hour)
	})

	for {
		select {
		case <-quit:
			return
		case msg := <-newMail:
			for _, url := range regexUrl.FindAllString(msg.Body, -1) {
				payload, err := FetchTorrent(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				torrent, err := torrent.NewTorrent(payload)
				if err != nil {
					fmt.Println(err)
					continue
				}

				// The request exist and is finished, just send an email
				// and leave.
				if r := db.GetRecord(torrent.InfoHash); r != nil {
					mailer.NotifyUser(r, []string{msg.From})
					continue
				}

				// The request is currently processing, just add the email
				// and leave.
				exist := (db.GetRequest(torrent.InfoHash) != nil)
				db.PutRequest(torrent.InfoHash, msg.From)
				if exist {
					fmt.Println("wait, record is processing")
					continue
				}

				r := Record{
					BeginTime:   time.Now(),
					RequestedBy: msg.From,
					Torrent:     torrent,
				}
				newJob <- r
			}
		case r := <-endJob:
			if r.err != nil {
				fmt.Println(r.err)
				continue
			}

			r.Name = util.CleanName(r.Name)
			r.Pathname = filepath.Join(cfg.DataPath, r.InfoHash+".tar")
			files := util.AddPathPrefix(cfg.DataPath, r.Files...)

			if err := util.CreateTarball(r.Pathname, files...); err != nil {
				fmt.Println(err)
				continue
			}

			// Clean files
			for _, path := range files {
				os.Remove(path)
			}

			r.EndTime = time.Now()
			db.PutRecord(r)
			mailer.NotifyUser(&r, db.GetRequest(r.InfoHash))
			db.DeleteRequest(r.InfoHash)
		}
	}
}

func startService() {
	fmt.Println("Start service...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill, syscall.SIGTERM)

	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatalf("creating tempory directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	newMail := make(chan Message, 100)
	newJob := make(chan Record, 100)
	endJob := make(chan Record, 100)

	mailer := NewMailer(cfg.Mailgun.Domain, cfg.Mailgun.SecretKey, cfg.Mailgun.PublicKey)
	mailer.ReceiveMsg("/mg-mail", newMail, func(_ string) bool {
		return true
	})

	var downloader = NewDownloader(newJob, endJob, tmpDir)
	downloader.Start()

	authenticator := auth.NewBasicAuthenticator("tr.dethi.fr",
		auth.HtpasswdFileProvider(cfg.HtpasswdPath))

	http.HandleFunc("/", auth.JustCheck(authenticator, listView))
	http.Handle("/data/", http.StripPrefix("/data/",
		http.FileServer(http.Dir(cfg.DataPath))))
	go http.ListenAndServe(":8000", nil)
	dispatcher(mailer, newMail, newJob, endJob, quit)
	fmt.Println("Gracefully stop service...")
}

func main() {
	flag.Parse()
	if *versionFlag {
		fmt.Println("rev", version)
		return
	}

	if err := LoadConfig(*configPath); err != nil {
		log.Fatal(err)
	}

	startService()
}
