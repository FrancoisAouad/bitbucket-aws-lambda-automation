#!/bin/bash
GO_VERSION=$(go version)

if [[ $GO_VERSION != *"go1.21.3"* ]]; then
  echo "Go version is not compatible"
  exit 1
fi