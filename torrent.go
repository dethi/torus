package torus

import (
	"bytes"
	"encoding/hex"
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

// NewTorrent creates a new torrent from a payload. It returns an error if
// the payload is invalid.
func NewTorrent(payload []byte) (Torrent, error) {
	var t Torrent

	buf := bytes.NewBuffer(payload)
	mi, err := metainfo.Load(buf)
	if err != nil {
		return t, errors.Wrap(err, "load metainfo")
	}
	info, err := mi.UnmarshalInfo()
	if err != nil {
		return t, errors.Wrap(err, "unmarshal info")
	}

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
	return hex.EncodeToString(b[:])
}
