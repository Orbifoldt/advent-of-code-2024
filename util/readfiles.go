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
	data, err := ReadInputMulti(day, useRealInput)
	if err != nil {
		return nil, err
	} else if len(data) == 0 {
		return nil, nil
	} else {
		return data[0], nil
	}
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


// Read all lines of the `input.txt` file of day [day]. Sections separated by a empty lines
// are put into separate string slices.
// When [useRealInput] is false, will use `sample.txt`
func ReadInputMulti(day int, useRealInput bool) ([][]string, error) {
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
	allData := make([][]string, 0)
	lines := []string{}
	for scanner.Scan() {
		trimmedLine := strings.TrimSpace(scanner.Text())
		if trimmedLine != "" {
			lines = append(lines, trimmedLine)
		} else {
			allData = append(allData, lines)
			lines = make([]string, 0)
		}
	}
	if len(lines) > 0 {
		allData = append(allData, lines)
	}

	return allData, nil
}
