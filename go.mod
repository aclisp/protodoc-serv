module github.com/aclisp/protodoc-serv

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.10.0
	github.com/gobwas/ws v1.0.4
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/mux v1.7.3
	github.com/micro/go-micro/v2 v2.9.1
	github.com/stretchr/testify v1.7.0
)

// Use a production ready go-micro/v2 stable version maintained by nano-kit.
replace github.com/micro/go-micro/v2 => github.com/nano-kit/go-micro/v2 v2.10.2
