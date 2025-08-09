package c9r

import (
	"encoding/xml"
	"log"

	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
)
type Action struct {
 	XMLName  xml.Name
 	InnerXML []byte  `xml:",innerxml"`
 	PearID   peer.ID `xml:"Root>Peer"`
}

type EndpointConfig struct {
	Name         string                 `yaml:"name"`
	SharedObject string                 `yaml:"shared_object"`
	Entrypoint   string                 `yaml:"entrypoint"`
	Extra        map[string]interface{} `yaml:",inline"` // all other fields will be included here
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
