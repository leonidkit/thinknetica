package rpcsrv

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/engine"
	"log"
	"net"
	"net/rpc"

	"github.com/powerman/rpc-codec/jsonrpc2"
)

type RPCsrv struct {
	engine *engine.Service
}

type Query struct {
	Data string
}

func (r *RPCsrv) Search(query Query, result *[]crawler.Document) error {
	res, err := r.engine.Search(query.Data)
	if err != nil {
		return err
	}

	*result = res
	return nil
}

func Serve(host, port string, e *engine.Service) {
	listener, err := net.Listen("tcp4", host+":"+port)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer listener.Close()

	err = rpc.Register(&RPCsrv{e})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		go jsonrpc2.ServeConn(conn)
	}
}
