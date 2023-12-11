package read_file

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func readFile(filePath string) ([]string, error) {
	// os.ReadFile(filePath) as alternative ??
	// Read entire contents of file
	data, err := ioutil.ReadFile(filename) 
	if err != nil {
		return nil, err
	}


	// Convert hte byte slice to a string and split it into lines
	lines := strings.Split(string(content), "\n")

	// Remove the last empty string (resulting from the trailing newline)
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return lines, nil
}


