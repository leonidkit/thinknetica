module gosearch/pkg/index/invert

replace (
    gosearch/pkg/crawler => ../../crawler
    gosearch/pkg/index/invert/btree => ./btree
)

go 1.15
