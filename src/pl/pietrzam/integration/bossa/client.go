package bossa

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

func RequestMetric(metric BossaMetric, properties BossaProperty) (string, error) {
	uri, ok := metricToUri[metric]
	if !ok {
		return "", errors.New(fmt.Sprintf("Cannot find uri for '%s' metric", metric))
	}
	return call(properties.BaseUrl + uri) // @todo maybe it should be safer?
}

func call(url string) (string, error) {

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, _ := netClient.Get(url)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return readBody(resp)
	} else {
		log.Warn("Requesting data failed with code '%d'", resp.StatusCode)
		return "", errors.New(fmt.Sprintf("Requesting data failed with code '%d'", resp.StatusCode))
	}
}

func readBody(response *http.Response) (string, error) {
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	bodyString := string(bodyBytes)
	return bodyString, nil
}
