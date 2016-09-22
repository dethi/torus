package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mailgun/mailgun-go"
)

type Mailer struct {
	mg     mailgun.Mailgun
	logger *log.Logger
}

type ValidUser func(email string) bool

func NewMailer(domain, apiKey, publicApiKey string) *Mailer {
	return &Mailer{
		mg:     mailgun.NewMailgun(domain, apiKey, publicApiKey),
		logger: log.New(os.Stderr, "Mailer: ", log.LstdFlags),
	}
}

func (s *Mailer) ReceiveMsg(pattern string, out chan<- Message,
	auth ValidUser) {

	http.HandleFunc(pattern, func(_ http.ResponseWriter, r *http.Request) {
		msg := readMsg(r)
		if auth(msg.From) {
			s.logger.Printf("accepted email from %v", msg.From)
			out <- msg
		} else {
			s.logger.Printf("rejected email from %v", msg.From)
		}
	})
}

func (s *Mailer) SendMsg(to, title, msg string) error {
	m := s.mg.NewMessage("Torrent Service <tr@dethi.fr>", title, msg, to)
	if _, _, err := s.mg.Send(m); err != nil {
		return fmt.Errorf("SendMessage: %v", err)
	}
	return nil
}

func (s *Mailer) NotifyUser(r *Record, to []string) {
	for _, email := range to {
		err := s.SendMsg(email, r.Name,
			"https://tr.dethi.fr/data/"+r.InfoHash+".tar")
		if err != nil {
			s.logger.Printf("failed to notify user %v: %v", email, err)
		}
	}
}

type Message struct {
	Timestamp time.Time
	To        string
	From      string
	Title     string
	Body      string
}

func readMsg(r *http.Request) Message {
	var msg Message

	timestamp := r.FormValue("timestamp")
	if ts, err := strconv.ParseInt(timestamp, 10, 64); err != nil {
		msg.Timestamp = time.Now()
	} else {
		msg.Timestamp = time.Unix(ts, 0)
	}
	msg.To = r.FormValue("recipient")
	msg.From = r.FormValue("sender")
	msg.Title = r.FormValue("subject")
	msg.Body = r.FormValue("stripped-text")
	return msg
}
