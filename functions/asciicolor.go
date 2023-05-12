package ascii

import (
	"fmt"
	"strings"
)

func AsciiColor(arg1 string, arg2 string) {
	banner := "./file/standard.txt"
	tabascii := Ascii(banner)
	if len(arg2) != 0 {
		split := strings.Split(arg2, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) && Supported(ColorFlag(arg1)) {
				for j := 0; j < 8; j++ {
					for i := 0; i < len(tabrune); i++ {
						MatchColored(tabrune[i], j, tabascii)
					}
					if len(tabrune) != 0 {
						fmt.Println()
					} else {
						fmt.Println()
						break
					}
				}
			} else {
				fmt.Println("Error : Non-displayable character or Color not supported!!!")
			}
		}
	}
}

func AsciiTobeColored(arg1 string, arg3 string) {
	banner := "./file/standard.txt"
	tabascii := Ascii(banner)
	if len(arg3) != 0 {
		split := strings.Split(arg3, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) && Supported(ColorFlag(arg1)) {
				for j := 0; j < 8; j++ {
					for i := 0; i < len(tabrune); i++ {
						ToBeColored(tabrune[i], j, tabascii)
					}
					if len(tabrune) != 0 {
						fmt.Println()
					} else {
						fmt.Println()
						break
					}
				}
			} else {
				fmt.Println("Error : Non-displayable character or Color not supported!!!")
			}
		}
	}
}
