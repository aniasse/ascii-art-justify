
# Ascii-Art-Justify

This program aims to convert user input into ASCII art.


## Installation

Clone this repository to your local machine.
Make sure Go is installed on your machine.
Navigate to the directory where you cloned the repository and open a terminal.
Run the command "go run ." to start the program.
```bash
  go run .
```
    
## Requirement

The program requires two arguments: 

a. A string to convert to ASCII art.

b. The name of the template.

Optional: You can add a third argument to specify the alignment of ASCII art:

a. "left" to align the text to the left.

b. "right" to align the text to the right.

c. "center" to align the text in the center.

d. "justify" to justify the text.

## Usage/Examples

Usage: go run . [OPTION] [STRING] [BANNER]

Example: 
```bash
  go run . --align=right something standard
```


## Note

The program will not display non-displayable characters.
If a non-existent ASCII file is specified, the program will display an error message and terminate.
If you want to add your own template, just create a new file in the "file" folder and follow the same format as the existing template.
## Authors

- [@fatidiop](https://learn.zone01dakar.sn/git/fatidiop)
- [@aboubakdiallo](https://learn.zone01dakar.sn/git/aboubakdiallo)
- [@aniasse](https://learn.zone01dakar.sn/git/aniasse)

