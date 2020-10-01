package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"github.com/timoth-y/kicksware-api/service-common/config"
	"github.com/timoth-y/kicksware-api/service-common/core"
)

type communicator struct {
	client               http.Client
	auth                 core.AuthService
	contentType          string
	innerServiceFormat   string
}

func NewCommunicator(auth core.AuthService, config config.CommonConfig) core.InnerCommunicator {
	return &communicator{
		http.Client{},
		auth,
		config.ContentType,
		config.InnerServiceFormat,
	}
}


func (c *communicator) PostMessage(service string, query interface{}, response interface{}) error {
	body, err := json.Marshal(query); if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", service, bytes.NewBuffer(body)); if err != nil {
		return err
	}
	token, err := c.authenticate(); if err != nil {
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
	err = json.Unmarshal(bytes, &response); if err != nil {
		return err
	}
	return nil
}

func (c *communicator) GetMessage(service string, response interface{}) error {
	req, err := http.NewRequest("GET", service, nil); if err != nil {
		return err
	}
	token, err := c.authenticate(); if err != nil {
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

func (c *communicator) authenticate() (string, error) {
	token, err := c.auth.Authenticate(); if err != nil {
		log.Fatalln(errors.Wrap(err, "service-common::communicator: authenticate failed"))
		return "", err
	}
	return token, nil
}
