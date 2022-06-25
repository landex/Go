// This program use the two packages math and strings
package main

import ( //|-> lets you import multiple packages at once
	"math"    //|-> Import the "math" package so we can use math Floor.
	"strings" //|-> Import the "strings" package so we can sue strings Title
)

func main() {
	math.Floor(2.75)
	strings.Title("head first go math")
}
