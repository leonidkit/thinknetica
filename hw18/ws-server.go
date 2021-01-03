package main

type Server struct {
	Clients    []*client
	Unregister chan client
	Broadcast  chan string
}

func (s *Server) run() {
	go func() {
		for {
			select {
			case msg := <-s.Broadcast:
				for _, cl := range s.Clients {
					cl.writeMsg(msg)
				}
			case cli := <-s.Unregister:
				for i, cl := range s.Clients {
					if *cl == cli {
						s.Clients[i] = s.Clients[len(s.Clients)-1]
						s.Clients[len(s.Clients)-1] = &client{}
						s.Clients = s.Clients[:len(s.Clients)-1]
					}
				}
			}
		}
	}()
}
