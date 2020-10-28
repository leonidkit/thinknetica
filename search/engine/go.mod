module engine

go 1.15

require (
	crawler/pkg/fakebot v0.0.0
	crawler/pkg/spider v0.0.0
	indexer/pkg/index v0.0.0
)

replace crawler/pkg/spider v0.0.0 => ../crawler/pkg/spider

replace crawler/pkg/fakebot v0.0.0 => ../crawler/pkg/fakebot

replace indexer/pkg/index v0.0.0 => ../indexer/pkg/index
