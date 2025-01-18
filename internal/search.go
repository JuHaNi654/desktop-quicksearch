package internal

/* search uses Boyer-Moore Algorithm */

import (
	"strings"
)

func computeFullShift(shiftArr, longerSuffArray *[]int, pattern string) {
	n := len(pattern)

	i := n
	j := n + 1

	(*longerSuffArray)[i] = j

	for i > 0 {
		for j <= n && pattern[i-1] != pattern[j-1] {
			if (*shiftArr)[j] == 0 {
				(*shiftArr)[j] = j - i
			}

			j = (*longerSuffArray)[j]
		}

		i--
		j--
		(*longerSuffArray)[i] = j
	}
}

func computeGoodSuffix(shiftArr, longerSuffArray *[]int, pattern string) {
	n := len(pattern)
	j := (*longerSuffArray)[0]

	for i := 0; i < n; i++ {
		if (*shiftArr)[i] == 0 {
			(*shiftArr)[i] = j
			if i == j {
				j = (*longerSuffArray)[j]
			}
		}
	}
}

func Search(text, pattern string) int {
	text = strings.ToLower(text)
	pattern = strings.ToLower(pattern)

	patternLen := len(pattern)
	textLen := len(text)

	longerSuffArray := make([]int, (patternLen + 1))
	shiftArr := make([]int, (patternLen + 1))

	// Initialize shift array
	for i := 0; i <= patternLen; i++ {
		shiftArr[i] = 0
	}

	computeFullShift(&shiftArr, &longerSuffArray, pattern)
	computeGoodSuffix(&shiftArr, &longerSuffArray, pattern)

	shift := 0
	index := -1
	for shift <= (textLen - patternLen) {
		j := patternLen - 1

		for j >= 0 && pattern[j] == text[shift+j] {
			j--
		}

		if j < 0 {
			index++
			shift += shiftArr[0]
		} else {
			shift += max(1, shiftArr[j+1])
		}
	}

	return index
}
