package helper

import (
	"os"
	"strconv"
)

func PaginationLimit() int {
	pagination, _ := strconv.Atoi(os.Getenv("pagination"))
	return pagination
}

func InArray(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

