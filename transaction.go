package torus

import "time"

type TransactionType int

const (
	AddTorrent TransactionType = iota
	DownloadTorrent
)

type Transaction struct {
	Type   TransactionType
	Cost   uint
	Tags   []string
	UserId uint64

	Datetime time.Time
}

type TransactionService interface {
	Add(t Transaction) error
	View(begin, end time.Time) ([]Transaction, error)
}
