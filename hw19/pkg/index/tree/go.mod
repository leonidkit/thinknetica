module gosearch/pkg/index/inverted

replace (
	gosearch/pkg/crawler => ../../crawler
	gosearch/pkg/index => ../index
	gosearch/pkg/index/tree/btree => ./btree
)

go 1.15