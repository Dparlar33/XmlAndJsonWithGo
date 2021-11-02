package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	NewFile, err := os.Open("deneme.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer NewFile.Close()
	data_file, _ := ioutil.ReadAll(NewFile)

	var x Message
	xml.Unmarshal(data_file, &x)
	fmt.Printf(print(x))
}

type Message struct {
	XMLname      xml.Name `xml:"note"`
	SenderName   string   `xml:"to"`
	ReceiverName string   `xml:"from"`
	Subject      string   `xml:"heading"`
	Body         string   `xml:"body"`
}

func print(p Message) string {
	return fmt.Sprintf("Sender: %s, Receiver :%s, Subject : %s,Message: %s", p.SenderName, p.ReceiverName, p.Subject, p.Body)
}
