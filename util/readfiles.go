package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Read all lines of the `input.txt` file of day [day] into a string slice. 
// When [useRealInput] is false, will use `sample.txt`
func ReadInput(day int, useRealInput bool) ([]string, error) {
	// Need to find project root since current work dir depends on where execution starts
	modulePath, err := findModuleFilePath()
	if err != nil {
		return nil, err
	}

	var filename string
	if useRealInput {
		filename = "input.txt"
	} else {
		filename = "sample.txt"
	}

	filePath := filepath.Join(modulePath, "resources", fmt.Sprintf("day%02d", day), filename)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())
		lines = append(lines, trimmedLine)
	}

	return lines, nil
}



func findModuleFilePath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dir := cwd
	for {
		currentDirEntries, err := os.ReadDir(dir)
		if err != nil {
			return "", err
		}

		for _, entry := range currentDirEntries {
			if(!entry.IsDir() && entry.Name() == "go.mod") {
				return dir, nil
			}
		}

		// No module found, move to parent folder
		oldPath := dir
		dir = filepath.Dir(dir)
		if(dir == oldPath){
			return "", fmt.Errorf("did not find Go module in '%v'", cwd)
		}
	}
}
