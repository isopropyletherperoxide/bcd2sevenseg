package main

import (
	"fmt"
	"math"
	"os"
)


func font_init() [10]string{

var font [10]string
	font[0] = " ### \n" +
		"#   #\n" +
		"#   #\n" +
		"     \n" +
		"#   #\n" +
		"#   #\n" +
		" ### "

	font[1] = "     \n" +
		"    #\n" +
		"    #\n" +
		"     \n" +
		"    #\n" +
		"    #\n" +
		"     "

	font[2] = " ### \n" +
		"    #\n" +
		"    #\n" +
		" ### \n" +
		"#    \n" +
		"#    \n" +
		" ### "

	font[3] = " ### \n" +
		"    #\n" +
		"    #\n" +
		" ### \n" +
		"    #\n" +
		"    #\n" +
		" ### "

	font[4] = "     \n" +
		"#   #\n" +
		"#   #\n" +
		" ### \n" +
		"    #\n" +
		"    #\n" +
		"     "

	font[5] = " ### \n" +
		"#    \n" +
		"#    \n" +
		" ### \n" +
		"    #\n" +
		"    #\n" +
		" ### "

	font[6] = " ### \n" +
		"#    \n" +
		"#    \n" +
		" ### \n" +
		"#   #\n" +
		"#   #\n" +
		" ### "

	font[7] = " ### \n" +
		"    #\n" +
		"    #\n" +
		"     \n" +
		"    #\n" +
		"    #\n" +
		"     "

	font[8] = " ### \n" +
		"#   #\n" +
		"#   #\n" +
		" ### \n" +
		"#   #\n" +
		"#   #\n" +
		" ### "

	font[9] = " ### \n" +
		"#   #\n" +
		"#   #\n" +
		" ### \n" +
		"    #\n" +
		"    #\n" +
		" ### "
return font;
}

func bcdconv(in string) int {
	var out int
	for i := 3; i >= 0; i-- {
		num := float64(in[i] - '0')
		out += int(num * math.Pow(2, float64(i)))
	}
	return out
}

func main() {
        font := font_init()
        letter := bcdconv(os.Args[1])
        if letter > 9 {
                println(letter)
        } else {
	fmt.Println(font[letter])
        }
}
