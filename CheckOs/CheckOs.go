package CheckOs

import (
	"fmt"
	"runtime"
)

func CheckOs() {
	os := runtime.GOOS
	fmt.Println("Your os is: ", os)
}
