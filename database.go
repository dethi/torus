package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/boltdb/bolt"
)

type Database struct {
	db     *bolt.DB
	index  *algoliasearch.Index
	logger *log.Logger

	requests map[string][]string
	*sync.Mutex
}

func NewDatabase(dbPath string) *Database {
	db, err := bolt.Open(dbPath, 0600, &bolt.Options{
		Timeout: 1 * time.Second,
	})
	if err != nil {
		log.Fatalln("opening database: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(tBucket); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatalln("creating bucket: %v", err)
	}

	client := algoliasearch.NewClient(*algoliaAppId, *algoliaApiKey)
	return &Database{
		db:       db,
		index:    client.InitIndex(*algoliaIndex),
		logger:   log.New(os.Stderr, "Database: ", log.LstdFlags),
		requests: make(map[string][]string),
		Mutex:    &sync.Mutex{},
	}
}

func (s *Database) Close() error {
	return s.db.Close()
}

func (s *Database) GetRecord(infoHash string) *Record {
	var buf []byte
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tBucket)
		if buf = b.Get([]byte(infoHash)); buf == nil {
			return fmt.Errorf("key %v not found", infoHash[:7])
		}
		return nil
	})
	if err != nil {
		return nil
	}

	var r Record
	if err := json.Unmarshal(buf, &r); err != nil {
		return nil
	}
	return &r
}

func (s *Database) ViewRecords() []Record {
	var records []Record

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(tBucket)

		return b.ForEach(func(k, v []byte) error {
			var blob Record
			if err := json.Unmarshal(v, &blob); err != nil {
				return fmt.Errorf("unmarshal %v: %v", k[:7], err)
			}
			records = append(records, blob)
			return nil
		})
	})
	if err != nil {
		s.logger.Printf("view records: %v", err)
	}

	return records
}

func (s *Database) PutRecord(r Record) {
	buf, err := json.Marshal(r)
	if err != nil {
		s.logger.Printf("json marshal %v: %v", r.InfoHash[:7], err)
		return
	}

	err = s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(tBucket)
		return b.Put([]byte(r.InfoHash), buf)
	})
	if err != nil {
		s.logger.Printf("put %v: %v", r.InfoHash[:7], err)
		return
	}

	if _, err = s.index.AddObject(r); err != nil {
		s.logger.Printf("indexing %v: %v", r.InfoHash[:7], err)
	}
}

func (s *Database) ReindexRecords() {
	if _, err := s.index.Clear(); err != nil {
		s.logger.Printf("clearing index: %v", err)
	}
	for _, r := range s.ViewRecords() {
		if _, err := s.index.AddObject(r); err != nil {
			s.logger.Printf("indexing %v: %v", r.InfoHash[:7], err)
		}
	}
}

func (s *Database) GetRequest(infoHash string) []string {
	s.Lock()
	defer s.Unlock()
	return s.requests[infoHash]
}

func (s *Database) PutRequest(infoHash string, email string) {
	s.Lock()
	defer s.Unlock()
	s.requests[infoHash] = append(s.requests[infoHash], email)
}

func (s *Database) DeleteRequest(infoHash string) {
	s.Lock()
	defer s.Unlock()
	delete(s.requests, infoHash)
}
