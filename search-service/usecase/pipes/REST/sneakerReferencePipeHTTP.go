package REST

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"github.com/pkg/errors"

	"github.com/timoth-y/kicksware-api/reference-service/core/model"

	"github.com/timoth-y/kicksware-api/service-common/core/meta"

	"github.com/timoth-y/kicksware-api/search-service/core/pipe"
	"github.com/timoth-y/kicksware-api/search-service/core/service"
	"github.com/timoth-y/kicksware-api/search-service/env"
)

type referencePipe struct {
	client               http.Client
	auth                 service.AuthService
	contentType          string
	bindingServiceFormat string
	bindingServiceEndpoint   string
}

func NewSneakerReferencePipe(auth service.AuthService, config env.CommonConfig) pipe.SneakerReferencePipe {
	return &referencePipe{
		http.Client{},
		auth,
		config.ContentType,
		config.InnerServiceFormat,
		"references/sneakers",
	}
}

func (p *referencePipe) authenticate() (string, error) {
	token, err := p.auth.Authenticate(); if err != nil {
		log.Fatalln(errors.Wrap(err, "search-service::startup.InnerServiceAuth: authenticate failed"))
		return "", err
	}
	return token, nil
}


func (p *referencePipe) FetchOne(code string) (ref *model.SneakerReference, err error) {
	err = p.getFromDataService(p.requestResource("/", code), ref)
	return
}

func (p *referencePipe) Fetch(codes []string, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	values := url.Values{}
	for _, code := range codes {
		values.Add("referenceId", code)
	}
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		for key, param := range paramValues {
			values.Add(key, param[0])
		}
	}
	err = p.getFromDataService(p.requestResource("?", values.Encode()), &refs)
	return
}

func (p *referencePipe) FetchAll(params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		err = p.getFromDataService(p.requestResource("?", paramValues.Encode()), &refs)
		return
	}
	err = p.getFromDataService(p.requestResource("/"), &refs)
	return
}

func (p *referencePipe) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs[]*model.SneakerReference, err error) {
	resource := p.requestResource("/query")
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		resource = p.requestResource("/query", "?", paramValues.Encode())
	}
	refs, err = p.postOnDataService(resource, query)
	return
}

func (p *referencePipe) getFromDataService(service string, target interface{}) error {
	req, err := http.NewRequest("GET", service, nil); if err != nil {
		return err
	}
	token, err := p.authenticate(); if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	resp, err := p.client.Do(req); if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, target); if err != nil {
		return err
	}
	return nil
}

func (p *referencePipe) postOnDataService(service string, query interface{}) (references []*model.SneakerReference, err error) {
	body, err := json.Marshal(query); if err != nil {
		return
	}
	req, err := http.NewRequest("POST", service, bytes.NewBuffer(body)); if err != nil {
		return nil, err
	}
	token, err := p.authenticate(); if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", p.contentType)
	resp, err := p.client.Do(req); if err != nil {
		return
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return
	}
	err = json.Unmarshal(bytes, &references); if err != nil {
		return
	}
	return
}

func (p *referencePipe) requestResource(res ...string) string {
	return fmt.Sprintf(p.bindingServiceFormat, p.bindingServiceEndpoint, strings.Join(res, ""))
}

func requestParamValues(params *meta.RequestParams) url.Values {
	if params == nil {
		return nil
	}
	values := url.Values{}
	properties := structs.Map(params)
	for prop := range properties {
		val := properties[prop]
		switch v := val.(type) {
		case nil:
			continue
		case string:
			if v != "" {
				values.Add("referenceId", v)
			}
		case int:
			if v != 0 {
				values.Add("referenceId", strconv.Itoa(v))
			}
		case float32, float64:
			if v != 0 {
				values.Add("referenceId", fmt.Sprint(v))
			}
		case bool:
			values.Add("referenceId", strconv.FormatBool(v))
		default:
			values.Add("referenceId", fmt.Sprint(v))
		}

	}
	return values
}
