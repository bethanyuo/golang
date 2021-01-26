package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d ", i)
			// fmt.Println(i)   // Prints downwards on each separate line
		}
	}
	fmt.Println("") // 7 14 21 28 35 42 49

	for i := 1; i <= 50; i++ {
		if i%7 != 0 { // if i is not divisible by 7
			continue
		}
		fmt.Printf("%d ", i)

	}
	fmt.Println("") // 7 14 21 28 35 42 49

	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 { // if i is not divisible by 7
			continue
		}
		fmt.Printf("%d ", i)
		count++

		if count == 3 { // if i've already found 3 numbers, then break
			break
		}

	}
	fmt.Println("") // 7 14 21

	for i := 1; i <= 500; i++ {
		if i%5 == 0 && i%7 == 0 { // if i is divisible both by 7 and 5
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("") // 35 70 105 140 175 210 245 280 315 350 385 420 455 490

	for i := 1996; i <= 2020; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("") // 1996 1997 1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017 2018 2019 2020

	birthYear, currentYear := 1996, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++ // POST Statement moved inside of FOR LOOP
	}
	fmt.Println() // 1996 1997 1998 1999 2000 2001 2002 2003 2004 2005 2006 2007 2008 2009 2010 2011 2012 2013 2014 2015 2016 2017 2018 2019 2020

	age := -9
	if age < 0 || age > 100 {
		fmt.Println("Invalid Age")
	} else if age <= 18 {
		fmt.Println("You are minor!")
	} else if age == 18 {
		fmt.Println("Congratulations! You've just become major!")
	} else {
		fmt.Println("You are major!")
	}

	age = 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are a minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
