package torus

import "time"

type RequestState int

const (
	Waiting RequestState = iota
	Downloading
	Completed
)

type Request struct {
	UserId uint64       `json:"user_id"`
	URLs   []string     `json:"urls"`
	State  RequestState `json:"state"`

	RequestTime        time.Time     `json:"request_time"`
	DownloadDuration   time.Duration `json:"download_duration"`
	ProcessingDuration time.Duration `json:"processing_duration"`
}

type RequestService interface {
	NewRequest(userId uint64, URLs []string)
}
