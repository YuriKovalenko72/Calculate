package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func romeToArab(romNum string) int {
	romNums := map[string]int {
		"I": 1,
		"II": 2,
		"III": 3,
		"IV": 4,
		"V": 5,
		"VI": 6,
		"VII": 7,
		"VIII": 8,
		"IX": 9,
		"X": 10,
	}
	return romNums[romNum]
}

func arabToRome(arabNum int) string {

	romeNumber := map[int]string {
		1: "I",			
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6:  "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16:  "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
		21: "XXI",
		22: "XXII",
		23: "XXIII",
		24: "XXIV",
		25: "XXV",
		26: "XXVI",
		27: "XXVII",
		28:  "XXVIII",
		29: "XXIX",
		30: "XXX",
		31: "XXXI",
		32: "XXXII",
		33: "XXXIII",
		34: "XXXIV",
		35: "XXXV",
		36: "XXXVI",
		37: "XXXVII",
		38: "XXXVIII",
		39: "XXXIX",
		40: "XL",
		41: "XLI",
		42: "XLII",
		43: "XLIII",
		44: "XLIV",
		45: "XLV",
		46: "XLVI",
		47: "XLVII",
		48: "XLVIII",
		49: "XLIX",
		50: "L",
		51: "LI",
		52: "LII",
		53: "LIII",
		54: "LIV",
		55: "LV",
		56: "LVI",
		57: "LVII",
		58: "LVIII",
		59: "LIX",
		60: "LX",
		61: "LXI",
		62: "LXII",
		63: "LXIII",
		64: "LXIV",
		65: "LXV",
		66: "LXVI",
		67: "LXVII",
		68: "LXVIII",
		69: "LXIX",
		70: "LXX",
		71: "LXXI",
		72: "LXXII",
		73: "LXXIII",
		74: "LXXIV",
		75: "LXXV",
		76: "LXXVI",
		77: "LXXVII",
		78: "LXXVIII",
		79: "LXXIX",
		80: "LXXX",
		81: "LXXXI",
		82: "LXXXII",
		83: "LXXXIII",
		84: "LXXXIV",
		85: "LXXXV",
		86: "LXXXVI",
		87: "LXXXVII",
		88: "LXXXVIII",
		89: "LXXXIX",
		90: "XC",
		91: "XCI",
		92: "XCII",
		93: "XCIII",
		94: "XCIV",
		95: "XCV",
		96: "XCVI",
		97: "XCVII",
		98: "XCVIII",
		99: "XCIX",
		100: "C",
	}

	return romeNumber[arabNum]
}

func isRome(num string) bool {
	for _, char := range num {
		if !unicode.IsLetter(char) || (unicode.ToUpper(char) != 'I') && (unicode.ToUpper(char) != 'V') && (unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
} 

func main() {
	var a, b, s string
	var err = ""

	fmt.Println("Напишите ваш пример на арабских или римских цифрах\n" + "Введите выражение")
	fmt.Fscanln(os.Stdin, &a, &s, &b, &err)

	if err != "" {
		panic("Ошибка: не правильное выражение, формат - цифра операнд цифра!") 
	}

	if isRome(a) != isRome(b) {
		panic("Ошибка: обе цифры должны быть либо арабские(1, 2, 3...), либо римские(I, II, III...)!")
	}

	var numA, numB int

	if isRome(a) {
		numA = romeToArab(a)
	} else {
		numA, _ = strconv.Atoi(a)
	}

	if isRome(b) {
		numB = romeToArab(b)
	} else {
		numB, _ = strconv.Atoi(b)
	}

	if numB > 10 || numB < 1 || numA > 10 || numA < 1 {
		panic("Ошибка: укажите цифры от 1 до 10!")
	}

	var res int 

	switch s {
	case "+":
		res = numA + numB
	case "-":
		res = numA - numB
	case "*":
		res = numA * numB
	case "/":
		res = numA / numB
	default:
		panic("Ошибка: неверный операнд")
	}

	if isRome(a) && isRome(b) {
		var e = arabToRome(res)
		if e == "" {
			panic("Ошибка: отрицательных римских цифр не существует")
		}
		fmt.Println("Результат: ", arabToRome(res), "\nАрабские: ", res)
	} else {
		fmt.Println("Результат: ", res)
	}

}