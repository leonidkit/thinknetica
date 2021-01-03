module gosearch

go 1.15

replace (
	gosearch/pkg/api => ./pkg/api
	gosearch/pkg/crawler v0.0.0 => ./pkg/crawler
	gosearch/pkg/crawler/spider v0.0.0 => ./pkg/crawler/spider

	gosearch/pkg/engine v0.0.0 => ./pkg/engine
	gosearch/pkg/index v0.0.0 => ./pkg/index

	gosearch/pkg/index/tree v0.0.0 => ./pkg/index/tree
	gosearch/pkg/index/tree/btree => ./pkg/index/tree/btree

	gosearch/pkg/netsrv => ./pkg/netsrv
	gosearch/pkg/webapp => ./pkg/webapp

)

require (
	github.com/gorilla/mux v1.8.0 // indirect
	golang.org/x/net v0.0.0-20201224014010-6772e930b67b // indirect
	gosearch/pkg/api v0.0.0-00010101000000-000000000000 // indirect
	gosearch/pkg/crawler v0.0.0 // indirect
	gosearch/pkg/crawler/spider v0.0.0 // indirect
	gosearch/pkg/engine v0.0.0 // indirect
	gosearch/pkg/index v0.0.0 // indirect
	gosearch/pkg/index/tree v0.0.0 // indirect
	gosearch/pkg/index/tree/btree v0.0.0-00010101000000-000000000000 // indirect
)
