package main

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
)

func migration() {
	fmt.Println("Migrate database...")

	dbs := NewDatabase(*dbPath)
	defer dbs.Close()

	err := dbs.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tBucket)

		var r Record
		return b.ForEach(func(k, v []byte) error {
			json.Unmarshal(v, &r)
			r.Name = CleanName(r.Name)
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
