module gosearch/pkg/engine

replace (
	gosearch/pkg/crawler => ../crawler
	gosearch/pkg/index => ../index
)

go 1.15