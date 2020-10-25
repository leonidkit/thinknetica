module engine

go 1.15

require (
	crawler/pkg/spider v0.0.0
	crawler/pkg/spiderblank v0.0.0
)

replace crawler/pkg/spider v0.0.0 => ../crawler/pkg/spider

replace crawler/pkg/spiderblank v0.0.0 => ../crawler/pkg/spiderblank