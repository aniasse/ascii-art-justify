package ascii

import (
	"fmt"
	"strings"
)

func AsciiFs(arg1 string, arg2 string) {
	tabascii := Ascii2(arg2)
	if len(arg1) != 0 {
		split := strings.Split(arg1, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) {
				for j := 0; j < 8; j++ {
					for i := 0; i < len(tabrune); i++ {
						Match(tabrune[i], j, tabascii)
					}
					if len(tabrune) != 0 {
						fmt.Println()
					} else {
						fmt.Println()
						break
					}
				}
			} else {
				fmt.Println("Error : Non-displayable character !!!")
			}
		}
	}
}
