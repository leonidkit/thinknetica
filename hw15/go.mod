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
