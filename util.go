package main

import (
	"crypto/sha1"
	"fmt"

	"github.com/zeebo/bencode"
)

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func InfoHash(data []byte) string {
	var torrent struct {
		Info interface{} `bencode:"info"`
	}

	if err := bencode.DecodeBytes(data, &torrent); err != nil {
		return ""
	}

	binfo, err := bencode.EncodeBytes(torrent.Info)
	if err != nil {
		return ""
	}
	hash := sha1.Sum(binfo)
	return fmt.Sprintf("%x", hash)
}
