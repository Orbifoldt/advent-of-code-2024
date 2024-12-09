package day09

import (
	"advent-of-code-2024/util"
	"fmt"
	"strconv"
)

func SolvePart1(useRealInput bool) (int64, error) {
	diskMap, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	disk := make([]int, 0)

	// Write all files onto the disk
	isFreeSpace := false
	fileId := 0
	for _, size := range diskMap {
		if !isFreeSpace {
			for range size {
				disk = append(disk, fileId)
			}
			fileId++
		} else {
			for range size {
				disk = append(disk, -1)
			}
		}
		isFreeSpace = !isFreeSpace
	}

	// Move the file blocks
	firstEmptyIndex := 0
	for i := len(disk) - 1; i >= 0 && firstEmptyIndex < i-1; i-- {
		sourceValue := disk[i]
		if sourceValue >= 0 {
			for j := firstEmptyIndex; j < len(disk); j++ {
				if disk[j] < 0 {
					disk[j] = sourceValue
					disk[i] = -1
					firstEmptyIndex = j
					break
				}
			}
		}
	}

	return checksum(disk), nil
}

func SolvePart2(useRealInput bool) (int64, error) {
	diskMap, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	disk := make([]int, 0)
	fileMap := make([]int, len(diskMap)/2+1)
	fileSizes := make([]int, len(diskMap)/2+1)

	// Write all files onto the disk and keep track of where they are and how big
	isFreeSpace := false
	fileId := 0
	diskIndex := 0
	for _, size := range diskMap {
		if !isFreeSpace {
			for range size {
				disk = append(disk, fileId)
			}
			fileMap[fileId] = diskIndex
			fileSizes[fileId] = size
			fileId++
		} else {
			for range size {
				disk = append(disk, -1)
			}
		}
		diskIndex += size
		isFreeSpace = !isFreeSpace
	}

	// Move whole files to where there is space
	for fileId := len(fileMap) - 1; fileId >= 0; fileId-- {
		currentIndex := fileMap[fileId]
		fileSize := fileSizes[fileId]
		if fileSize == 0 {
			continue
		}

		// Find where there is sufficient space
		targetIndex := -1
		consequetiveFreeSpace := 0
		for j := 0; j < currentIndex; j++ {
			if disk[j] < 0 {
				consequetiveFreeSpace++
			} else {
				consequetiveFreeSpace = 0
			}

			if consequetiveFreeSpace == fileSize {
				targetIndex = (j + 1) - consequetiveFreeSpace
				break
			}
		}

		// Move the file
		if targetIndex > 0 {
			fileMap[fileId] = targetIndex
			for i := range fileSize {
				disk[targetIndex+i] = fileId
				disk[currentIndex+i] = -1
			}
		}

	}

	return checksum(disk), nil
}

func checksum(disk []int) int64 {
	checksum := int64(0)
	for idx, x := range disk {
		if x >= 0 {
			checksum += int64(idx * x)
		}
	}
	return checksum
}

func parseInput(useRealInput bool) ([]int, error) {
	data, err := util.ReadInput(9, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	diskMap := make([]int, 0)
	for _, r := range data[0] {
		x, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, err
		}
		diskMap = append(diskMap, x)
	}

	return diskMap, nil
}
