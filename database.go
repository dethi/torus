package torus

type DatabaseService interface {
	AddUser(u User) error
	AddTransaction(t Transaction) error
	AddRequest(r Request) error
	AddTorrent(t Torrent) error
	AddRecord(r Record) error
}
