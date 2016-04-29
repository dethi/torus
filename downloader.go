package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	tr "github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/metainfo"
)

type Downloader struct {
	in  <-chan Record
	out chan<- Record
	dir string

	logger *log.Logger
	client *tr.Client
}

func NewDownloader(in <-chan Record, out chan<- Record,
	dir string) *Downloader {

	return &Downloader{
		in:  in,
		out: out,
		dir: dir,

		logger: log.New(os.Stderr, "Downloader: ", log.LstdFlags),
	}
}

func (s *Downloader) Start(addr string) error {
	cl, err := tr.NewClient(&tr.Config{
		DataDir:    s.dir,
		ListenAddr: addr,
		NoUpload:   true,
		Seed:       false,
		Debug:      false,
	})
	if err != nil {
		return fmt.Errorf("creating client: %v", err)
	}
	s.client = cl

	s.accept()
	return nil
}

func (s *Downloader) accept() {
	go func() {
		for record := range s.in {
			s.logger.Printf("start request: %v", record.InfoHash[:7])
			err := s.download(&record)
			if err != nil {
				s.logger.Printf("error while downloding: %v: %v",
					record.InfoHash[:7], err)
				record.err = err
			}

			s.out <- record
			s.logger.Printf("end request: %v", record.InfoHash[:7])
		}

		defer s.client.Close()
	}()
}

func (s *Downloader) download(record *Record) error {
	tBuf := bytes.NewBuffer(record.torrent)
	metaInfo, err := metainfo.Load(tBuf)
	if err != nil {
		return fmt.Errorf("loading metainfo: %v", err)
	}

	t, err := s.client.AddTorrent(metaInfo)
	if err != nil {
		return fmt.Errorf("adding metainfo: %v", err)
	}

	<-t.GotInfo()
	record.Name = t.Name()

	t.DownloadAll()
	s.client.WaitAll()
	record.EndDownloadTime = time.Now()

	for _, f := range t.Files() {
		path := filepath.Join(s.dir, f.Path())
		record.tFiles = append(record.tFiles, path)
	}
	t.Drop()
	return nil
}
