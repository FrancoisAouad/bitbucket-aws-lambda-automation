#!/bin/bash
BRANCH_FILE=$(cat .tmp/branches.txt)
JENKINS_IP=$1
JENKINS_USERNAME=$2
JENKINS_PASSWORD=$3

WORKSPACE_DIR="/var/jenkins_home/jobs/bitbucket-lambda"

ssh -i "$JENKINS_USERNAME:$JENKINS_PASSWORD@$JENKINS_IP"

sudo su - 
mkdir -p $WORKSPACE_DIR/workspace/.tmp
cd "$WORKSPACE_DIR/.tmp"
echo "$BRANCH_FILE" > branches.txt
cp  $WORKSPACE_DIR/config.xml $WORKSPACE_DIR/config.txt

# handle key insertions here in xml



