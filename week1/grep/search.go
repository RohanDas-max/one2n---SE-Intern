package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

//function to search string from a file/folder
func search(filename, args string) ([]string, error) {
	var result []string
	var mapData = make(map[string]string)
	if err := filepath.Walk(filename, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else {
			dt, _ := os.ReadFile(path)
			i := 0
			for range dt {
				mapData[path] += string(dt[i])
				i++
			}
			return nil
		}
	}); err != nil {
		return nil, err
	} else {
		for fn, res := range mapData {
			if strings.Contains(res, args) {
				result = append(result, fmt.Sprintf("%s: %s found in this file\n", fn, args))
			} else if !strings.Contains(res, args) {
				result = append(result, fmt.Sprintf("%s: %s is not available in this file \n", fn, args))
			}
		}
	}
	return result, nil
}
