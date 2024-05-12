package p2p

import (
	"net"
	"sync"

	"github.com/dev-hack95/localstorage/utilities/logs"
)

type TCPTransportsOpts struct {
	ListenAddr    string        // A network with a port i.e 127.0.0.1:8083
	HandshakeFunc HandshakeFunc // A ananomous function which accepts any type of data i.e type any = interface{}
	Decoder       Decoder
}

// TCPTransport is a struct that have information about sender and reciver
// listenaddress will holds address on which the transport
// will listen incoming connection
// listner holds the listner object and allows to accept incoming packets
// mu Mutexs are the advanced deadlocks that helps to get access over memory
// and helps to manage the memory over the transport period
// peer It maps associated network address i.e net.Addr with Peer interface
// net.Addr represents network address
type TCPTransport struct {
	TCPTransportsOpts
	listener net.Listener      // A listner accepts a packet request and return connection interface which will have access of read and write the streams
	mu       sync.RWMutex      // Mutex are often known as deadlocks and they are used to protect the data loss
	peers    map[net.Addr]Peer // peers is a dictonary that map as key: value pair
	// net.Addr define what type of packets we are sending and on which remote node location
	// i.e Network -> tcp and udp, String -> Network Node Info i.e 127.0.0.1:8083
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
func NewTCPTransport(opts TCPTransportsOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportsOpts: opts,
	}
}

// ListenAndAccept function will accept the type of network connection created
// Either tcp or udp
func (t *TCPTransport) ListenAndAccept() (err error) {
	// Listen will listen on given addrress via tcp protocol
	// Returns listener
	//	var err error
	// A methods which is assocaited with TCPTransport struct
	// We are passing network() which is either tcp or udp and string() which is address of new node
	t.listener, err = net.Listen("tcp", t.ListenAddr)
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

// The function uses for loop to create the connection between new devices
// It will accepts packets from connection
func (t *TCPTransport) startAcceptLoop() {
	// Until user revoke its connections the listener will listen the signals
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
	// Create a new tcp peer node

	_ = peer

	if err := t.HandshakeFunc(conn); err != nil {
		logs.Error("Tcp handshake error: ", err.Error())
		conn.Close()
		return
	}

	logs.Info("New endpoint connection: ", peer.conn.LocalAddr())

	msg := &Message{}

	for {
		// Decode message coming from connection
		if err := t.Decoder.Decode(conn, msg); err != nil {
			logs.Error("Decoder error: ", err.Error())
			continue
		}
	}

}
