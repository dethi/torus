package torus

type Record struct {
	ID       uint64    `json:"id"`
	Request  Request   `json:"request"`
	Torrents []Torrent `json:"torrents"`
}
