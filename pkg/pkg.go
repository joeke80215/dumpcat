package pkg

import "sync"

var (
	// PacketChan packet channel
	PacketChan chan Packet
)

func init() {
	PacketChan = make(chan Packet)
}

//
// ---base---
// *Time(*time.Time): package event time
//
// ---PackIPv4 package information---
// *Proto(string): portocal
// *Src(string): source <ip:port>
// *Dst(string): destination  <ip:port>
// *Lenght(uint16): package lenght
//
// refernece: https://en.wikipedia.org/wiki/Transmission_Control_Protocol
// ---PackTCP tcp package information---
// *Proto(string): portocal
// *IPv4(PackIPv4): ipv4 package info
// *Src(string): source port
// *Dst(string): destination port
// *NS(bool): ECN-nonce - concealment protection
// *CWR(bool): Congestion Window Reduced (CWR) flag is set by the sending host to indicate that
//  it received a TCP segment with the ECE flag set and had responded in congestion control
//  mechanism (added to header by RFC 3168).
// *ECE(bool): ECN-Echo has a dual role, depending on the value of the SYN flag
// *URG(bool):  indicates that the Urgent pointer field is significant
// *Urgent: if the URG flag is set, then this 16-bit field is an offset from the sequence number
//  indicating the last urgent data byte
// *ACK(bool): indicates that the Acknowledgment field is significant. All packets after the
//  initial SYN packet sent by the client should have this flag set.
// *Ack(uint32): if the ACK flag is set then the value of this field is the next sequence number
//  that the sender of the ACK is expecting. This acknowledges receipt of all prior bytes (if any).
//  The first ACK sent by each end acknowledges the other end's initial sequence number itself, but no data.
// *PSH(bool): Push function. Asks to push the buffered data to the receiving application.
// *RST(bool): Reset the connection
// *SYN(bool): Synchronize sequence numbers. Only the first packet sent from each end should have this flag set.
//  Some other flags and fields change meaning based on this flag, and some are only valid when it is set,
//  and others when it is clear.
// *SEQ(uint32):
//  Has a dual role:
//  If the SYN flag is set (1), then this is the initial sequence number.
//  The sequence number of the actual first data byte and the acknowledged number in the corresponding
//  ACK are then this sequence number plus 1.
//  If the SYN flag is clear (0), then this is the accumulated sequence number of the first data byte of this segment
//  for the current session.
// *FIN(bool): Last packet from sender.
// *Window(uint16): The size of the receive window, which specifies the number of window size units (by default, bytes)
//  (beyond the segment identified by the sequence number in the acknowledgment field) that the sender of this segment is
//  currently willing to receive
// *Checksum(uint16): The 16-bit checksum field is used for error-checking of the header, the Payload and a Pseudo-Header.
//  The Pseudo-Header consists of the Source IP Address, the Destination IP Address, the protocol number for the TCP-Protocol (0x0006)
//  and the length of the TCP-Headers including Payload (in Bytes).
//

// Packet interface
type Packet interface {
	// Set packet parameter
	Set(string, interface{})

	// Get key value
	Get(string) interface{}

	// GetData get data map[string]interface{}
	GetData() map[string]interface{}
}

type packet struct {
	data  map[string]interface{}
	rwMut sync.RWMutex
}

// NewPacket new packet
func NewPacket() Packet {
	return &packet{
		data: make(map[string]interface{}),
	}
}

// Set package value
func (p *packet) Set(key string, value interface{}) {
	p.rwMut.Lock()
	defer p.rwMut.Unlock()
	p.data[key] = value
}

// Get package value
func (p *packet) Get(key string) interface{} {
	p.rwMut.RLock()
	defer p.rwMut.RUnlock()
	return p.data[key]
}

// GetData get data map[string]interface{}
func (p *packet) GetData() map[string]interface{} {
	p.rwMut.RLock()
	defer p.rwMut.RUnlock()
	return p.data
}
