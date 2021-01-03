module hw18/pkg/ws-api

go 1.15

require github.com/gorilla/websocket v1.4.2

replace (
    hw18/pkg/ws-server => ../ws-server
)
