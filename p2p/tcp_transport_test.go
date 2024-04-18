package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing code snippet for TCPTransport
func TestTCPTransport(t *testing.T) {
	listenAddr := ":4344"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, tr.listenAddress, listenAddr)
	assert.Nil(t, tr.ListenAndAccept())
}
