package main

import (
	"fmt"
	"github.com/xian-zx/quanban"
)

func main() {
	// 半角字符 => 全角字符
	fmt.Println(quanban.B2Q(BanStr))
	// 全角字符 => 半角字符
	fmt.Println(quanban.Q2B(QuanStr))
}
