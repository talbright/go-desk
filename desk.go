package desk

import (
	"log"
)

const (
	DeskLibVersion = "0.1"
	DeskApiVersion = "v2"
	DeskHost       = "desk.com"
	DeskUserAgent  = "go-desk/" + DeskLibVersion
)

func init() {
	log.SetPrefix("[desk] ")
	log.Println("init")
	log.Printf("Desk client library (%v) for desk.com API %v\n", DeskLibVersion, DeskApiVersion)
}
