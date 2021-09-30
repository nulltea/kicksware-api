package REST

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/fatih/structs"
	"go.kicksware.com/api/shared/api/rest"
	"go.kicksware.com/api/shared/config"
	"go.kicksware.com/api/shared/core"

	"go.kicksware.com/api/services/references/core/model"

	"go.kicksware.com/api/shared/core/meta"

	"go.kicksware.com/api/services/search/core/pipe"
)

type referencePipe struct {
	config        config.CommonConfig
	communicator  core.InnerCommunicator
	fetchResource string
}

func NewSneakerReferencePipe(auth core.AuthService, config config.CommonConfig) pipe.SneakerReferencePipe {
	return &referencePipe{
		config,
		rest.NewCommunicator(auth, config),
		"sneaker/references",
	}
}


func (p *referencePipe) FetchOne(code string) (ref *model.SneakerReference, err error) {
	err = p.communicator.GetMessage(p.requestEndpoint(code), ref)
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
	err = p.communicator.GetMessage(p.requestEndpoint("?", values.Encode()), &refs)
	return
}

func (p *referencePipe) FetchAll(params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		err = p.communicator.GetMessage(p.requestEndpoint("?", paramValues.Encode()), &refs)
		return
	}
	err = p.communicator.GetMessage(p.requestEndpoint(), &refs)
	return
}

func (p *referencePipe) FetchQuery(query meta.RequestQuery, params *meta.RequestParams) (refs []*model.SneakerReference, err error) {
	resource := p.requestEndpoint("query")
	if paramValues := requestParamValues(params); paramValues != nil && len(paramValues) != 0 {
		resource = p.requestEndpoint("query", "?", paramValues.Encode())
	}
	err = p.communicator.PostMessage(resource, query, &refs)
	return
}

func (p *referencePipe) requestEndpoint(res ...string) string {
	return path.Join(p.fetchResource, strings.Join(res, ""))
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
