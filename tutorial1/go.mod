module tutorial/hello

go 1.23.2

require (
	golang.org/x/example/hello v0.0.0-20241014184706-d7b0ac127859
	rsc.io/quote v1.5.2
	tutorial/greetings v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace tutorial/greetings => ../greetings
