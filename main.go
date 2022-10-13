package main

import "flag"

var (
	dir  = flag.String("dir", "", "")
	port = flag.Int("port", 8080, "listen port")
)

func main() {
	flag.Parse()

	if *dir == "" {
		panic("empty dir")
	}

	Run()
}
