module gosearch/pkg/engine

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/index => ../index
	gosearch/pkg/index/fakeindex => ../index/fakeindex
)

go 1.15