package torus

import (
	"bytes"
	"fmt"
	"path/filepath"

	"github.com/anacrolix/torrent/metainfo"
	"github.com/dethi/torus/util"
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

	t.Name, _ = util.SplitFilename(util.CleanName(info.Name))
	t.Size = uint64(info.TotalLength())
	t.InfoHash = convertInfoHash(mi.HashInfoBytes())
	t.Payload = payload

	for _, fileInfo := range info.UpvertedFiles() {
		if fileInfo.Path != nil {
			// Multiple files
			for _, path := range fileInfo.Path {
				t.Files = append(t.Files, filepath.Join(info.Name, path))
			}
		} else {
			// Simple file
			t.Files = append(t.Files, info.Name)
		}
	}

	return t, nil
}

func convertInfoHash(b [20]byte) string {
	slice := b[:]
	return fmt.Sprintf("%x", slice)
}
