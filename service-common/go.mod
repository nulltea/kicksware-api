module github.com/timoth-y/kicksware-api/service-common

go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fatih/structs v1.1.0
	github.com/go-chi/chi v4.1.2+incompatible
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/soheilhy/cmux v0.1.4
	github.com/thoas/go-funk v0.7.0
	github.com/timoth-y/kicksware-api/user-service v0.0.0-20200917011049-79140f0e7480
	go.elastic.co/apm/module/apmgrpc v1.8.0
	go.elastic.co/apm/module/apmhttp v1.8.0
	go.mongodb.org/mongo-driver v1.4.1
	google.golang.org/grpc v1.32.0
)
