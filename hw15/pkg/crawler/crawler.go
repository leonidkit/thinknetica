package crawler

type Interface interface {
	BatchScan([]string, int, int) (chan Document, chan error)
	Scan(string, int) ([]Document, error)
}

type Document struct {
	ID    uint64
	Title string
	URL   string
}

func (d Document) Ident() uint64 {
	return d.ID
}
