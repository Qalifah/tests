package iteration  

import (

)

// Repeat outputs a string 5 times
func Repeat(arg string, times int) string {
	var holder string
	for i := 1; i <= times; i++ {
		holder += arg
	}
	return holder
}