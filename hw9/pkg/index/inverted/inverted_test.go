package inverted

import (
	"flag"
	"gosearch/pkg/crawler"
	"gosearch/pkg/storage"
	"testing"
)

var (
	dataFilename = flag.String("f", "", "имя файла с данными в виде map в формате gob")

	data = []crawler.Document{
		crawler.Document{
			ID:    uint64(1),
			Title: "Как использовать git?",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(2),
			Title: "Прикладное применение подорожника",
			URL:   "http://localhost",
		},
		crawler.Document{
			ID:    uint64(3),
			Title: "Криптовалюта - будущее финансовой системы?",
			URL:   "http://localhost",
		},
	}
)

func TestInvertedTree_Find(t *testing.T) {
	type args struct {
		record string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"Тестирование поиска элемента",
			args{
				record: "прикладное",
			},
			data[1].URL,
			false,
		},
		{
			"Тестирование поиска отсутствующего элемента",
			args{
				record: "кек",
			},
			"index not found",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := NewIndexTree(data)
			got, err := i.Find(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Fatalf("InvertedTree.Find() error = %v, wantErr %v", err, tt.wantErr)
			}
			if (err != nil) && tt.wantErr {
				if err.Error() != tt.want {
					t.Fatalf("InvertedTree.Find() error = %v, want %v", err, tt.wantErr)
				}
				return
			}

			if len(got) == 0 {
				t.Fatal("len(InvertedTree.Find()) = 0, want 1")
			}
			if got[0].URL != tt.want {
				t.Errorf("InvertedTree.Find().URL = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		value  uint64
		source []crawler.Document
	}
	tests := []struct {
		name    string
		args    args
		want    uint64
		wantErr bool
	}{
		{
			"Тестирование поиска элемента",
			args{
				uint64(1),
				data,
			},
			uint64(0),
			false,
		},
		{
			"Тестирование поиска отсутствующего элемента",
			args{
				uint64(10),
				data,
			},
			uint64(0),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := binarySearch(tt.args.value, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("binarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkNewIndexTree(b *testing.B) {
	var err error
	if *dataFilename != "" {
		flr := storage.New()
		data, err = flr.LoadFile(*dataFilename)
		if err != nil {
			b.Fatal(err.Error())
		}
	}

	for i := 0; i <= b.N; i++ {
		indx := NewIndexTree(data)
		_ = indx
	}
}

func BenchmarkNewIndexList(b *testing.B) {
	var err error
	if *dataFilename != "" {
		flr := storage.New()
		data, err = flr.LoadFile(*dataFilename)
		if err != nil {
			b.Fatal(err.Error())
		}
	}

	for i := 0; i <= b.N; i++ {
		indx := NewIndexList(data)
		_ = indx
	}
}

func BenchmarkInvertedTree_Find(b *testing.B) {
	var err error
	if *dataFilename != "" {
		flr := storage.New()
		data, err = flr.LoadFile(*dataFilename)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
	indx := NewIndexTree(data)

	for i := 0; i <= b.N; i++ {
		res, err := indx.Find("habr")
		if err != nil {
			continue
		}
		_ = res
	}
}

func BenchmarkInvertedList_Find(b *testing.B) {
	var err error
	if *dataFilename != "" {
		flr := storage.New()
		data, err = flr.LoadFile(*dataFilename)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
	indx := NewIndexList(data)

	for i := 0; i <= b.N; i++ {
		res, err := indx.Find("habr")
		if err != nil {
			continue
		}
		_ = res
	}
}
