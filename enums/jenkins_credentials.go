package enums

type JenkinsCredential string

const (
	JENKINS_BASIC_AUTH JenkinsCredential = "JENKINS_BASIC_AUTH"
	JENKINS_IP         JenkinsCredential = "JENKINS_IP"
)
