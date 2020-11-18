module gosearch/pkg/index/inverted

replace (
	gosearch/pkg/crawler => ../../crawler
	gosearch/pkg/index/inverted/btree => ./btree
)

go 1.15

require (
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/inverted/btree v0.0.0-00010101000000-000000000000
)
