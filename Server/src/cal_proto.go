package main

import "fmt"

func main() {
	const rent int = 70000
	const init_pay int = 300000
	const support int = 27000
	const internet int = 2800

	var res = rent + int(init_pay/24) + internet - support
	fmt.Printf("result:%d\n", res)
}
