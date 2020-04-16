package shis

import (
	"errors"
	"strconv"
)

func ArgParser(args []string, histFileName *string, invertFlag *bool, dupAllowFlag *bool, maxRec *int, andScanFlag *bool, color *bool) (error, int, []string) {

	var words []string
	var andScanReaded = false

	if len(args) == 1 {
		return errors.New("asked for help"), -1, nil
	} else {
		i := 1
		for i < len(args) {

			if string(args[i][0]) == "-" {

				switch args[i] {

				case "-v":
					*invertFlag = true

				case "-d":
					*dupAllowFlag = true

				case "-n":
					if i < len(args)-1 {
						var err error
						if *maxRec, err = strconv.Atoi(args[i+1]); err != nil {
							return errors.New("arguments parse err"), 101, nil
						}
						i += 1
					} else {
						return errors.New("arguments parse err"), 101, nil
					}

				case "-c":
					if i < len(args)-1 {
						if args[i+1] == "true" || args[i+1] == "t" {
							*color = true
						} else if args[i+1] == "false" || args[i+1] == "f" {
							*color = false
						} else {
							return errors.New("arguments parse err"), 101, nil
						}
						i += 1
					} else {
						return errors.New("arguments parse err"), 101, nil
					}

				case "-f":
					if i < len(args)-1 {
						*histFileName = args[i+1]
						i += 1
					} else {
						return errors.New("arguments parse err"), 101, nil
					}

				case "-a":
					if andScanReaded == false {
						andScanReaded = true
						*andScanFlag = true
					} else {
						return errors.New("arguments parse err"), 101, nil
					}

				case "-o":
					if andScanReaded == false {
						andScanReaded = true
						*andScanFlag = false
					} else {
						return errors.New("arguments parse err"), 101, nil
					}

				case "-h":
					return errors.New("asked for help"), -1, nil

				case "--version":
					return errors.New("require version print"), -2, nil

				default:
					return errors.New(string(args[i]) + " arguments parse err"), 101, nil
				}

			} else {
				words = append(words, args[i])
			}
			i += 1

		}
		return nil, 0, words
	}
}
