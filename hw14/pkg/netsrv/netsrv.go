package netsrv

import (
	"bufio"
	"gosearch/pkg/engine"
	"net"
	"strings"
)

type Netsrv struct {
	host   string
	port   string
	engine engine.Service
}

func New(host string, port string, engine engine.Service) *Netsrv {
	return &Netsrv{
		host:   host,
		port:   port,
		engine: engine,
	}
}

// Принимает соединение с запросом, в ответ пишет результат поиска по запросу
func (n *Netsrv) searchInteractive(conn net.Conn) {
	query := bufio.NewReader(conn)

	for {
		line, err := query.ReadString('\n')
		if err != nil {
			conn.Write([]byte("error occured:" + err.Error() + "\n"))
			return
		}

		lineclr := strings.TrimRight(line, "\n\r")
		if lineclr == "exit" {
			conn.Close()
			return
		}

		docs, err := n.engine.Search(lineclr)
		if err != nil {
			conn.Write([]byte("error occured:" + err.Error() + "\n"))
			continue
		}

		for _, doc := range docs {
			_, err := conn.Write([]byte(doc.URL + " - " + doc.Title + "\n"))
			if err != nil {
				conn.Write([]byte("error occured:" + err.Error() + "\n"))
				return
			}
		}
	}
}

func (n *Netsrv) Serve() error {
	lstnr, err := net.Listen("tcp4", n.host+":"+n.port)
	if err != nil {
		return err
	}

	for {
		conn, err := lstnr.Accept()
		if err != nil {
			return err
		}

		go n.searchInteractive(conn)
	}
}
