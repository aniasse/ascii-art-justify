package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiOutput(arg1 string, arg2 string, arg3 string) {
	tabascii := Ascii2(arg3)
	if Extension(arg1) {
		if len(arg2) != 0 {
			split := strings.Split(arg2, "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			var res string
			for k, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							res += Print(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							res += "\n"
						} else {
							res += "\n"
							break
						}
					}
					if k == len(split)-1 {
						res += "\n"
					}
					Filename := OutputName(arg1)
					sortie, err := os.Create(Filename)
					if err != nil {
						fmt.Println(err)
						return
					}
					_, err = sortie.WriteString(res)
					if err != nil {
						fmt.Println(err)
						sortie.Close()
						return
					}
					err = sortie.Close()
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
				}
			}
		}
	} else {
		fmt.Println("Error: invalid extension!")
	}
}

func AsciiOutput1(arg1 string, arg2 string) {
	banner := "./file/standard.txt"
	tabascii := Ascii(banner)
	if Extension(arg1) {
		if len(arg2) != 0 {
			split := strings.Split(arg2, "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			var res string
			for k, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							res += Print(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							res += "\n"
						} else {
							res += "\n"
							break
						}
					}
					if k == len(split)-1 {
						res += "\n"
					}

					Filename := OutputName(arg1)
					sortie, err := os.Create(Filename)
					if err != nil {
						fmt.Println(err)
						return
					}
					_, err = sortie.WriteString(res)
					if err != nil {
						fmt.Println(err)
						sortie.Close()
						return
					}
					err = sortie.Close()
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
				}
			}
		}
	} else {
		fmt.Println("Error: invalid extension!")
	}
}
