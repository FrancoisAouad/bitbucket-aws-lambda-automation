package http

import (
	"example/lib"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Jeffail/gabs"
)

var logger = lib.Log

func GetApiData(url string) *gabs.Container {
	//aws.GetAWSSecrets()
	response, err := http.Get(url)

	if err != nil {
		logger.WithError(err).Error("Failed to make GET request")
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.WithError(err).Fatal("Error occurred when trying to read response body")
		os.Exit(1)
	}
	jsonParsed, err := gabs.ParseJSON(responseData)
	if err != nil {
		panic(err)
	}

	return jsonParsed
}

func PostApi(url string) {
	// reponse, err := http.Post()
}
