module gosearch/pkg/index/invert

replace (
	gosearch/pkg/crawler => ../../crawler
	gosearch/pkg/index/invert/btree => ./btree
)

go 1.15

require (
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/invert/btree v0.0.0-00010101000000-000000000000
)
