package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	decode := flag.Bool("d", false, "Enable Decoding")
	encode := flag.Bool("e", false, "Enable Encoding")

	flag.Parse()
	args := os.Args[1:]

	if *decode {
		str, err := base64.StdEncoding.DecodeString(args[1])
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(str))
	}

	if *encode {
		str := base64.StdEncoding.EncodeToString([]byte(args[1]))

		fmt.Println(str)
	}
}
