param([Int32]$day) 

$day_string = 'day{0:D2}' -f $day

# Create input and sample files
pushd
cd .\resources
mkdir $day_string
cd $day_string
touch "input.txt"
touch "sample.txt"
popd

# Create the solution and unit test .go files
pushd
cd .\days
mkdir $day_string
cd $day_string

touch "$day_string.go"
$solutionTemplate = @"
package $day_string

import (
	"advent-of-code-2024/util"
	"fmt"
)

func SolvePart1(useRealInput bool) (int, error) {
	_, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func SolvePart2(useRealInput bool) (int, error) {
	_, err := parseInput(useRealInput)
	if err != nil {
		return 0, err
	}

	return 0, nil
}

func parseInput(useRealInput bool) ([]string, error) {
	data, err := util.ReadInputMulti($day, useRealInput)
	if err != nil {
		return nil, err
	}
	if len(data) != 1 {
		return nil, fmt.Errorf("expected single line of input")
	}

	return data[0], nil
}
"@ 
$solutionTemplate | Out-File -FilePath "$day_string.go"

touch "${day_string}_test.go"
$testTemplate = @"
package $day_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCorrectlyDeterminePart1OnExampleInput(t *testing.T) {
	solution, err := SolvePart1(false)
	if err != nil {
		t.Fatalf("Error during Solve: %v", err)
	}
	assert.Equal(t, 111, solution)
}

func TestShouldCorrectlyDeterminePart2OnExampleInput(t *testing.T) {
	solution, err := SolvePart2(false)
	if err != nil `{
		t.Fatalf("Error dkuring Solve: %v", err)
	}
	assert.Equal(t, 111, solution)
}
"@ 
$testTemplate | Out-File -FilePath "${day_string}_test.go"

popd
