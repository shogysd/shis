package main

import (
	"../shis"
	"fmt"
	"os"
	"strings"
	"time"
)

var versionString string = "shis: Version information not found"

func main() {

	errMessFormat := "err: %s [code:%d]\n"

	var err error
	var errCode int
	var words []string
	var histFileName = os.Getenv("HOME") + "/" + ".bash_history"
	var invertFlag = false
	var dupAllowFlag = false
	var maxRec = 100
	var andScanFlag = true
	var color = true
	var histories []string
	var sigTreeHead = shis.NewDataCV("")
	var histLines []shis.HistData
	var latestSigNo = -1

	// parse of arguments
	if err, errCode, words = shis.ArgParser(os.Args, &histFileName, &invertFlag, &dupAllowFlag, &maxRec, &andScanFlag, &color); err != nil {

		switch errCode {
		case -1:
			// error code '-1' mean asked for help by user
			shis.HelpPrinter()
			os.Exit(0)
		case -2:
			// error code '-2' mean asked for version
			fmt.Println(versionString)
			os.Exit(0)
		case 101:
			// error code '101' mean arguments parse err
			// commandline latest argument is '-v'
			fmt.Printf(errMessFormat, err, errCode)
			fmt.Printf("\nPlease see below for how to use the command.\n'shis -h'\n")
			os.Exit(1)
		}
	}

	latestSigNo = len(words)

	// make sig tree
	if err, errCode = shis.MakeTree(&words, sigTreeHead); err != nil {
		switch errCode {
		case 201:
			fmt.Printf(errMessFormat, err, errCode)
			os.Exit(1)
		}
	}

	// histry file read
	if err, errCode = shis.HistFileReader(&histories, &histFileName); err != nil {
		switch errCode {
		case 301:
			// error code '301' mean histFile not found
			// if use -f option you must to use file path for history file
			fmt.Printf(errMessFormat, err, errCode)
			os.Exit(1)
		}
	}

	// host scanner
	if err, errCode = shis.HistoryScanner(&histories, sigTreeHead, invertFlag, dupAllowFlag, &histLines, &maxRec, andScanFlag, latestSigNo); err != nil {
		switch errCode {
		case 301:
			// error code '301' mean histFile not found
			// if use -f option you must to use file path for history file
			fmt.Printf(errMessFormat, err, errCode)
			os.Exit(1)
		}
	}

	// hist printer
	i := len(histLines) - 1
	for i >= 0 {

		if histLines[i].Ts == -999999 {
			fmt.Printf(" No date information    ")
		} else {
			fmt.Printf(" %s   ", time.Unix(histLines[i].Ts, 0).Format("01-02-2006 (15:04:05)"))
			// fmt.Printf(" %s   ", time.Unix(histLines[i].Ts, 0).Format("Jan-_2-2006 (15:04:05)"))
		}

		if color == true {
			outStr := histLines[i].HistLineData
			for _, word := range words {
				outStr = strings.Replace(outStr, word, "\x1b[1;31m"+word+"\x1b[0;39m", -1)
			}
			fmt.Printf("%s\n", outStr)
		} else {
			fmt.Printf("%s\n", histLines[i].HistLineData)
		}

		i -= 1

	}

	os.Exit(0)
}
