module gosearch

go 1.15

replace (
	gosearch/pkg/crawler v0.0.0 => ./pkg/crawler
	gosearch/pkg/crawler/spider v0.0.0 => ./pkg/crawler/spider

	gosearch/pkg/engine v0.0.0 => ./pkg/engine
	gosearch/pkg/index v0.0.0 => ./pkg/index

	gosearch/pkg/index/inverted v0.0.0 => ./pkg/index/inverted
	gosearch/pkg/index/inverted/btree => ./pkg/index/inverted/btree

	gosearch/pkg/netsrv => ./pkg/netsrv
	gosearch/pkg/webapp => ./pkg/webapp
)

require (
	gosearch/pkg/crawler v0.0.0 // indirect
	gosearch/pkg/crawler/spider v0.0.0 // indirect
	gosearch/pkg/engine v0.0.0 // indirect
	gosearch/pkg/index v0.0.0 // indirect
	gosearch/pkg/index/inverted v0.0.0 // indirect
	gosearch/pkg/index/inverted/btree v0.0.0-00010101000000-000000000000 // indirect
	gosearch/pkg/webapp v0.0.0-00010101000000-000000000000 // indirect
)
