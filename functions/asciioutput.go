package ascii

import (
	"fmt"
	"os"
	"strings"
)

func AsciiOutput(arg1 string, arg2 string, arg3 string) {
	ascii := make(map[byte][]string)
	var index byte = 32
	banner := Banner(arg3)
	file, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error : Not a ascci file in the repertory")
	}
	if arg3 == "thinkertoy" {
		Split := strings.Split(string(file), "\r\n")
		for i := 1; i+8 < len(Split); i += 9 {
			ascii[index] = Split[i : i+8]
			index++
		}
	} else {
		Split := strings.Split(string(file), "\n")
		for i := 1; i+8 < len(Split); i += 9 {
			ascii[index] = Split[i : i+8]
			index++
		}
	}

	tabascii := ascii

	if len(arg2) != 0 {
		split := strings.Split(arg2, "\\n")
		if NewLine(split) {
			split = split[:len(split)-1]
		}
		var res string
		for _, v := range split {
			tabrune := []rune(v)
			if Printable(tabrune) {
				for j := 0; j < 8; j++ {
					for i := 0; i < len(tabrune); i++ {
						res += Print(tabrune[i], j, tabascii)
					}
					if len(tabrune) != 0 {
						res += "\n"
					} else {
						break
					}
				}
				res += "\n"
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
}
