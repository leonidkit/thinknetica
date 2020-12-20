module auth

go 1.15

replace pkg/jwtauth => ./pkg/jwtauth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	pkg/jwtauth v0.0.0-00010101000000-000000000000 // indirect
)
