module gosearch/pkg/engine

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/index => ../index
	gosearch/pkg/index/fakeindex => ../index/fakeindex
)

go 1.15

require (
	gosearch/pkg/crawler v0.0.0-00010101000000-000000000000
	gosearch/pkg/index v0.0.0-00010101000000-000000000000
	gosearch/pkg/index/fakeindex v0.0.0-00010101000000-000000000000
)
