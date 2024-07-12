package models

import "github.com/gliderlabs/ssh"

type Drop struct {
	ID               string
	Sender           ssh.Session
	Receiver         ssh.Session
	TransferredBytes int64
	WaitCh           chan struct{}
	DoneCh           chan struct{}
}
