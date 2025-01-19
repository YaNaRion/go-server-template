module main

go 1.23.4

replace controller => ./controller

replace router => ./router

require (
	controller v0.0.0-00010101000000-000000000000
	router v0.0.0-00010101000000-000000000000
)
