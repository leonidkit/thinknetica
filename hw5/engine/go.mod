module engine

go 1.15

replace engine/pkg/index v0.0.0 => ./pkg/index

replace engine/pkg/spider v0.0.0 => ./pkg/spider

replace engine/pkg/index/btree v0.0.0 => ./pkg/index/btree

require (
	engine/pkg/index/btree v0.0.0 // indirect
	engine/pkg/index v0.0.0
	engine/pkg/spider v0.0.0
)
