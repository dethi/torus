package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/anacrolix/torrent/metainfo"
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
			// Wait for disk space (if needed)
			if err := waitDiskSpace(record); err != nil {
				d.logger.Print(err)
				continue
			}
			if err := readInfo(d.service.DataDir, &record); err != nil {
				d.logger.Print(err)
				continue
			}

			d.logger.Printf("start request: %v", record.InfoHash[:7])

			// Download
			t := torrent.Torrent{Body: record.torrent}
			ch := d.service.Add(t)
			for task := range ch {
				record.err = task.Error
			}

			record.EndDownloadTime = time.Now()
			d.out <- record

			d.logger.Printf("end request: %v", record.InfoHash[:7])
		}
	}()
}

func readInfo(dataDir string, r *Record) error {
	buf := bytes.NewBuffer(r.torrent)
	mi, err := metainfo.Load(buf)
	if err != nil {
		return errors.Wrap(err, "loading metainfo")
	}
	info := mi.UnmarshalInfo()
	r.Name = info.Name

	for _, fileinfo := range info.UpvertedFiles() {
		if fileinfo.Path == nil {
			// Simple file
			r.tFiles = append(r.tFiles, filepath.Join(dataDir, info.Name))
		} else {
			for _, path := range fileinfo.Path {
				r.tFiles = append(r.tFiles, filepath.Join(dataDir, info.Name, path))
			}
		}
	}

	return nil
}

func waitDiskSpace(r Record) error {
	tBuf := bytes.NewBuffer(r.torrent)
	metaInfo, err := metainfo.Load(tBuf)
	if err != nil {
		return fmt.Errorf("loading metainfo: %v", err)
	}

	info := metaInfo.UnmarshalInfo()
	size := uint64(info.TotalLength())
	fsStat, err := fs.GetFsStats(cfg.DataPath)
	if err != nil {
		return fmt.Errorf("stat fs: %v", err)
	}

	// Waiting for filesystem space
	for fsStat.Available < 2*size {
		time.Sleep(10 * time.Minute)

		// Update value
		if v, _ := fs.GetFsStats(cfg.DataPath); v != nil {
			fsStat = v
		}
	}

	return nil
}
