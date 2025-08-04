package examples

import (
	"fmt"
	"log"

	"github.com/emersion/go-message/mail"
	bp "github.com/sarff/bounce-parser"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

func main() {
	// Auth
	c, err := client.DialTLS("imap.example.com:993", nil) // imap.gmail.com:993
	if err != nil {
		log.Fatal("IMAP connection error:", err)
	}
	defer c.Logout()

	if err := c.Login("your-email@example.com", "password"); err != nil {
		log.Fatal("IMAP login error:", err)
	}

	// Folder INBOX
	mbox, err := c.Select("INBOX", false)
	if err != nil {
		log.Fatal("Select INBOX error:", err)
	}

	if mbox.Messages == 0 {
		log.Println("No messages in INBOX")
		return
	}

	// Last 5 letters
	from := uint32(1)
	to := mbox.Messages
	if mbox.Messages > 5 {
		from = to - 4
	}
	seqSet := new(imap.SeqSet)
	seqSet.AddRange(from, to)

	section := &imap.BodySectionName{}
	items := []imap.FetchItem{imap.FetchEnvelope, section.FetchItem()}
	messages := make(chan *imap.Message, 5)
	go func() {
		if err := c.Fetch(seqSet, items, messages); err != nil {
			log.Fatal("Fetch error:", err)
		}
	}()

	// Обробка кожного повідомлення
	for msg := range messages {
		if msg == nil {
			continue
		}

		subject := msg.Envelope.Subject
		r := msg.GetBody(section)
		if r == nil {
			log.Println("Empty body for message:", subject)
			continue
		}

		mr, err := mail.CreateReader(r)
		if err != nil {
			log.Println("Error creating mail reader:", err)
			continue
		}

		// Bounce check
		b, err := bp.Parse(subject, mr)
		if err != nil {
			log.Println("Error parsing mail:", err)
			continue
		}
		if b.Type != bp.BounceNone {
			fmt.Printf("Bounce detected: [%s] %s (Subject: %s) Mailbox: %s\n", b.Type, b.Reason, b.Subject, b.Mailbox)
		}
	}
}
