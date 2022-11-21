package modules

import "strings"

func GetExtension(filename []string) string {

	leng := 0

	for i, v := range filename {
		if v == "." {
			leng = i
			break
		}
	}

	return strings.Join(filename[leng:], "")
}
