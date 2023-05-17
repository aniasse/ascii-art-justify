package ascii

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Ascii(banner string) map[byte][]string {
	ascii := make(map[byte][]string)
	var index byte = 32
	file, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal("Error : Not a ascci file in the repertory")
	}
	Split := strings.Split(string(file), "\n")
	for i := 1; i+8 < len(Split); i += 9 {
		ascii[index] = Split[i : i+8]
		index++
	}
	return ascii
}

func Ascii2(arg3 string) map[byte][]string {
	ascii := make(map[byte][]string)
	var index byte = 32
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
	} else {
		Split := strings.Split(string(file), "\n")
		for i := 1; i+8 < len(Split); i += 9 {
			ascii[index] = Split[i : i+8]
			index++
		}
	}
	return ascii
}

func Banner(s string) string {
	return "./file/" + s + ".txt"
}

func Flag(s string) bool {
	tab := []rune(s)
	if len(tab) >= 7 && string(tab[:7]) == "--color" {
		return true
	} else {
		return false
	}
}

func Output(s string) bool {
	tab := []rune(s)
	if len(tab) >= 8 && string(tab[:8]) == "--output" {
		return true
	} else {
		return false
	}
}

func Align(s string) bool {
	tab := []rune(s)
	if len(tab) >= 7 && string(tab[:7]) == "--align" {
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
	t2 := []rune{}
	if t[7] == '=' {
		t2 = t[8:]

	} else {
		res := "false"
		t2 = []rune(res)
	}
	return string(t2)
}
func taille(s string) int {
	tab := strings.Split(s, "\n")

	return len(tab[0])
}
func Taille_term() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	output, err := cmd.Output() // stocker le résultat de la commande
	if err != nil {
		return 0, 0, err
	}

	var width, height int
	_, err = fmt.Sscanf(string(output), "%d %d", &height, &width) // stocker les coordonnées dans les 2 variables
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func AjoutSpace1(nb int, s string) {
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
func AjoutSpace2(nb int, s string) {
	var space string
	for i := 0; i < nb; i++ {
		space += " "
	}
	str := strings.Split(s, "\n")
	for i := 0; i < len(str); i++ {
		str[i] = str[i] + space
	}
	fin := strings.Join(str, "\n")
	fmt.Println(fin)
}
func lenAscii(s string) int {
	str := strings.Split(s, "\n")
	return len(str[0])
}
func Left(s string) {
	largeur, _, err := Taille_term()
	if err != nil {
		panic(err)
	}
	length := taille(s)
	padding := largeur - length

	AjoutSpace2(padding, s)
}
func left(s string) {
	tab := strings.Split(s, "\n")

	for _, v := range tab {
		Left(v)
	}
}

func center(s string) {
	largeur, _, err := Taille_term()
	if err != nil {
		panic(err)
	}
	length := taille(s)
	padding := (largeur - length) / 2

	AjoutSpace1(padding, s)
}
func Center(s string) {
	tab := strings.Split(s, "\n")

	for _, v := range tab {
		center(v)
	}
}

func Right(s string) {
	largeur, _, err := Taille_term()
	if err != nil {
		panic(err)
	}
	length := taille(s)
	padding := largeur - length

	AjoutSpace1(padding, s)
}

func right(s string) {
	tab := strings.Split(s, "\n")
	for _, v := range tab {
		Right(v)
	}
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
func justify1(text string, width int) string {
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

func justify2(s string) {
	world := strings.Split(s, "\n")
	largeur, _, _ := Taille_term()
	for i := 0; i < len(world); i++ {
		world[i] = justify1(world[i], largeur)
	}
	justify := strings.Join(world, "\n")
	fmt.Println(justify)
}

func justify(s string) {
	tab := strings.Split(s, "\n")

	for _, v := range tab {
		justify2(v)
	}
}

func Alignment(s string) bool {
	if s[8:] != "left" && s[8:] != "right" && s[8:] != "justify" && s[8:] != "center" {
		return false
	}
	return true
}

func Extension(s string) bool {
	name := OutputName(s)
	if len(name) > 4 && name[len(name)-4:] == ".txt" {
		return true
	}
	return false
}
