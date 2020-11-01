module engine

go 1.15

replace engine/pkg/index v0.0.0 => ./pkg/index

replace engine/pkg/spider v0.0.0 => ./pkg/spider

require (
	engine/pkg/index v0.0.0
	engine/pkg/spider v0.0.0
)
