package hw11

import (
	"os"
	"reflect"
	"testing"
)

func TestMaxAge(t *testing.T) {
	type args struct {
		people []Human
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Тестирование основной функциональности",
			args{
				people: []Human{
					&Customer{24}, &Employee{22},
				},
			},
			24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAge(tt.args.people...); got != tt.want {
				t.Errorf("MaxAge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxAgeHuman(t *testing.T) {
	type args struct {
		people []interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			"Тестирование основной функциональности",
			args{
				people: []interface{}{
					&Customer{24}, &Employee{22},
				},
			},
			&Customer{24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxAgeHuman(tt.args.people...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxAgeHuman() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExamplePrint() {
	c1 := 1
	c2 := true
	c3 := "some string"

	Print(os.Stdout, c1, c2, c3)
	// Output: some string
}
