package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing code snippet for TCPTransport
func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportsOpts{
		ListenAddr:    ":3000",
		HandshakeFunc: NPOHandShakeFunc,
		Decoder:       GOBDecoder{},
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tr.ListenAddr, tcpOpts)
	assert.Nil(t, tr.ListenAndAccept())
}
