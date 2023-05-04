package ascii

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Banner(s string) string {
	return "./file/" + s + ".txt"
}

func Flag(s string) bool {
	tab := []rune(s)
	if len(tab) > 8 && string(tab[:8]) == "--color=" {
		return true
	} else {
		return false
	}
}

func Output(s string) bool {
	tab := []rune(s)
	if len(tab) > 9 && string(tab[:9]) == "--output=" {
		return true
	} else {
		return false
	}
}

func Align(s string) bool {
	tab := []rune(s)
	if len(tab) > 8 && string(tab[:8]) == "--align=" {
		return true
	} else {
		return false
	}
}

func ColorFlag(S string) string {
	var color string
	var j int
	for i := 0; i < len(S); i++ {
		if S[i] != '=' {
			j++
		} else {
			break
		}
	}
	for k := j + 1; k < len(S); k++ {
		color += string(S[k])
	}
	return color
}

func ToColor(s string, r rune) bool {
	tab := []rune(s)
	for i := 0; i < len(tab); i++ {
		if tab[i] == r {
			return true
		}
	}
	return false
}

func OutputName(S string) string {
	var filename string
	var j int
	for i := 0; i < len(S); i++ {
		if S[i] != '=' {
			j++
		} else {
			break
		}
	}
	for k := j + 1; k < len(S); k++ {
		filename += string(S[k])
	}
	return filename
}

func Print(r rune, i int, ascii map[byte][]string) string {
	var result string
	for ind, v := range ascii {
		if rune(ind) == r {
			result = v[i]
			break
		}
	}
	return result
}

func Write(r rune, i int, ascii map[byte][]string) string {
	var str string
	for ind, v := range ascii {
		if rune(ind) == r {
			str += v[i]
		}
	}
	return str
}

func Match(r rune, i int, ascii map[byte][]string) {
	for ind, v := range ascii {
		if rune(ind) == r {
			fmt.Print(v[i])
		}
	}
}
func MatchColored(r rune, i int, ascii map[byte][]string) {
	couleur := ColorFlag(os.Args[1])
	for ind, v := range ascii {
		if rune(ind) == r {
			color(couleur, v[i])
		}
	}
}

func ToBeColored(r rune, i int, ascii map[byte][]string) {
	couleur := ColorFlag(os.Args[1])
	for ind, v := range ascii {
		if rune(ind) == r && ToColor(os.Args[2], r) {
			color(couleur, v[i])
		}
		if rune(ind) == r && !ToColor(os.Args[2], r) {
			fmt.Print(v[i])
		}
	}
}

func NewLine(tab []string) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] != "" {
			return false
		}
	}
	return true
}

func Printable(tab []rune) bool {
	for i := 0; i < len(tab); i++ {
		if tab[i] < 32 || tab[i] > 126 {
			return false
		}
	}
	return true
}

func opt(s string) string {
	t := []rune(s)
	t2 := t[8:]
	return string(t2)
}
func taille(s string) int {
	tab := strings.Split(s, "\n")

	return len(tab[0])
}
func AjoutSpace(nb int, s string) {
	var space string
	for i := 0; i < nb; i++ {
		space += " "
	}
	str := strings.Split(s, "\n")
	for i := 0; i < len(str); i++ {
		str[i] = space + str[i]
	}
	fin := strings.Join(str, "\n")

	fmt.Println(fin)
}
func lenAscii(s string) int {
	str := strings.Split(s, "\n")
	return len(str[0])
}
func left(s string) {
	fmt.Println(s)
}

func Center(s string) {
	largeur, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	length := taille(s)
	padding := (largeur - length) / 2

	AjoutSpace(padding, s)
}

func Right(s string) {
	largeur, _, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	length := taille(s)
	padding := largeur - length

	AjoutSpace(padding, s)
}

func dollars(s string) string {
	tab1 := []rune(s)
	tab2 := []rune{}

	for _, v := range tab1 {
		if v == '$' {
			tab2 = append(tab2, v)
		} else {
			continue
		}
	}
	return string(tab2)
}
func justify(text string, width int) string {
	words := strings.Split(text, "$$$$$$")
	concat := strings.Join(words, "")
	numWords := len(words)

	if numWords == 1 {
		return text
	}
	numSpaces := width - len(concat)
	spaceWidth := numSpaces / (numWords - 1)
	reste1 := numSpaces % (numWords - 1)
	spaces := strings.Repeat(" ", spaceWidth)
	result := ""
	for i, word := range words {
		result += word
		if i < numWords-1 {
			result += spaces
			if i < reste1 {
				result += " "
			}
		}
	}
	return result
}

func Justify(s string) {
	world := strings.Split(s, "\n")
	largeur, _, _ := term.GetSize(0)
	for i := 0; i < len(world); i++ {
		world[i] = justify(world[i], largeur)
	}
	justify := strings.Join(world, "\n")
	fmt.Println(justify)
}
