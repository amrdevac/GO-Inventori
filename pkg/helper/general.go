package helper

import (
	"os"
	"strconv"
)

func PaginationLimit() int {
	pagination, _ := strconv.Atoi(os.Getenv("pagination"))
	return pagination
}
