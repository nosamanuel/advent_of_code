package main

import (
	"sort"
)

func locate(ids []string) string {
	size := len(ids[0])
	for i := 0; i < size; i++ {
		difference := locateSingleDifference(ids, i)
		if difference != "" {
			return difference
		}
	}

	return "N/A"
}

func locateSingleDifference(ids []string, location int) string {
	var sortIds []string
	for _, id := range ids {
		sortId := id[:location] + id[location+1:]
		sortIds = append(sortIds, sortId)
	}

	sort.Strings(sortIds)
	for i := 1; i < len(sortIds); i++ {
		if sortIds[i] == sortIds[i-1] {
			return sortIds[i]
		}
	}

	return ""
}
