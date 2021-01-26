package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {

	var name string = "Bethany"
	country := "USA"

	fmt.Printf("Your name: %q\n", name)  // Your name: "Bethany"
	fmt.Printf("Country: %s\n", country) // Country: USA

	fmt.Printf("Your name: %s\nCountry: %s\n", name, country)

	//equivalent to: // Just like this below
	fmt.Printf(`Your name: %s
Country: %s
`, name, country)

	fmt.Println()
	fmt.Println(strings.Repeat("#", 50))
	fmt.Println("He says: \"Hello\"") // He says: "Hello"
	fmt.Println("C:\\Users\\a.txt")   // C:\Users\a.txt

	var r rune = 'ă'
	fmt.Printf("Type of r is = %T\n", r) // Type of r is = int32 ===> rune is alias to int32

	str1, str2 := "ma", "m"

	fmt.Println("New Word = " + str1 + str2 + string(r)) // New Word = mamă

	s1 := "țară means country in Romanian"

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s1); i++ {

		fmt.Printf("%v ", s1[i])
	}

	fmt.Println() // Bytes in string: 200 155 97 114 196 131 32 109 101 97 110 115 32 99 111 117 110 116 114 121 32 105 110 32 82 111 109 97 110 105 97 110

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}
	fmt.Println()

	/*
	   byte index: 0 -> rune: ț
	   byte index: 2 -> rune: a
	   byte index: 3 -> rune: r
	   byte index: 4 -> rune: ă
	   byte index: 6 -> rune:
	   byte index: 7 -> rune: m
	   byte index: 8 -> rune: e
	   byte index: 9 -> rune: a
	   byte index: 10 -> rune: n
	   byte index: 11 -> rune: s
	   byte index: 12 -> rune:
	   byte index: 13 -> rune: c
	   byte index: 14 -> rune: o
	   byte index: 15 -> rune: u
	   byte index: 16 -> rune: n
	   byte index: 17 -> rune: t
	   byte index: 18 -> rune: r
	   byte index: 19 -> rune: y
	   byte index: 20 -> rune:
	   byte index: 21 -> rune: i
	   byte index: 22 -> rune: n
	   byte index: 23 -> rune:
	   byte index: 24 -> rune: R
	   byte index: 25 -> rune: o
	   byte index: 26 -> rune: m
	   byte index: 27 -> rune: a
	   byte index: 28 -> rune: n
	   byte index: 29 -> rune: i
	   byte index: 30 -> rune: a
	   byte index: 31 -> rune: n
	*/

	str3 := "Go is cool!"
	x := str3[0]
	fmt.Println(x) // 71

	//str3[0] = 'x' // ERROR -> cannot assign to str3[0]

	// printing the number of runes in the string
	fmt.Println(len(str3))                    // 11
	fmt.Println(utf8.RuneCountInString(str3)) // 11

	s := "你好 Go!"

	// converting string to byte slice
	b := []byte(s)

	// printing out the byte slice
	fmt.Printf("%#v\n", b) // []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd, 0x20, 0x47, 0x6f, 0x21}

	for i, v := range b {
		fmt.Printf("Byte Index: %d :: Value: %v\n", i, v)
	}

	// Iterating over the byte slice and printing out each index and byte in the byte slice
	/*
		Byte Index: 0 :: Value: 228
		Byte Index: 1 :: Value: 189
		Byte Index: 2 :: Value: 160
		Byte Index: 3 :: Value: 229
		Byte Index: 4 :: Value: 165
		Byte Index: 5 :: Value: 189
		Byte Index: 6 :: Value: 32
		Byte Index: 7 :: Value: 71
		Byte Index: 8 :: Value: 111
		Byte Index: 9 :: Value: 33
	*/

	// converting string to rune slice
	r1 := []rune(s)

	// printing out the rune slice
	fmt.Printf("%#v\n", r1) // []int32{20320, 22909, 32, 71, 111, 33}

	for i, v := range r1 {
		fmt.Printf("Byte Index: %d :: Value: %c\n", i, v)
	}

	// Iterating over the rune slice and printing out each index and rune in the rune slice
	/*
		Byte Index: 0 :: Value: 你
		Byte Index: 1 :: Value: 好
		Byte Index: 2 :: Value:
		Byte Index: 3 :: Value: G
		Byte Index: 4 :: Value: o
		Byte Index: 5 :: Value: !
	*/

	var i = 3
	var f = 3.2
	var ss1, ss2 = "3.14", "5"

	// 1. int to string
	ss := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", ss, ss) // s Type is string, s value is "3"

	// 2. string to int
	is, err := strconv.Atoi(ss2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is) // i type is int, i value is 5
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	sss1 := fmt.Sprintf("%f", f)
	fmt.Printf("sss1's type: %T, sss1's value: %s\n", sss1, sss1) // sss1's type: string, sss1's value: 3.200000

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(ss1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1) // f1's type: float64, f1's value: 3.14
	} else {
		fmt.Println("Cannot convert string to float64.")
	}

}
