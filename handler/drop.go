package handler

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
	"github.com/knid/drop/models"
	"github.com/knid/drop/utils"
)

type DropHandler struct {
	Drops map[string]*models.Drop
}

func (dh *DropHandler) Default(s ssh.Session) {
	if len(s.Command()) == 0 {
		drop := dh.NewDrop(s)
		dh.Drops[drop.ID] = drop

		log.Printf("%s: New drop created: %s\n", drop.Sender.RemoteAddr(), drop.ID)
		io.WriteString(s, fmt.Sprintf("New drop created: %s\n", drop.ID))
		io.WriteString(s, "Waiting for receiver...\n\n")

		for {
			<-drop.WaitCh

			io.WriteString(s, fmt.Sprintf("Start sending to %s\n", drop.Receiver.RemoteAddr()))
			log.Printf("Start sending to %s from %s\n", drop.Receiver.RemoteAddr(), drop.Sender.RemoteAddr())

			var err error
			drop.TransferredBytes, err = io.Copy(drop.Receiver, drop.Sender)
			if err != nil {
				log.Printf("Failed to send data to %s from %s\n", drop.Receiver.RemoteAddr(), drop.Sender.RemoteAddr())
				io.WriteString(s, fmt.Sprintf("Failed to send data to %s\n", drop.Receiver.RemoteAddr()))
				io.WriteString(s, "Waiting for new receiver...\n")
				continue
			}
			defer delete(dh.Drops, drop.ID)

			log.Printf("Successfully sent %d bytes to %s from %s\n", drop.TransferredBytes,
				drop.Receiver.RemoteAddr(), drop.Sender.RemoteAddr())
			io.WriteString(s, fmt.Sprintf("Successfully sent %d bytes to %s\n", drop.TransferredBytes, drop.Sender.RemoteAddr()))

			drop.DoneCh <- struct{}{}
			break
		}

	} else {
		id := s.Command()[0]
		drop, ok := dh.Drops[id]
		if !ok {
			io.WriteString(s, "Drop not found\n")
			return
		}

		drop.Receiver = s
		drop.WaitCh <- struct{}{}
		<-drop.DoneCh
	}
}

func (dh *DropHandler) NewDrop(sender ssh.Session) *models.Drop {
	return &models.Drop{
		ID:     utils.RandomString(8),
		Sender: sender,
		WaitCh: make(chan struct{}),
		DoneCh: make(chan struct{}),
	}
}
