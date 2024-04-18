package p2p

import (
	"net"
	"sync"

	"github.com/dev-hack95/localstorage/utilities/logs"
)

// TCPTransport is a struct that have information about sender and reciver
// listenaddress will holds address on which the transport
// will listen incomming connection
// listner holds the listner object and allows to accept incoming packets
// mu Mutexs are the advanced deadlocks that helps to get access over memory
// and helps to manage the memory over the transport period
// peer It maps associated network address i.e net.Addr with Peer interface
// net.Addr represents network address
type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakefunc HandshakeFunc
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

// TCP Peer represents remote node over TCP established connection
// if we dial and retrive connection --> outbound == true and inbound == false
// if we accept and retrive connection --> outbound == false and inbound == true
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

// NewTCP Peer function accepts the connection and check that outbound is true or not
// When you are sending request away from server its called outbound and outbound == true
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// NewTCPTransport function accepts addres and retuns TCPTransport
// struct thatt will have acces to Peer and Listner
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		handshakefunc: NPOHandShakeFunc,
		listenAddress: listenAddr,
	}
}

// ListenAndAccept method listen to a given address
func (t *TCPTransport) ListenAndAccept() error {
	// Listen will listen on given addrress via tcp protocol
	// Returns listener
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	// If error occured it will return the error
	if err != nil {
		logs.Error("Error: ", err.Error())
		return err
	}
	_ = t.listener

	// Using Goroutines to start accepts the messages from sender
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		// It will accept message from listener
		conn, err := t.listener.Accept()
		// Error is managed here
		if err != nil {
			logs.Error("Error: ", err.Error())
		}
		// Using Goroutines to handle connction
		go t.handleConnection(conn)
	}
}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.handshakefunc(conn); err != nil {
		logs.Error("Error: ", err.Error())
	}
	logs.Info("New endpoint connection: ", peer.conn.LocalAddr())
}
