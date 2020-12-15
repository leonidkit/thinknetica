package engine

import (
	"gosearch/pkg/crawler"
	"gosearch/pkg/index"
	"gosearch/pkg/index/fakeindex"
	"reflect"
	"testing"
)

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
			name: "Тестирование поиска",
			fields: fields{
				fakeindex.New(),
			},
			args: args{
				"как",
			},
			want:    want,
			wantErr: false,
		},
		{
			name: "Тестирование поиска пустого запроса",
			fields: fields{
				fakeindex.New(),
			},
			args: args{
				"",
			},
			want:    nil,
			wantErr: true,
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
