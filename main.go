// test project main.go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	maxNumberTriad = 3
	numberThousand = 1
	numberDigits   = 10
	numberForms    = 3
	female         = 1
	male           = 0
	maxNumber      = 1000000000
)

var (
	wordFirstDecade = [2][numberDigits]string{
		{"ноль", "один", "два", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"},
		{"ноль", "одна", "две", "три", "четыре", "пять", "шесть", "семь", "восемь", "девять"},
	}

	wordSecondDecade  = [numberDigits]string{"", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать", "пятнадцать",
		"шестнадцать", "семнадцать", "восемнадцать", "девятнадцать"}

	wordOverDecade = [numberDigits]string{"", "десять", "двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто"}

	wordHundred  = [numberDigits]string{"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот"}

	wordOverHundred = [maxNumberTriad][numberForms]string{
		{"", "", ""},
		{"тысяча", "тысячи", "тысяч"},
		{"миллион", "миллиона", "миллионов"},
	}
)

/*
triadToWords - функция преобразует трехзначное число в словестную форму,
а также добавляет к числу обозначение разрядности числа (миллион, тысяча и т.д.)
triad uint - трехзначное число
numberTriad -номер разряда
sex - число обозначающее какое род будет использоваться для обозначения разрядности
*/

func triadToWords(triad uint, numberTriad uint8, sex uint8) string {

	firstDigit := triad / 10 / 10 % 10
	secondDigit := triad / 10 % 10
	thirdDigit := triad % 10
	wordSecondDecadeState := false

	var words string

	if firstDigit != 0 {
		words += wordHundred[firstDigit] + " "
	}

	if secondDigit != 0 {
		if secondDigit == 1 && thirdDigit > 0 {
			words += wordSecondDecade[thirdDigit] + " "
			wordSecondDecadeState = true
		} else {
			words += wordOverDecade[secondDigit] + " "
		}
	}

	if thirdDigit != 0 && !wordSecondDecadeState {
		words += wordFirstDecade[sex][thirdDigit] + " "
	}

	if thirdDigit == 0 || thirdDigit > 4 {

		words += wordOverHundred[numberTriad][2] + " "

	}

	if thirdDigit > 1 && thirdDigit < 5 {
		if wordSecondDecadeState {
			words += wordOverHundred[numberTriad][2] + " "
		} else {
			words += wordOverHundred[numberTriad][1] + " "
		}
	}

	if thirdDigit == 1 {
		words += wordOverHundred[numberTriad][0] + " "
	}

	return words
}

/*
numberToWords - функция преобразует трехзначное число в словестную форму,
а также добавляет к числу обозначение разрядности числа (миллион, тысяча и т.д.)
возращает словестную форму числа
*/

func numberToWords(number uint) string {

	if number == 0 {
		return wordFirstDecade[0][0]
	}

	var words string

	numberTriad :=0

	for number > 0 && numberTriad <= maxNumberTriad {
		if number%1000 != 0 {
			if numberTriad == numberThousand {
				words = triadToWords(number%1000, uint8(numberTriad), female) + words
			} else {
				words = triadToWords(number%1000, uint8(numberTriad), male) + words
			}
		}

		number = number / 1000
		numberTriad++
	}

	words = strings.TrimSpace(words)

	return words
}

func main() {
	var numberStr string

	fmt.Scanf("%s", &numberStr)

	if number, err := strconv.Atoi(numberStr); err == nil {

		if number > 0 && number < maxNumber {
			fmt.Println(numberToWords(uint(number)))
		} else {
			fmt.Println("number is not in range")
		}
	} else {
		fmt.Println("not number")
	}

}
