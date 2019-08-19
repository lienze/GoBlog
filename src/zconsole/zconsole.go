package zconsole

import "fmt"

func PrintRedText(sInfo string) {
	fmt.Printf("%c[1;40;31m%s%c[0m\n", 0x1B, sInfo, 0x1B)
}

func PrintYellowText(sInfo string) {
	fmt.Printf("%c[1;40;33m%s%c[0m\n", 0x1B, sInfo, 0x1B)
}

func PrintGreenText(sInfo string) {
	fmt.Printf("%c[1;40;32m%s%c[0m\n", 0x1B, sInfo, 0x1B)
}
