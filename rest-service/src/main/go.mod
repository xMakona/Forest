module main

go 1.22.1

require router v0.0.0-00010101000000-000000000000

require (
	auth v0.0.0-00010101000000-000000000000 // indirect
	user v0.0.0-00010101000000-000000000000 // indirect
	utils v0.0.0-00010101000000-000000000000 // indirect
)

replace auth => ../auth

replace router => ../router

replace utils => ../utils

replace user => ../user
