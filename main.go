package main

import (
	"archive/tar"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strconv"
	"syscall"
	"time"
)

var (
	mgDomain    = flag.String("mgDomain", "", "mailgun domain")
	mgKey       = flag.String("mgKey", "", "mailgun key")
	mgPublicKey = flag.String("mgPublicKey", "", "mailgun public key")

	algoliaAppId     = flag.String("algoliaAppId", "", "algolia app id")
	algoliaApiKey    = flag.String("algoliaApiKey", "", "algolia api key")
	algoliaSearchKey = flag.String("algoliaSearchKey", "", "algolia search key")
	algoliaIndex     = flag.String("algoliaIndex", "", "algolia index name")

	dbPath   = flag.String("dbPath", "/static/torrents.db", "path to database")
	dataPath = flag.String("dataPath", "/static/data", "path to data")

	migrate    = flag.Bool("migrate", false, "migrate database")
	activeJobs = flag.Int("activeJobs", 3, "number of active torrents")
)

const (
	findUrl   = `((?:https?:\/{2})(?:[-\w]+\.)+(?:\w+)(?:\/[-\w]+)*(?:\.[-\w]+)?)`
	cleanName = `(?i)((\[ *)?[a-z]+.cpasbien.[a-z]+( *\])?)|(web(-?dl)?)|(xvid)`
)

const listHtml = `Nothing for now sorry`

var regexUrl = regexp.MustCompile(findUrl)
var regexClean = regexp.MustCompile(cleanName)

type Record struct {
	BeginTime       time.Time
	EndDownloadTime time.Time
	EndTime         time.Time

	InfoHash    string
	Name        string
	FilePath    string
	RequestedBy string

	err error

	torrent []byte
	tFiles  []string
}

func dispatcher(mailer *MailService, newMail <-chan Message, newJob chan<- Record,
	endJob <-chan Record, quit <-chan os.Signal) {

	db := NewDatabase(*dbPath)
	defer db.Close()

	for {
		select {
		case <-quit:
			return
		case msg := <-newMail:
			for _, url := range regexUrl.FindAllString(msg.Body, -1) {
				torrent, err := FetchTorrent(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				infoHash := InfoHash(torrent)
				if infoHash == "" {
					// TODO log
					fmt.Println("skip because empty hash")
					continue
				}

				// The request exist and is finished, just send an email
				// and leave.
				if r := db.GetRecord(infoHash); r != nil {
					mailer.NotifyUser(r, []string{msg.From})
					continue
				}

				// The request is currently processing, just add the email
				// and leave.
				exist := (db.GetRequest(infoHash) != nil)
				db.PutRequest(infoHash, msg.From)
				if exist {
					fmt.Println("wait, record is processing")
					continue
				}

				r := Record{
					BeginTime:   time.Now(),
					InfoHash:    infoHash,
					RequestedBy: msg.From,
					torrent:     torrent,
				}
				newJob <- r
			}
		case r := <-endJob:
			if r.err != nil {
				// TODO: error handling
				fmt.Println(r.err)
				continue
			}

			r.Name = CleanName(r.Name)
			r.FilePath = filepath.Join(*dataPath, r.InfoHash+".tar")
			if err := createTarball(r); err != nil {
				// TODO: error handling
				fmt.Println(err)
				continue
			}

			// Clean files
			for _, path := range r.tFiles {
				os.Remove(path)
			}

			r.EndTime = time.Now()
			db.PutRecord(r)
			mailer.NotifyUser(&r, db.GetRequest(r.InfoHash))
			db.DeleteRequest(r.InfoHash)
		}
	}
}

func createTarball(r Record) error {
	tarball, err := os.Create(r.FilePath)
	if err != nil {
		return fmt.Errorf("createTarball: %v", err)
	}
	defer tarball.Close()

	tw := tar.NewWriter(tarball)
	defer tw.Close()

	for _, path := range r.tFiles {
		err := func() error {
			f, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("open file: %v: %v", path, err)
			}
			defer f.Close()

			stat, err := f.Stat()
			if err != nil {
				return fmt.Errorf("stat file: %v: %v", path, err)
			}

			hdr := &tar.Header{
				Name: filepath.Base(CleanName(path)),
				Mode: 0644,
				Size: stat.Size(),
			}
			if err := tw.WriteHeader(hdr); err != nil {
				return fmt.Errorf("write header: %v: %v", path, err)
			}

			if _, err := io.Copy(tw, f); err != nil {
				return fmt.Errorf("write file: %v: %v", path, err)
			}
			return nil
		}()
		if err != nil {
			return fmt.Errorf("create tarball %v: %v", r.InfoHash[:7], err)
		}
	}
	return nil
}

func startService(activeJobs int) {
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

	mailer := NewMailService(*mgDomain, *mgKey, *mgPublicKey)
	mailer.ReceiveMsg("/mg-mail", newMail, func(_ string) bool {
		return true
	})

	var downloader []*DownloaderService
	for i := 0; i < activeJobs; i++ {
		downloader = append(downloader,
			NewDownloaderService(newJob, endJob, tmpDir))
	}
	for i, srv := range downloader {
		listenAddr := ":" + strconv.Itoa(50007+i)
		srv.Start(listenAddr)
	}

	http.HandleFunc("/list", listHandler)
	http.Handle("/data/", http.StripPrefix("/data/",
		http.FileServer(http.Dir(*dataPath))))
	go http.ListenAndServe(":80", nil)
	dispatcher(mailer, newMail, newJob, endJob, quit)
	fmt.Println("Gracefully stop service...")
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, listHtml)
}

func main() {
	flag.Parse()

	if *migrate {
		migration()
	} else {
		startService(*activeJobs)
	}
}
