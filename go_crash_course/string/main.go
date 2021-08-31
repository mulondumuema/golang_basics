package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println()
	name := "Hello World"
	fmt.Println("String:", name)
	fmt.Println()

	fmt.Println("Bytes")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%x", name[i])
	}

	fmt.Println()

	fmt.Println("Index\tRune\tString")
	for x, y := range name {
		fmt.Println(x, "\t", y, "\t", string(y))
	}

	fmt.Println()

	fmt.Println()
	fmt.Println("Three ways to concatenate strings")

	h := "Hello"
	w := "World"

	myString := h + w
	fmt.Println(myString)
	fmt.Println()

	// using fmt (this is more efficient than the previous)
	myString = fmt.Sprintf("%s%s", h, w)
	fmt.Println(myString)
	fmt.Println()

	// using stringbuilder (very efficient)
	var sb strings.Builder
	sb.WriteString(h)
	sb.WriteString(w)
	fmt.Println(sb.String())

	fmt.Println()
	name = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println("Getting a substring")
	fmt.Println(name[10:13])

	// indexing a string

	courseName := "Learn Go For Beginners Crash Course"

	fmt.Println(string(courseName[0]))

	for i := 0; i <= 21; i++ {
		fmt.Println(string(courseName[i]))
	}
	fmt.Println()

	// string length
	fmt.Println("Length of CourseName is", len(courseName))

	var mySlice []string
	mySlice = append(mySlice, "one")
	mySlice = append(mySlice, "two")
	mySlice = append(mySlice, "three")

	fmt.Println("My slice has", len(mySlice), "elements")

	fmt.Println("the last element in mySlice is", mySlice[len(mySlice)-1])

	courses := []string {
		"Learn Go for Beginners Crash Course",
		"Learn Java for Beginners Crash Course",
		"Learn Python for Beginners Crash Course",
		"Learn C for Beginners Crash Course",
	}

	for _, x := range courses {
		if strings.Contains(x, "Go") {
			fmt.Println("Go is found in", x, "and index is", strings.Index(x, "Go"))
		}
	}

	newString := "Go is a great programming language. Go for it!"
	fmt.Println(strings.HasPrefix(newString, "Go"))
	fmt.Println(strings.HasPrefix(newString, "Python"))
	fmt.Println(strings.HasPrefix(newString, "!"))
	fmt.Println(strings.Count(newString, "Go"))
	fmt.Println(strings.Count(newString, "Fish"))
	fmt.Println(strings.Index(newString, "Go"))

	// string manipulation
	if strings.Contains(newString, "Go") {
		// newString = strings.Replace(newString, "Go", "Golang", 1) // To select all use -1 or use the ReplaceAll() function
		newString = strings.ReplaceAll(newString, "Go", "Golang")
	}

	fmt.Println(newString)

	// string comparison
	if "A" > "B" {
		fmt.Println("A is greater than B")
	} else {
		fmt.Println("A is not greater than B")
	}

	badEmail := " me@here.com "
	badEmail = strings.TrimSpace(badEmail)
	fmt.Printf("%s", badEmail)
	fmt.Println()

	myOtherString := "This is a clear example of why we search in one case only"

	searchString := strings.ToLower(myOtherString)

	if strings.Contains(searchString, "this") {
		fmt.Println("Found it!")
	} else {
		fmt.Println("Did not find it!")
	}

	// other case functions
	fmt.Println(strings.ToLower(myOtherString))
	fmt.Println(strings.ToUpper(myOtherString))
	fmt.Println(strings.Title(((myOtherString))))

	var myNewString = "Goodbye, cruel world!"
	fmt.Println(len(myNewString))
	fmt.Println(myNewString[3:5])
}