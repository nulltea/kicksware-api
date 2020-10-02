package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"go.kicksware.com/api/service-common/config"
	"go.kicksware.com/api/service-common/core"
	"go.kicksware.com/api/service-common/core/meta"
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


func (c *communicator) PostMessage(endpoint string, query interface{}, response interface{}, params *meta.RequestParams) error {
	url := fmt.Sprintf(c.endpointFormat, endpoint)
	body, err := json.Marshal(query); if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body)); if err != nil {
		return err
	}
	token, err := c.Authenticate(params); if err != nil {
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

func (c *communicator) GetMessage(endpoint string, response interface{}, params *meta.RequestParams) error {
	url := fmt.Sprintf(c.endpointFormat, endpoint)
	req, err := http.NewRequest("GET", url, nil); if err != nil {
		return err
	}
	token, err := c.Authenticate(params); if err != nil {
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

func (c *communicator) Authenticate(params *meta.RequestParams) (string, error) {
	if params != nil && len(params.Token()) != 0 {
		return params.Token(), nil
	}
	token, err := c.auth.Authenticate(); if err != nil {
		log.Fatalln(errors.Wrap(err, "service-common::communicator: authenticate failed"))
		return "", err
	}
	return token, nil
}
