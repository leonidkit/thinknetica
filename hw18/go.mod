module hw18

go 1.15

replace (
	hw18/pkg/api => ./pkg/api
	hw18/pkg/ws-server => ./pkg/ws-server
)

require (
	hw18/pkg/api v0.0.0-00010101000000-000000000000 // indirect
	hw18/pkg/ws-server v0.0.0-00010101000000-000000000000 // indirect
)
