package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"go.kicksware.com/api/shared/config"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/core/meta"
)

type communicator struct {
	client         http.Client
	auth           core.AuthService
	contentType    string
	endpointFormat string
}

func NewCommunicator(auth core.AuthService, config config.CommonConfig) core.InnerCommunicator {
	return &communicator{
		http.Client{},
		auth,
		config.ContentType,
		config.ApiEndpointFormat,
	}
}


func (c *communicator) PostMessage(endpoint string, query interface{}, response interface{}, params ...*meta.RequestParams) error {
	url := fmt.Sprintf(c.endpointFormat, endpoint)
	body, err := json.Marshal(query); if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body)); if err != nil {
		return err
	}
	token, err := c.authenticate(params...); if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", c.contentType)
	resp, err := c.client.Do(req); if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, response); if err != nil {
		return err
	}
	return nil
}

func (c *communicator) GetMessage(endpoint string, response interface{}, params ...*meta.RequestParams) error {
	url := fmt.Sprintf(c.endpointFormat, endpoint)
	req, err := http.NewRequest("GET", url, nil); if err != nil {
		return err
	}
	token, err := c.authenticate(params...); if err != nil {
		return err
	}
	req.Header.Set("Authorization", token)
	resp, err := c.client.Do(req); if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body); if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, response); if err != nil {
		return err
	}
	return nil
}

func (c *communicator) authenticate(params ...*meta.RequestParams) (string, error) {
	if len(params) > 0 {
		if param := params[0]; param != nil && len(param.Token()) != 0 {
			return param.Token(), nil
		}
	}
	token, err := c.auth.Authenticate(); if err != nil {
		log.Fatalln(errors.Wrap(err, "service-common::communicator: authenticate failed"))
		return "", err
	}
	return token, nil
}
