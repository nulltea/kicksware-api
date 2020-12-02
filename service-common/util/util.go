package util

import (
	"bufio"
	"context"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/fatih/structs"
	"github.com/golang/glog"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/metadata"

	"go.kicksware.com/api/service-common/api/gRPC"
	"go.kicksware.com/api/service-common/core/meta"
)

func ToMap(v interface{}) map[string]interface{} {
	return structs.Map(v)
}

func ToBsonMap(v interface{}) (d bson.M, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &d)
	return
}

func ToBsonDoc(v interface{}) (d bson.D, err error) {
	data, err := bson.Marshal(v)
	if err != nil {
		return
	}
	err = bson.Unmarshal(data, &d)
	return
}

func GetAllInsertValues(v interface{}) []interface{} {
	return structs.Values(v)
}

func GetAllInsertColumns(v interface{}) []string {
	return structs.Names(v)
}

func GetInsertValues(v interface{}, fields []string) []interface{} {
	filter  := ToMap(v)
	values := funk.Map(fields, func(k interface{}) interface{}{
		key := k.(string)
		return filter[key]
	}).([]interface{})
	return values
}

func ToQueryMap(v url.Values) (qm map[string]interface{}) {
	qm = make(map[string]interface{})
	keys := funk.Keys(v).([]string)
	for _, key := range keys {
		if len(v[key]) > 1 {
			qm[key] = v[key]
			continue
		}
		qm[key] = v[key][0]
	}
	return
}

func GetPublicKey(keyPath string) *rsa.PublicKey {
	publicKeyFile, err := os.Open(keyPath)
	if err != nil {
		panic(err)
	}

	pemFileInfo, _ := publicKeyFile.Stat()
	var size int64 = pemFileInfo.Size()
	pemBytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pemBytes)

	data, _ := pem.Decode(pemBytes)

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes); if err != nil {
		panic(err)
	}

	publicKey, ok := publicKeyImported.(*rsa.PublicKey); if !ok {
		return nil
	}

	return publicKey
}

func RetrieveUserID(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		userIDs := md.Get(gRPC.UserContextKey)
		if len(userIDs) != 0 {
			return userIDs[0], true
		}
	}
	return "", false
}

func RetrieveAuthToken(ctx context.Context) (string, bool) {
	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		tokens := md.Get(gRPC.AuthMetaKey)
		if len(tokens) != 0 {
			return tokens[0], true
		}
	}
	return "", false
}

func GetErrorMsg(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func SetAuthParamsFromMetaData(ctx context.Context, params **meta.RequestParams) (ok bool) {
	var userID, token string
	userID, ok = RetrieveUserID(ctx)
	token, ok = RetrieveAuthToken(ctx)
	if ok {
		if *params == nil {
			*params = &meta.RequestParams{}
		}
		(*params).SetToken(token)
		(*params).SetUserID(userID)
		return true
	}
	return false
}

func NewTLSConfig(tlsConfig *meta.TLSCertificate) *tls.Config {
	if !tlsConfig.EnableTLS {
		return nil
	}
	certs := x509.NewCertPool()
	pem, err := ioutil.ReadFile(tlsConfig.CertFile); if err != nil {
		glog.Fatalln(err)
	}
	certs.AppendCertsFromPEM(pem)
	return &tls.Config{
		RootCAs: certs,
	}
}