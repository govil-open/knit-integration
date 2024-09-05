package util

import "strings"

func ConvertNewline(str, nlcode string) string {
	return strings.NewReplacer(
		"\r\n", nlcode,
		"\r", nlcode,
		"\n", nlcode,
	).Replace(str)
}

func Suffix(data string, separator string) string {
	var output string
	if data == "" {
		output = ""
	} else {
		output = data + separator
	}
	return output
}
