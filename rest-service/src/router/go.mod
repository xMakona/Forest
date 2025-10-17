module router

go 1.22.1

require (
	auth v0.0.0-00010101000000-000000000000
	user v0.0.0-00010101000000-000000000000
	utils v0.0.0-00010101000000-000000000000
)

replace auth => ../api/auth

replace utils => ../utils

replace user => ../api/user
