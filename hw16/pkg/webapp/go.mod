module gosearch/pkg/webapp

go 1.15

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/engine => ../engine
	gosearch/pkg/index => ../index
	gosearch/pkg/index/fakeindex => ../index/fakeindex
)