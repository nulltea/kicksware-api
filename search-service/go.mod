module github.com/timoth-y/kicksware-api/search-service

go 1.14

require (
	github.com/Masterminds/squirrel v1.4.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/jmoiron/sqlx v1.2.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/olivere/elastic/v7 v7.0.20
	github.com/pkg/errors v0.9.1
	github.com/thoas/go-funk v0.7.0
	github.com/timoth-y/kicksware-api/product-service v0.0.0-20200917011049-79140f0e7480
	github.com/timoth-y/kicksware-api/reference-service v0.0.0-20200917011049-79140f0e7480
	github.com/timoth-y/kicksware-api/service-common v0.0.0-20200917005139-98b85de071c8
	github.com/timoth-y/kicksware-api/user-service v0.0.0-20200917011049-79140f0e7480
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	go.mongodb.org/mongo-driver v1.4.1
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
)
