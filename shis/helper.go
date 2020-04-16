package shis

import "fmt"

func HelpPrinter() {

	message := `
shis [OPTION] PATTERN (In no particular order)  ......

options [ -v | -d | -n [int] | -f | -h | --version ]

-v: invert scan
	Those that do NOT match the keywords are output

-d: Allow duplication
	Output duplicate content

-n [int]: number of outputs
	Default is 100 (Duplicates are eliminated)

-f: bash history file name
	Default is "~/.bash_history"
	Please specify HISTFILE

-h: print help message (this message)

--version: print version
`
	fmt.Println(message)

}
