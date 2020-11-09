package crawler

type Interface interface {
	Scan(string, int) ([]Document, error)
}

type Document struct {
	ID    uint64
	Title string
	URL   string
}

func (d *Document) Ident() uint64 {
	return d.ID
}
