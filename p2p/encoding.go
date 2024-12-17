package p2p

import (
	"encoding/gob"
	"io"

	"github.com/dev-hack95/localstorage/utilities/logs"
)

type GOBDecoder struct{}

type NPODecoder struct{}

// A type used to decode data packets
// Decode is function which takes input and output reader i.e io.Reader
type Decoder interface {
	Decode(io.Reader, *Message) error
}

func (dec GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

func (dec NPODecoder) Decode(r io.Reader, msg *Message) error {
	buffer := make([]byte, 2048)
	n, err := r.Read(buffer)
	if err != nil {
		logs.Error("Error: ", err.Error())
	}

	msg.Payload = buffer[:n]

	return nil
}
