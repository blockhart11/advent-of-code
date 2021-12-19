package _18

import (
	"strconv"
)

func sumAndReduceAll(sn []string) string {
	result := sn[0]
	for _, v := range sn[1:] {
		//fmt.Println("Processing numbers", result, " and ", v)
		result = reduce(sum(result, v))
	}
	return result
}

func sum(lhs, rhs string) string {
	result := "[" + lhs + "," + rhs + "]"
	//fmt.Println("Sum:", result)
	return result
}

func reduce(sn string) string {
	result := sn
	// go til you can't go no more
	for {
		// look for explosion
		//fmt.Println("Scanning:", result)
		var openCount int
		var foundSomething bool
		for i, v := range result {
			switch v {
			case '[':
				openCount++
				if openCount > 4 {
					// explode!
					result = explode(result, i)
					foundSomething = true
				}
			case ']':
				openCount--
			}
			if foundSomething {
				break
			}
		}
		// if exploded, restart loop
		if foundSomething {
			continue
		}

		// check for split
		for i, _ := range result[:len(result)-1] {
			nextNum, err := strconv.Atoi(result[i : i+2])
			if err == nil {
				// found a 2 digit number
				result = split(result, nextNum, i)
				foundSomething = true
				break
			}
		}
		if !foundSomething {
			break
		}
	}
	//fmt.Println(result, " is fully reduced")
	return result
}

//     i
// [[[[[1,2]]]]]
// [[[[[15,20]]]]]
func explode(sn string, index int) string {
	lhsResult := sn[:index]
	//parse things
	marker := index + 1
	var lhsExplodeInt, rhsExplodeInt int
	lhsExplodeInt, _ = strconv.Atoi(string(sn[marker]))
	marker++
	if sn[marker] != ',' {
		// two digit left num
		nextDigit, _ := strconv.Atoi(string(sn[marker]))
		lhsExplodeInt = lhsExplodeInt*10 + nextDigit
		marker++
	}
	marker++
	rhsExplodeInt, _ = strconv.Atoi(string(sn[marker]))
	marker++
	if sn[marker] != ']' {
		// two digit right num
		nextDigit, _ := strconv.Atoi(string(sn[marker]))
		rhsExplodeInt = rhsExplodeInt*10 + nextDigit
		marker++
	}
	marker++
	rhsResult := sn[marker:]

	//fmt.Println("Exploding", sn[index:marker], ", ", lhsExplodeInt, rhsExplodeInt)

	// to the left, to the left
	for i := len(lhsResult) - 1; i >= 0; i-- {
		leftNum, err := strconv.Atoi(string(lhsResult[i]))
		if err == nil {
			// we got one
			// check if 2 digits
			nextDigit, err := strconv.Atoi(string(lhsResult[i-1]))
			if err == nil {
				leftNum = nextDigit*10 + leftNum
				lhsResult = lhsResult[:i-1] + strconv.Itoa(leftNum+lhsExplodeInt) + lhsResult[i+1:]
			} else {
				lhsResult = lhsResult[:i] + strconv.Itoa(leftNum+lhsExplodeInt) + lhsResult[i+1:]
			}
			break
		}
	}

	// to the right, to the right
	for i, v := range rhsResult {
		rightNum, err := strconv.Atoi(string(v))
		if err == nil {
			// we got a number
			// check if 2 digits
			nextDigit, err := strconv.Atoi(string(rhsResult[i+1]))
			if err == nil {
				rightNum = rightNum*10 + nextDigit
				rhsResult = rhsResult[:i] + strconv.Itoa(rightNum+rhsExplodeInt) + rhsResult[i+2:]
			} else {

				rhsResult = rhsResult[:i] + strconv.Itoa(rightNum+rhsExplodeInt) + rhsResult[i+1:]
			}
			break
		}
	}
	return lhsResult + "0" + rhsResult
}

//      i
// [[[[[10,2]]]]]
func split(sn string, splitInt int, index int) string {
	lhsResult := sn[:index]
	rhsResult := sn[index+2:]
	lhs := splitInt / 2
	rhs := splitInt / 2
	if splitInt%2 == 1 {
		rhs++
	}
	//fmt.Println("Splitting", splitInt, "into", lhs, rhs)
	return lhsResult + "[" + strconv.Itoa(lhs) + "," + strconv.Itoa(rhs) + "]" + rhsResult
}

func magnitude(sn string) int {
	// find delimiter, lhs and rhs
	marker := 1
	var depth int
	var lhsStr, rhsStr string
	for {
		var done bool
		switch sn[marker] {
		case '[':
			depth++
		case ']':
			depth--
		case ',':
			if depth == 0 {
				// found the middle
				lhsStr = sn[1:marker]
				rhsStr = sn[marker+1 : len(sn)-1]
				done = true
			}
		}
		marker++
		if done {
			break
		}
	}

	var lhs, rhs int
	// base case
	if len(lhsStr) == 1 {
		lhs, _ = strconv.Atoi(lhsStr)
	} else {
		lhs = magnitude(lhsStr)
	}
	if len(rhsStr) == 1 {
		rhs, _ = strconv.Atoi(rhsStr)
	} else {
		rhs = magnitude(rhsStr)
	}
	return lhs*3 + rhs*2
}
