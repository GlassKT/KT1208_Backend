package modules

import (
	"strings"
)

func FileName(filename string) string {

	fullname := strings.Split(filename, "")

	ind := 0

	for i, v := range fullname {
		if v == "." {
			ind = i
			break
		}
	}
	return strings.Join(fullname[:ind], "")
}
