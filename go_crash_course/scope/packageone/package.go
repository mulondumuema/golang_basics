package packageone

import "fmt"

var PackageVar = "I am a package variable"

func PrintMe(myVar, blockVar string)  {
	fmt.Println(myVar, blockVar, PackageVar)
	
}

