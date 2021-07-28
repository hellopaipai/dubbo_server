package hash

import "fmt"

func New(a,b int) (ret int) {
	ret = a + b
	fmt.Println("这是2.0.0版本")
	return ret
}
