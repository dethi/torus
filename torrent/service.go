package torrent

import (
	"bytes"
	"log"
	"os"
	"strconv"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/dethi/torus"
	"github.com/pkg/errors"
)

type TorrentTask struct {
	Torrent torus.Torrent
	Error   error
}

type Service struct {
	DataDir string

	token  chan uint
	logger *log.Logger
}

// NewService initializes a new service using dataDir as its storage
// directory. The token limits the number of requests completed at a given
// time. Behind the scene, it uses a new torrent client for each request. Each
// torrent client have a dedicated port equal to port+current_token.
func NewService(port uint, token uint, dataDir string) *Service {
	ch := make(chan uint, token)
	for i := uint(0); i < token; i++ {
		ch <- port + i
	}

	return &Service{
		DataDir: dataDir,
		token:   ch,
		logger:  log.New(os.Stderr, "TorrentService: ", log.LstdFlags),
	}
}

// Add adds torrents to the download queue and return a response channel.
func (ts *Service) Add(torrents ...torus.Torrent) <-chan TorrentTask {
	ch := make(chan TorrentTask, len(torrents))
	go func() {
		port := <-ts.token
		// Don't forget to put the token back and to close the channel.
		defer func() {
			ts.token <- port
			close(ch)
		}()

		if tasks, err := ts.download(port, torrents...); err != nil {
			ts.logger.Print(err)
		} else {
			for _, t := range tasks {
				ch <- t
			}
		}
	}()

	return ch
}

// download creates a new torrent client with port and block until all
// downloads have finished. The client is closed at the end.
func (ts *Service) download(port uint, torrents ...torus.Torrent)
	([]TorrentTask, error) {

	client, err := torrent.NewClient(&torrent.Config{
		DataDir:    ts.DataDir,
		ListenAddr: ":" + strconv.Itoa(int(port)),
		NoUpload:   true,
		Seed:       false,
		Debug:      false,
		//NoDHT:      true,
		//DisableUTP: true,
		//ForceEncryption: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "initialize client")
	}
	defer client.Close()

	var tasks = make([]TorrentTask, len(torrents))
	for i, t := range torrents {
		tasks[i].Torrent = t

		buf := bytes.NewBuffer(t.Payload)
		mi, err := metainfo.Load(buf)
		if err != nil {
			// This should never happened because metainfo are pre-loaded
			// when added to the service.
			tasks[i].Error = errors.Wrap(err, "load metainfo")
			continue
		}

		tr, err := client.AddTorrent(mi)
		if err != nil {
			tasks[i].Error = errors.Wrap(err, "add torrent")
			continue
		}

		<-tr.GotInfo()
		tr.DownloadAll()
	}

	client.WaitAll()

	// Remove all torrents from the client to avoid write after close panic.
	for _, tr := range client.Torrents() {
		tr.Drop()
	}
	return tasks, nil
}
