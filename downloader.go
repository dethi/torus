package main

import (
	"log"
	"os"
	"time"

	"github.com/dethi/goutil/fs"
	"github.com/dethi/torus/torrent"
	"github.com/pkg/errors"
)

type Downloader struct {
	in     <-chan Record
	out    chan<- Record
	logger *log.Logger

	service *torrent.Service
}

func NewDownloader(in <-chan Record, out chan<- Record,
	dir string) *Downloader {

	return &Downloader{
		in:     in,
		out:    out,
		logger: log.New(os.Stderr, "Downloader: ", log.LstdFlags),

		service: torrent.NewService(cfg.DownloadToken, dir),
	}
}

func (d *Downloader) Start() {
	go func() {
		for record := range d.in {
			if err := waitDiskSpace(record.Size); err != nil {
				d.logger.Print(err)
				continue
			}
			d.logger.Printf("start request: %v", record.InfoHash[:7])

			// Download
			ch := d.service.Add(record.Torrent)
			for task := range ch {
				record.err = task.Error
			}

			record.EndDownloadTime = time.Now()
			d.out <- record

			d.logger.Printf("end request: %v", record.InfoHash[:7])
		}
	}()
}

func waitDiskSpace(size uint64) error {
	fsStat, err := fs.GetFsStats(cfg.DataPath)
	if err != nil {
		return errors.Wrap(err, "wait disk space")
	}

	// Waiting for filesystem space
	for fsStat.Available < 2*size {
		time.Sleep(10 * time.Minute)
		if v, _ := fs.GetFsStats(cfg.DataPath); v != nil {
			fsStat = v
		}
	}

	return nil
}
