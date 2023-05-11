package main

import (
	"fmt"
	"os"

	Func "ascii/functions"
)

func main() {
	if len(os.Args) == 2 && !Func.Flag(os.Args[1]) && !Func.Align(os.Args[1]) && !Func.Output(os.Args[1]) {
		Func.AsciiArt(os.Args[1])
	} else if len(os.Args) == 3 && !Func.Flag(os.Args[1]) && !Func.Align(os.Args[1]) && !Func.Output(os.Args[1]) {
		Func.AsciiFs(os.Args[1], os.Args[2])
	} else if len(os.Args) == 3 && Func.Flag(os.Args[1]) {
		Func.AsciiColor(os.Args[1], os.Args[2])
	} else if len(os.Args) == 4 && Func.Flag(os.Args[1]) {
		Func.AsciiTobeColored(os.Args[1], os.Args[3])
	} else if len(os.Args) == 4 && Func.Output(os.Args[1]) {
		Func.AsciiOutput(os.Args[1], os.Args[2], os.Args[3])
	} else if len(os.Args) == 3 && Func.Output(os.Args[1]) {
		Func.AsciiOutput1(os.Args[1], os.Args[2])
	} else if len(os.Args) == 4 && Func.Align(os.Args[1]) {
		Func.ASciiJystify(os.Args[1], os.Args[2], os.Args[3])
	} else if len(os.Args) == 3 && Func.Align(os.Args[1]) {
		Func.ASciiJystify1(os.Args[1], os.Args[2])
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
	}
}
