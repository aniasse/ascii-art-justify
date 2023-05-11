package ascii

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ASciiJystify(arg1 string, arg2 string, arg3 string) {
	ascii := make(map[byte][]string)
	var index byte = 32
	// if len(arg1) > 8 {
	option := opt(arg1)
	banner := Banner(arg3)

	file, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal("Error : Not a ascci file in the repertory")
	}
	if arg3 == "thinkertoy" {
		Split := strings.Split(string(file), "\r\n")
		for i := 1; i+8 < len(Split); i += 9 {
			ascii[index] = Split[i : i+8]
			index++
		}
		if option == "justify" {
			ascii[32] = []string{"$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$"}
		}
	} else {
		Split := strings.Split(string(file), "\n")
		for i := 1; i+8 < len(Split); i += 9 {
			ascii[index] = Split[i : i+8]
			index++
		}
		if option == "justify" {
			ascii[32] = []string{"$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$"}
		}
	}
	tabascii := ascii
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
				largeur, _, _:= Taille_term()
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
	ascii := make(map[byte][]string)
	var index byte = 32
	option := opt(arg1)

	file, err := os.ReadFile("./file/standard.txt")
	if err != nil {
		log.Fatal("Error : Not a ascci file in the repertory")
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	if option == "justify" {
		ascii[32] = []string{"$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$", "$$$$$$"}
	}

	tabascii := ascii
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
				largeur, _, _:= Taille_term()
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
