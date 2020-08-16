module github.com/timoth-y/kicksware-platform/middleware-service/search-service

go 1.14

require (
	github.com/Masterminds/squirrel v1.2.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/go-chi/chi v4.1.1+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/olivere/elastic/v7 v7.0.12
	github.com/pkg/errors v0.9.1
	github.com/thoas/go-funk v0.6.0
	github.com/timoth-y/kicksware-platform/middleware-service/product-service v0.0.0-20200807133030-5f847defdf0b
	github.com/timoth-y/kicksware-platform/middleware-service/reference-service v0.0.0-20200807133030-5f847defdf0b
	github.com/timoth-y/kicksware-platform/middleware-service/service-common v0.0.0-20200816165922-25c13d486d1e
	github.com/timoth-y/kicksware-platform/middleware-service/user-service v0.0.0-20200807131113-a8928eed241c
	github.com/vmihailenco/msgpack v3.3.3+incompatible
	go.mongodb.org/mongo-driver v1.3.3
	google.golang.org/grpc v1.31.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.2.4
)
