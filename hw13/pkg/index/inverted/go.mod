module gosearch/pkg/index/inverted

replace (
	gosearch/pkg/crawler => ../../crawler
	gosearch/pkg/index/inverted/btree => ./btree
)

go 1.15