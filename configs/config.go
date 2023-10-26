package configs

var JENKINS_URL = "http://${BASIC_AUTH}@${JENKINS_ADDRESS}/job/{JOB_NAME}/config.xml"
var BITBUCKET_URL = "https://api.bitbucket.org/2.0/repositories/{WORKSPACE}/{SLUG}/refs/branches?q~name+%7E+%2Frelease/i&sort=-name"
var DIR = ".tmp"
var BRANCH_FILE = ".tmp/branches.txt"
