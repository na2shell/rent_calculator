package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		rent     = flag.Int("r", 70000, "int flag")
		init_pay = flag.Int("ini", 300000, "int flag")
	)

	flag.Parse()
	const support int = 27000
	const internet int = 2800

	var res = *rent + int(*init_pay/24) + internet - support
	fmt.Printf("result:%d\n", res)
}
