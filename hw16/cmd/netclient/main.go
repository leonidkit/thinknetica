package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	host := "localhost"
	port := "8000"

	conn, err := net.Dial("tcp4", host+":"+port)
	if err != nil {
		log.Fatal(err.Error())
	}

	enter := "Enter word to find: "
	snr := bufio.NewScanner(os.Stdin)

	for fmt.Print(enter); snr.Scan(); fmt.Print(enter) {
		word := snr.Text()
		if strings.Replace(word, " ", "", -1) == "exit" {
			break
		}
		if word != "" {
			_, err := conn.Write([]byte(word + "\n"))
			if err != nil {
				log.Fatalf("writing error: %s", err.Error())
			}
		}

		rdr := bufio.NewReader(conn)
		for {
			res, err := rdr.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("error occured: %s", err.Error())
			}

			fmt.Println(res)
			if res == "done\n" || strings.Contains(res, "error occured") {
				break
			}
		}
	}
}
