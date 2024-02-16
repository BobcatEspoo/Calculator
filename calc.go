package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romeToArab(x string) int {
	var i int
	inter := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rome := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := range inter {
		z := x[0:1] == `X`
		l := len(x) > 2

		a := strconv.Itoa(inter[i])
		if x == a {
			panic(fmt.Sprintf("введите только арабские или только римские"))
		}
		if z == true {
			if l == true {
				h := x[0:2]
				if (h == "I") || (h == "V") {
					panic(fmt.Sprintf("Введите число меньшее или равное 10"))
				}
			}

		}
		if rome[i] == x {
			return i + 1 //добавить конвертацию римских в арабские
		}

	}
	return i + 1
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rez := 0
	var rezR string
	//var in bool
	for {
		fmt.Println("Введите выражение в формате a+b")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text) // считывание
		text = strings.ReplaceAll(text, " ", "")
		rome1 := strings.Contains(text, "I")
		rome5 := strings.Contains(text, "V")
		rome10 := strings.Contains(text, "X")
		roman := rome1 || rome5 || rome10
		if strings.Contains(text, ".") || strings.Contains(text, ",") {
			panic(fmt.Sprintf("Введите целое число"))
		}
		if roman == false {
			in := strings.Contains(text, "+")
			if strings.Count(text, "+") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "+", -2) // массив
				toNumber0, _ := strconv.Atoi(a[0])
				toNumber1, _ := strconv.Atoi(a[1])
				if (toNumber0 < 1) || (toNumber1 < 1) {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				if (toNumber0 > 10) || (toNumber1 > 10) {
					panic(fmt.Sprintf("Введите число меньшее или равное 10"))
				}
				rez = toNumber0 + toNumber1
				fmt.Println(rez)
			}
			in = strings.Contains(text, "-")
			if strings.Count(text, "-") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "-", -2) // массив
				toNumber0, _ := strconv.Atoi(a[0])
				toNumber1, _ := strconv.Atoi(a[1])
				if (toNumber0 < 1) || (toNumber1 < 1) {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				if (toNumber0 > 10) || (toNumber1 > 10) {
					panic(fmt.Sprintf("Введите число меньшее или равное 10"))
				}
				rez = toNumber0 - toNumber1
				fmt.Println(rez)
			}
			in = strings.Contains(text, "*")
			if strings.Count(text, "*") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "*", -2) // массив
				toNumber0, _ := strconv.Atoi(a[0])
				toNumber1, _ := strconv.Atoi(a[1])
				if (toNumber0 < 1) || (toNumber1 < 1) {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				if (toNumber0 > 10) || (toNumber1 > 10) {
					panic(fmt.Sprintf("Введите число меньшее или равное 10"))
				}
				rez = toNumber0 * toNumber1
				fmt.Println(rez)
			}
			in = strings.Contains(text, "/")
			if strings.Count(text, "/") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "/", -2) // массив
				toNumber0, _ := strconv.Atoi(a[0])
				toNumber1, _ := strconv.Atoi(a[1])
				if (toNumber0 < 1) || (toNumber1 < 1) {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				if (toNumber0 > 10) || (toNumber1 > 10) {
					panic(fmt.Sprintf("Введите число меньшее или равное 10"))
				}
				rez = toNumber0 / toNumber1
				fmt.Println(rez)
			}

		} else {
			in := strings.Contains(text, "+")
			if strings.Count(text, "+") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "+", -2) // массив
				if strings.Contains(text, "0") {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				toNumber0 := romeToArab(a[0])
				toNumber1 := romeToArab(a[1])
				rez = toNumber0 + toNumber1
				rezR = integerToRoman(rez)
				fmt.Println(rezR)

			}
			in = strings.Contains(text, "-")
			if strings.Count(text, "-") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "-", -2) // массив
				if strings.Contains(text, "0") {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				toNumber0 := romeToArab(a[0])
				toNumber1 := romeToArab(a[1])
				rez = toNumber0 - toNumber1
				if rez < 1 {
					panic(fmt.Sprintf("Введите число от I до X"))
				}
				rezR = integerToRoman(rez)
				fmt.Println(rezR)

			}

			in = strings.Contains(text, "*")
			if strings.Count(text, "*") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "*", -2) // массив
				if strings.Contains(text, "0") {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				toNumber0 := romeToArab(a[0])
				toNumber1 := romeToArab(a[1])
				rez = toNumber0 * toNumber1
				rezR = integerToRoman(rez)
				fmt.Println(rezR)

			}

			in = strings.Contains(text, "/")
			if strings.Count(text, "/") > 1 {
				panic(fmt.Sprintf("Введите корректно"))
			}
			if in == true {
				a := strings.SplitN(text, "/", -2) // массив
				if strings.Contains(text, "0") {
					panic(fmt.Sprintf("Введите число большее или равное 1"))
				}
				toNumber0 := romeToArab(a[0])
				toNumber1 := romeToArab(a[1])
				rez = toNumber0 / toNumber1
				rezR = integerToRoman(rez)
				fmt.Println(rezR)
			}

		}
		prov1 := strings.Contains(text, "+")
		prov2 := strings.Contains(text, "-")
		prov3 := strings.Contains(text, "*")
		prov4 := strings.Contains(text, "/")
		if prov1 || prov2 || prov3 || prov4 {
			fmt.Println("")
		} else {
			panic(fmt.Sprintf("Введите корректно"))
		}

	}
}
