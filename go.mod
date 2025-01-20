module main

go 1.23.4

replace controller => ./controller

replace router => ./router

require (
	controller v0.0.0-00010101000000-000000000000
	github.com/gorilla/websocket v1.5.3
	router v0.0.0-00010101000000-000000000000
)
