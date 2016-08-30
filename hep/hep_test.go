package hep

import (
	"testing"
)

/**
* Naive HEP3 message parse test...
 */
func TestBasicParse(t *testing.T) {
	t.Log("Test basic parse")

	udpPacket := []byte{
		0x48, 0x45, 0x50, 0x33, // HepID
		0x00, 0x71, 0x00, 0x00, 0x00, 0x01, 0x00, 0x07, 0x02,
		0x00, 0x00, 0x00, 0x02, 0x00, 0x07, 0x11, // protocol ID = 17 (UDP)
		0x00, 0x00, 0x00, 0x03, 0x00, 0x0a, 0xd4, 0xca, 0x00, 0x01, // IPv4 source address = 212.202.0.1
		0x00, 0x00, 0x00, 0x04, 0x00, 0x0a, 0x52, 0x74, 0x00, 0xd3, // IPv4 destination address = 82.116.0.211
		0x00, 0x00, 0x00, 0x07, 0x00, 0x08, 0x2e, 0xea, // source port = 12010
		0x00, 0x00, 0x00, 0x08, 0x00, 0x08, 0x13, 0xc4, // destination port = 5060
		0x00, 0x00, 0x00, 0x09, 0x00, 0x0a, 0x4e, 0x49, 0x82, 0xcb, // seconds timestamp 1313440459 = Mon Aug 15 22:34:19 2011
		0x00, 0x00, 0x00, 0x0a, 0x00, 0x0a, 0x00, 0x01, 0xd4, 0xc0, // micro-seconds timestamp offset 120000 = 0.12 seconds
		0x00, 0x00, 0x00, 0x0b, 0x00, 0x07, 0x01, // 01 – SIP
		0x00, 0x00, 0x00, 0x0c, 0x00, 0x0a, 0x00, 0x00, 0x00, 0xE4, // capture ID (228)
		0x00, 0x00, 0x00, 0x0f, 0x00, 0x14, 0x49, 0x4e, 0x56, 0x49, 0x54, 0x45, 0x20, 0x73, 0x69, 0x70, 0x3a, 0x62, 0x6f, 0x62, // SIP payload “INVITE sip:bob” (shortened)
	}

	t.Log("Parsing sample HEP3 message from spec...")

	msg, err := NewHepMsg(udpPacket)
	if err != nil {
		t.Log("Parse failed")
		t.Fatal(err.Error())
	}

	if msg.IP4SourceAddress != "212.202.0.1" {
		t.Fatalf("HepMsg.Parse: Expected source address to be 212.202.0.1, but it was %s", msg.IP4SourceAddress)
	}
	if msg.IP4DestinationAddress != "82.116.0.211" {
		t.Fatalf("HepMsg.Parse: Expected source address to be 212.202.0.1, but it was %s", msg.IP4DestinationAddress)
	}
	if msg.SourcePort != 12010 {
		t.Fatalf("HepMsg.Parse: Expected source port to be 12010, but it was %d", msg.SourcePort)
	}
	if msg.DestinationPort != 5060 {
		t.Fatalf("HepMsg.Parse: Expected source port to be 5060, but it was %d", msg.DestinationPort)
	}
}
