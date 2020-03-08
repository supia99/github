// https://go-tour-jp.appspot.com/flowcontrol/8

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannnot Sqrt negative number:%v\n", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x <= 0 {
		//var e ErrNegativeSqrt = ErrNegativeSqrt(x)
		return 0, ErrNegativeSqrt(x)
		//ErrNegativeSqrt(x)
	}
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// func main() {
// 	fmt.Println(Sqrt(2))
// 	fmt.Println(Sqrt(-2))
// }
