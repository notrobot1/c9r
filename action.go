package c9r

import (
	"encoding/xml"
	"log"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)

type Action struct {
	XMLName xml.Name
	Action  string
	PearID  peer.ID
}

func MarshalXML(v interface{}, s network.Stream) {
	xmlData, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		ErrorXML(err, s)
		return
	}
	// Sending the response back through the stream
	s.Write(xmlData)

	s.Close()
}

func ErrorXML(err error, s network.Stream) {

	log.Printf("%v", err)

	type Response struct {
		XMLName xml.Name `xml:"Response"`
		Status  int      `xml:"Status"`
	}

	response := Response{
		Status: 400,
	}

	xmlData, _ := xml.MarshalIndent(response, "", "  ")

	// Sending the response back through the stream
	_, _ = s.Write(xmlData)

	s.Close()
}
