package engine

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"gosearch/pkg/index/fakeindex"
	"reflect"
	"testing"
)

func TestSearch(t *testing.T) {
	idx := fakeindex.New()

	engine := New(idx)

	_, err := engine.Search("как")
	if err != nil {
		t.Fatalf("ожидался результат, а получена ошибка: %s", err.Error())
	}

	wantErr := "пустой запрос"
	_, err = engine.Search("")
	if err != nil {
		if err.Error() != wantErr {
			t.Fatalf("ожидалась ошибка %s, а получена: %s", wantErr, err.Error())
		}
	}
}

func TestService_Search(t *testing.T) {
	want := []crawler.Document{
		crawler.Document{
			ID:    uint64(1),
			Title: "Как использовать git?",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(2),
			Title: "Прикладное применение подорожника как лекарство",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(3),
			Title: "Криптовалюта как будущее финансовой системы?",
			URL:   "http://localhost",
		},
	}
	type fields struct {
		Index index.Interface
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []crawler.Document
		wantErr bool
	}{
		{
			"Тестирование поиска",
			fields{
				fakeindex.New(),
			},
			args{
				"как",
			},
			want,
			false,
		},
		{
			"Тестирование поиска пустого запроса",
			fields{
				fakeindex.New(),
			},
			args{
				"",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Index: tt.fields.Index,
			}
			got, err := s.Search(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
