module gosearch/pkg/api

go 1.15

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/engine => ../engine
	gosearch/pkg/index => ../index
	gosearch/pkg/index/fakeindex => ../index/fakeindex
)

require (
	github.com/gorilla/mux v1.8.0
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000 // indirect
	gosearch/pkg/engine v0.0.0-00010101000000-000000000000
	gosearch/pkg/index v0.0.0-00010101000000-000000000000 // indirect
	gosearch/pkg/index/fakeindex v0.0.0-00010101000000-000000000000
)
