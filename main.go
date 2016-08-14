package main

import (
	"log"

	"github.com/yogendrars/yogi3100/cloud-torrent/server"
	"github.com/yogendrars/opts"
)

var VERSION = "0.0.0-src" //set with ldflags

func main() {
	s := server.Server{
		Title:      "Cloud Torrent",
		Port:       3000,
		ConfigPath: "cloud-torrent.json",
	}

	o := opts.New(&s)
	o.Version(VERSION)
	o.PkgRepo()
	o.LineWidth = 96
	o.Parse()

	if err := s.Run(VERSION); err != nil {
		log.Fatal(err)
	}
}
