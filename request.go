package torus

import "time"

type RequestState int

const (
	Waiting RequestState = iota
	Downloading
	Completed
)

type Request struct {
	UserId uint64
	URLs   []string
	State  RequestState

	RequestTime        time.Time
	DownloadDuration   time.Duration
	ProcessingDuration time.Duration
}

type RequestService interface {
	NewRequest(userId uint64, URLs []string)
}
