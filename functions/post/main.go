package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/apex/log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/unee-t/env"
)

func handler(ctx context.Context, evt json.RawMessage) (string, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return "", err
	}
	err = post2Case(cfg, evt)
	return "", err
}

// For event notifications https://github.com/unee-t/lambda2sns/tree/master/tests/events
func post2Case(cfg aws.Config, evt json.RawMessage) (err error) {
	e, err := env.New(cfg)
	if err != nil {
		return err
	}
	casehost := fmt.Sprintf("https://%s", e.Udomain("case"))
	APIAccessToken := e.GetSecret("API_ACCESS_TOKEN")
	url := casehost + "/api/ses" // refer to Postman
	log.Infof("Posting to: %s, payload %s, with header key %s", url, evt, APIAccessToken)
	req, err := http.NewRequest("POST", url, strings.NewReader(string(evt)))
	if err != nil {
		log.WithError(err).Error("constructing POST")
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+APIAccessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.WithError(err).Error("POST request")
		return err
	}
	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.WithError(err).Error("failed to read body")
		return err
	}
	if res.StatusCode == http.StatusOK {
		log.Infof("Response code %d, Body: %s", res.StatusCode, string(resBody))
	} else {
		log.Warnf("Response code %d, Body: %s", res.StatusCode, string(resBody))
	}
	return err
}

func main() {
	lambda.Start(handler)
}
