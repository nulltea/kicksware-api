package pipes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/fatih/structs"

	"search-service/core/meta"
	"search-service/core/model"
	"search-service/core/pipe"
	"search-service/env"
)

type referencePipe struct {
	contentType          string
	bindingServiceFormat string
	bindingServiceName   string
}

func NewSneakerReferencePipe(config env.CommonConfig) pipe.SneakerReferencePipe {
	return &referencePipe{
		config.ContentType,
		config.InnerServiceFormat,
		"references-service",
	}
}

func (p *referencePipe) FetchOne(code string) (ref *model.SneakerReference, err error) {
	err = p.getFromDataService(p.requestResource("/", code), ref)
	return
}

func (p *referencePipe) Fetch(codes []string, params meta.RequestParams) (refs []*model.SneakerReference, err error) {
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

func (p *referencePipe) FetchAll(params meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		err = p.getFromDataService(p.requestResource("?", paramValues.Encode()), &refs)
		return
	}
	err = p.getFromDataService(p.requestResource("/"), refs)
	return
}

func (p *referencePipe) FetchQuery(query meta.RequestQuery, params meta.RequestParams) (refs[]*model.SneakerReference, err error) {
	resource := p.requestResource("/query")
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		resource = p.requestResource("/query", "?", paramValues.Encode())
	}
	refs, err = p.postOnDataService(resource, query)
	return
}

func (p *referencePipe) getFromDataService(service string, target interface{}) error {
	resp, err := http.Get(service); if err != nil {
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
	resp, err := http.Post(service, p.contentType, bytes.NewBuffer(body))
	if err != nil {
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
	return fmt.Sprintf(p.bindingServiceFormat, p.bindingServiceName, strings.Join(res, ""))
}

func requestParamValues(params meta.RequestParams) url.Values {
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

