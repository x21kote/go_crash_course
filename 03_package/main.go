package main

import (
	"fmt"
	"math"

	"github.com/bradtraversy/go_crash_course/03_package/strutil"
)

func main() {
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(strutil.Reverse("hello"))
}
