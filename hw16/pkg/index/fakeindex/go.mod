module gosearch/pkg/index/fakeindex

replace (
	gosearch/pkg/crawler => ../../crawler
	gosearch/pkg/index/tree/btree => ../tree/btree
)


go 1.15
