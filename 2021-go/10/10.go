package _10

import (
	"sort"
)

const (
	// ): 3 points.
	// ]: 57 points.
	// }: 1197 points.
	// >: 25137 points.
	parenthesis = 3
	squareBrace = 57
	curlyBrace  = 1197
	greaterThan = 25137
)

func totalScoreCorrupted(lines []string) int {
	var totalScore int
	for _, line := range lines {
		totalScore += scoreCorrupted(line)
	}
	return totalScore
}

func scoreCorrupted(line string) int {
	stack := make([]byte, 0)
	for _, v := range line {
		b := byte(v)
		if isOpener(b) {
			stack = append(stack, b)
			continue
		}
		// closer
		lastOpener := stack[len(stack) - 1]
		if closes(b, lastOpener) {
			stack = stack[:len(stack) - 1]
		} else {
			return scoreOfCorrupted(b)
		}
	}
	return 0
}

func middleScoreIncomplete(lines []string) int {
	var scores []int
	for _, l := range lines {
		if isIncomplete(l) {
			scores = append(scores, scoreCompletion(getCompletion(l)))
		}
	}
	sort.Ints(scores)
	return scores[(len(scores)-1) / 2]
}

func isIncomplete(line string) bool {
	return scoreCorrupted(line) == 0
}

func getCompletion(line string) []byte {
	stack := make([]byte, 0)
	for _, v := range line {
		b := byte(v)
		if isOpener(b) {
			stack = append(stack, b)
			continue
		}
		// closer
		lastOpener := stack[len(stack) - 1]
		if closes(b, lastOpener) {
			stack = stack[:len(stack) - 1]
		} else {
			return nil
		}
	}
	// reverse the stack and switch to closers
	for i, j := 0, len(stack) - 1; i <= j; i, j = i + 1, j - 1 {
		stack[i], stack[j] = toCloser(stack[j]), toCloser(stack[i])
	}
	//fmt.Printf("%s - %d points\n", stack, scoreCompletion(stack))
	return stack
}

func scoreCompletion(closers []byte) int {
	var score int
	for _, b := range closers {
		score *= 5
		score += scoreOfCompletion(b)
	}
	return score
}

func toCloser(opener byte) byte {
	switch opener {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	return opener
}

func isOpener(b byte) bool {
	switch b {
	case '(', '[', '{', '<':
		return true
	}
	return false
}

func closes(closer byte, opener byte) bool {
	switch opener {
	case '(':
		return closer == ')'
	case '[':
		return closer == ']'
	case '{':
		return closer == '}'
	case '<':
		return closer == '>'
	}
	return false
}

func scoreOfCorrupted(b byte) int {
	switch b {
	case ')':
		return parenthesis
	case ']':
		return squareBrace
	case '}':
		return curlyBrace
	case '>':
		return greaterThan
	}
	return 0
}

func scoreOfCompletion(b byte) int {
	switch b {
	case ')':
		return 1
	case ']':
		return 2
	case '}':
		return 3
	case '>':
		return 4
	}
	return 0
}