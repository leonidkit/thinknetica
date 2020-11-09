module gosearch

go 1.15

replace (
	gosearch/pkg/crawler v0.0.0 => ./pkg/crawler
	gosearch/pkg/crawler/spider v0.0.0 => ./pkg/crawler/spider

	gosearch/pkg/engine v0.0.0 => ./pkg/engine
	gosearch/pkg/index v0.0.0 => ./pkg/index

	gosearch/pkg/index/invert v0.0.0 => ./pkg/index/invert
	gosearch/pkg/index/invert/btree => ./pkg/index/invert/btree
)

require (
	golang.org/x/net v0.0.0-20201109172640-a11eb1b685be // indirect
	gosearch/pkg/crawler v0.0.0 // indirect
	gosearch/pkg/crawler/spider v0.0.0 // indirect
	gosearch/pkg/engine v0.0.0 // indirect
	gosearch/pkg/index v0.0.0 // indirect
	gosearch/pkg/index/invert v0.0.0 // indirect
	gosearch/pkg/index/invert/btree v0.0.0-00010101000000-000000000000 // indirect
)
