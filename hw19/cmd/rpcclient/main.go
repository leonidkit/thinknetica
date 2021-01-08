package main

import (
	"bufio"
	"fmt"
	"gosearch/pkg/crawler"
	"log"
	"os"
	"strings"

	"net/rpc/jsonrpc"
)

type Query struct {
	Data string
}

func main() {
	host := "0.0.0.0"
	port := "8001"

	conn, err := jsonrpc.Dial("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		word := snr.Text()
		if strings.Replace(word, " ", "", -1) == "exit" {
			break
		}

		if word != "" {
			res := []crawler.Document{}
			query := Query{Data: word}

			err = conn.Call("RPCsrv.Search", query, &res)
			if err != nil {
				log.Printf("error: %q\n", err)
			}

			fmt.Println(res)
		}
	}
}
