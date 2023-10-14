package main

import (
	"fmt"
	"strconv"
	"strings"
)

func checkIfNumber(char string) bool {
	if _, err := strconv.Atoi(char); err == nil {
		return true
	}
	return false
}

func checkIfBackslash(char string) bool {
	if strings.Contains(char, "\\") {
		return true
	}
	return false
}

func checkNumsUntilSymbol(s string) (int, int) {
	i := 0
	num := 0
	for _, c := range s {
		if checkIfNumber(string(c)) {
			temp, _ := strconv.Atoi(string(c))
			num = num*10 + temp
			i++
		} else {
			return num, i
		}
	}
	return num, i
}

func unpacking(str string) string {
	var result, temp, prev string
	var indexFlag, numFlag, slashFlag bool
	var num int = 0

	s := []rune(str)
	for i := 0; i < len(s); i++ {
		temp = string(s[i])
		if checkIfNumber(temp) && !slashFlag {
			numFlag = true
		} else {
			numFlag = false
		}

		if i != len(s)-1 {
			indexFlag = true
		} else {
			indexFlag = false
		}

		if i != 0 {
			prev = string(s[i-1])
		}

		if indexFlag && checkIfNumber(string(s[i+1])) && !numFlag {

			if slashFlag = checkIfBackslash(temp); slashFlag {
				if prev == "\\" {
					num, _ = strconv.Atoi(string(s[i+1]))
					result += strings.Repeat(temp, num)
					slashFlag = false
				} else {
					continue
				}
			} else {
				num, k := checkNumsUntilSymbol(string(s[i+1:]))
				result += strings.Repeat(string(s[i]), num)
				i += k
			}

		} else if slashFlag = checkIfBackslash(temp); slashFlag {
			continue
		} else if !numFlag {
			result += temp
		}
	}
	return result
}

func main() {
	var givenStr, result string

	fmt.Print("Given string: ")
	fmt.Scan(&givenStr)

	result = unpacking(givenStr)
	if result == "" {
		fmt.Println("Given incorrect string")
	} else {
		fmt.Println("Unpacked string: ", result)
	}
}
