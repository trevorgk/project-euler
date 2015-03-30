package main

import (
	"fmt"
	"strings"
	//"github.com/trevorgk/project-euler/eulerlib"
)

//	If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?
func euler_17() {

	sum := 0
	for i := 1; i <= 1000; i++ {
		numericText := strings.TrimSpace(chunk(i))
		sum += len(strings.Replace(numericText, " ", "", -1))
	}
	fmt.Println("Answer is:", sum)
}

func chunk(num int) string {

	million, thousand, hundred := 1000000, 1000, 100

	if (num >= million) {
		mod := num%million
		if (mod == 0) {
			return chunk(num/million) + " million"
		}
		return chunk(num/million) + " million " + chunk(num%million)
	}

	if (num >= thousand) {
		mod := num%thousand
		if (mod == 0) {
			return chunk(num/thousand) + " thousand"
		}
		return chunk(num/thousand) + " thousand " + chunk(num%thousand)
	}

	if (num > hundred) {
		mod := num%hundred
		if (mod == 0) {
			return chunk(num/hundred) + " hundred"
		}

		return chunk(num/hundred) + " hundred and " + chunk(num%hundred)
	}

	if (num >= 20){
		rem := num%10
		switch (num / 10){
		case 2:
			return "twenty " + chunk(rem)
		case 3:
			return "thirty " + chunk(rem)
		case 4:
			return "forty " + chunk(rem)
		case 5:
			return "fifty " + chunk(rem)
		case 6:
			return "sixty " + chunk(rem)
		case 7:
			return "seventy " + chunk(rem)
		case 8:
			return "eighty " + chunk(rem)
		case 9:
			return "ninety " + chunk(rem)
		}
	}

	switch(num){
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	}

	return ""
}
