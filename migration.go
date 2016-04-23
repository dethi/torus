package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

type Torrent struct {
	Timestamp   time.Time
	Hash        string
	Title       string
	Path        string
	RequestedBy string
}

func migration() {
	fmt.Println("Migrate database...")

	dbs := NewDatabase(*dbPath)
	defer dbs.Close()

	err := dbs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tBucket)

		var blob Torrent
		return b.ForEach(func(k, v []byte) error {
			json.Unmarshal(v, &blob)

			r := Record{
				BeginTime:       blob.Timestamp,
				EndDownloadTime: blob.Timestamp,
				EndTime:         blob.Timestamp,

				InfoHash:    blob.Hash,
				Name:        CleanName(blob.Title),
				FilePath:    blob.Path,
				RequestedBy: blob.RequestedBy,
			}

			buf, _ := json.Marshal(r)
			return b.Put(k, buf)
		})
	})
	if err != nil {
		fmt.Println("migration failed!")
	} else {
		dbs.ReindexRecords()
	}
}
