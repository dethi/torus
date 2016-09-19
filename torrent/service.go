package torrent

import (
	"bytes"
	"log"
	"os"
	"path/filepath"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
	"github.com/pkg/errors"
)

type Torrent struct {
	Name     string
	Size     uint64
	InfoHash string
	Files    []string

	Payload []byte
}

func NewTorrent(payload []byte) (Torrent, error) {
	var t Torrent

	buf := bytes.NewBuffer(payload)
	mi, err := metainfo.Load(buf)
	if err != nil {
		return t, errors.Wrap(err, "load metainfo")
	}
	info := mi.UnmarshalInfo()

	t.Name = info.Name
	t.Size = uint64(info.TotalLength())
	t.InfoHash = convertInfoHash(mi.HashInfoBytes())
	t.Payload = payload

	for _, fileInfo := range info.UpvertedFiles() {
		if fileInfo.Path != nil {
			// Multiple files
			for _, path := range fileInfo.Path {
				t.Files = append(t.Files, filepath.Join(t.Name, path))
			}
		} else {
			// Simple file
			t.Files = append(t.Files, t.Name)
		}
	}

	return t, nil
}

func convertInfoHash(b [20]byte) string {
	slice := b[:]
	return string(slice)
}

type TorrentTask struct {
	Torrent Torrent
	Error   error
}

type Service struct {
	DataDir string

	token  chan struct{}
	logger *log.Logger
}

func NewService(token uint, dataDir string) *Service {
	ch := make(chan struct{}, token)
	for i := uint(0); i < token; i++ {
		ch <- struct{}{}
	}

	return &Service{
		DataDir: dataDir,
		token:   ch,
		logger:  log.New(os.Stderr, "TorrentService: ", log.LstdFlags),
	}
}

// Add torrents to the download queue and return a response channel.
func (ts *Service) Add(torrents ...Torrent) <-chan TorrentTask {
	ch := make(chan TorrentTask, len(torrents))
	go func() {
		<-ts.token
		// Don't forget to put the token back and to close the channel.
		defer func() {
			ts.token <- struct{}{}
			close(ch)
		}()

		if tasks, err := ts.download(torrents...); err != nil {
			ts.logger.Print(err)
		} else {
			for _, t := range tasks {
				ch <- t
			}
		}
	}()

	return ch
}

// Create a new client with a random free port and block until all downloads
// have finished. The client is closed at the end.
func (ts *Service) download(torrents ...Torrent) ([]TorrentTask, error) {
	client, err := torrent.NewClient(&torrent.Config{
		DataDir:    ts.DataDir,
		ListenAddr: ":0", // Pick a random free port
		NoUpload:   true,
		Seed:       false,
		Debug:      false,
		//NoDHT:      true,
		//DisableUTP: true,
		//ForceEncryption: true, // TODO: update dep to use this
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
	return tasks, nil
}
