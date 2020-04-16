package shis

import (
	"os"
	"strconv"
	"strings"
)

type HistData struct {
	Ts           int64
	HistLineData string
}

func NewHist() *HistData {
	cont := new(HistData)
	cont.Ts = -999999
	cont.HistLineData = ""
	return cont
}

func HistoryScanner(
	histories *[]string, sigTreeHead *DataCV, invertFlag bool, dupAllowFlag bool, histLines *[]HistData, maxRec *int, andScanFlag bool, latestSigNo int) (error, int) {

	// history scan
	i := len(*histories) - 1
	recCounter := 0
	insertHist := NewHist()
	for i >= 0 {
		if (*histories)[i] != "" && string((*histories)[i][0]) != "#" {
			// hitFlag := false

			searchLine := strings.TrimRight((*histories)[i], " ")

			if err, errCode, hitFlag := SearchTree(&(searchLine), sigTreeHead, andScanFlag, latestSigNo); err != nil {
				switch errCode {
				case 401:
					// return errors.New("arguments parse err"), 401
				}
			} else if hitFlag == !invertFlag {

				if i >= 2 && string((*histories)[i-1][0]) == "#" {
					if ts, err := strconv.ParseInt((*histories)[i-1][1:], 10, 64); err != nil {
						// timestamo parse err
						os.Exit(1)
					} else {
						insertHist.Ts = ts
						insertHist.HistLineData = searchLine
						// insertHist.HistLineData = (*histories)[i]
					}
				} else {
					insertHist.Ts = -999999
					insertHist.HistLineData = searchLine
					// insertHist.HistLineData = (*histories)[i]
				}

				hitLineDupFlag := false // Match search criteria

				if dupAllowFlag == false {

					for _, hitLine := range *histLines {
						if hitLine.HistLineData == insertHist.HistLineData {
							hitLineDupFlag = true
							break
						}
					}

					if hitLineDupFlag == false {
						*histLines = append(*histLines, *insertHist)
					}

				} else {
					*histLines = append(*histLines, *insertHist)
				}
				recCounter += 1
				if recCounter >= *maxRec {
					break
				}
			}
		}
		i -= 1
	}
	return nil, 0
}
