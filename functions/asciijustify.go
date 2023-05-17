package ascii

import (
	"fmt"
	"os"
	"strings"
)

func ASciiJystify(arg1 string, arg2 string, arg3 string) {
	tabascii := Ascii2(arg3)
	option := opt(arg1)
	if option == "justify" {
		tabascii[32] = []string{"$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$"}
	}
	if option == "false" {
		fmt.Println("Error: invalid syntax")
	}
	var affiche string
	if Alignment(arg1) {
		if len(arg2) != 0 {
			var split []string
			split = strings.Split(os.Args[2], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							affiche += Write(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							affiche += "\n"
						} else {
							affiche += "\n"
							break
						}
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
					return
				}

			}
		}

		if option == "left" {
			left(affiche)
		}
		if option == "right" {
			right(affiche)
		}
		if option == "center" {
			Center(affiche)
		}
		if option == "justify" {
			if len(strings.Split(os.Args[2], " ")) == 1 {
				left(affiche)
			} else {
				largeur, _, _ := Taille_term()
				if lenAscii(affiche) > largeur {
					fmt.Println("Ce texte ne peut pas etre justifie car son ASCII depasse la taille du terminal ")
					return
				}
				justify(affiche)
			}
		}
	} else {
		fmt.Println("Error: invalid alignment!")
	}
}
func ASciiJystify1(arg1 string, arg2 string) {
	banner := "./file/standard.txt"
	tabascii := Ascii(banner)
	option := opt(arg1)
	if option == "justify" {
		tabascii[32] = []string{"$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$"}
	}
	if option == "false" {
		fmt.Println("Error: invalid syntax")
	}
	var affiche string
	if Alignment(arg1) {

		if len(arg2) != 0 {
			var split []string
			split = strings.Split(os.Args[2], "\\n")
			if NewLine(split) {
				split = split[:len(split)-1]
			}
			for _, v := range split {
				tabrune := []rune(v)
				if Printable(tabrune) {
					for j := 0; j < 8; j++ {
						for i := 0; i < len(tabrune); i++ {
							affiche += Write(tabrune[i], j, tabascii)
						}
						if len(tabrune) != 0 {
							affiche += "\n"
						} else {
							affiche += "\n"
							break
						}
					}
				} else {
					fmt.Println("Error : Non-displayable character !!!")
					return
				}

			}
		}

		if option == "left" {
			left(affiche)
		}
		if option == "right" {
			right(affiche)
		}
		if option == "center" {
			Center(affiche)
		}
		if option == "justify" {
			if len(strings.Split(os.Args[2], " ")) == 1 {
				left(affiche)
			} else {
				largeur, _, _ := Taille_term()
				if lenAscii(affiche) > largeur {
					fmt.Println("Ce texte ne peut pas etre justifie car son ASCII depasse la taille du terminal ")
					return
				}
				justify(affiche)
			}
		}

	} else {
		fmt.Println("Error: invalid alignment!")
	}
}
