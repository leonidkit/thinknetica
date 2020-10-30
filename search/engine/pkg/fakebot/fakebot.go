package fakebot

type LocalScan struct{}

func NewBot() *LocalScan {
	return &LocalScan{}
}

func (l *LocalScan) Scan(url string, depth int) (data map[string]string, err error) {
	data = map[string]string{
		"https://habr.ru":         "Главная",
		"https://habr.ru/contact": "Контакты",
	}

	return data, nil
}
