package handler

import (
	"strings"
)

const base = ""

func PathMultipleEntries(url string) string {
	return strings.Trim(base + url, "/")
}

func PathSingleEntry(url string, insert string) (colPath, docPath, id string) {
	url = strings.Trim(base + url, "/")
	splitted := strings.Split(url, "/")
	last := len(splitted) - 1
	splittedDocPath := make([]string, last)
	splittedColPath := make([]string, last)
	copy(splittedDocPath, splitted[:last])
	copy(splittedColPath, splitted[:last])
	docPath = strings.Join(append(splittedDocPath, insert, splitted[last]), "/")
	colPath = strings.Join(append(splittedColPath, insert), "/")
	return colPath, docPath, splitted[last]
}