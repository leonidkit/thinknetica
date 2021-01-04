package wshub

type Hub struct {
	Clients    []*Client
	Unregister chan Client
	Broadcast  chan string
}

func (s *Hub) Run() {
	go func() {
		for {
			select {
			case msg := <-s.Broadcast:
				for _, cl := range s.Clients {
					cl.WriteMsg(msg)
				}
			case cli := <-s.Unregister:
				for i, cl := range s.Clients {
					if *cl == cli {
						s.Clients[i] = s.Clients[len(s.Clients)-1]
						s.Clients[len(s.Clients)-1] = &Client{}
						s.Clients = s.Clients[:len(s.Clients)-1]
					}
				}
			}
		}
	}()
}
