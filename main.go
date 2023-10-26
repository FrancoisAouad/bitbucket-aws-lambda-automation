package main

import (
	"example/src/configs"
	"example/src/fs"
	"example/src/http"
	"example/src/lib"
	"example/src/services"
	"example/src/services/bitbucket"

	"github.com/sirupsen/logrus"
)

var logger = lib.Log

func main() {
	lib.LoggerInit(logger)
	logger.WithFields(logrus.Fields{"type": "System", "job": "AWS-Lambda"}).Info("Started executing Lambda function to Update Jenkins Job Param")
	// const url = "http://pokeapi.co/api/v2/pokedex/kanto/"
	data := http.GetApiData(configs.BITBUCKET_URL)
	branchList := bitbucket.HandleBranches(data)
	fs.Write(branchList, configs.DIR, configs.BRANCH_FILE)
	services.ExecShell()
	fs.Clean(configs.DIR, configs.BRANCH_FILE)
	// POST to jenkins in order to update
	http.PostApi(configs.JENKINS_URL)
	logger.Info("Updated Jenkins job config.xml successfully.")
}
