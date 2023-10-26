package bitbucket

import (
	"example/lib"
	"os"
	"sort"

	"github.com/Jeffail/gabs"
)

var logger = lib.Log

func HandleBranches(data *gabs.Container) []string {
	var branchVector []string
	childrenMap, err := data.ChildrenMap()
	if err != nil {
		logger.WithError(err).Error("Error occurred while performing the operation")
		os.Exit(1)
	}
	// loop through the bitbucket api response
	for key := range childrenMap {
		branchVector = append(branchVector, key)
	}
	//sort the vector of branches in desc order
	sort.Slice(branchVector, func(i, j int) bool {
		return branchVector[i] > branchVector[j]
	})

	logger.Info("New Branches have been added.")
	return branchVector
}
