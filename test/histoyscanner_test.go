package test

import (
	"../shis"
	"testing"
)

func TestHistoryScanner(t *testing.T) {

	var err error
	var errCode int
	var words []string
	var invertFlag = false
	var dupAllowFlag = false
	var maxRec = 100
	var andScanFlag = true
	var histories []string
	var sigTreeHead = shis.NewDataCV("")
	var histLines []shis.HistData
	var latestSigNo = -1

	words = []string{"01", "02"}
	latestSigNo = len(words)
	histories = []string{"history line 01", "history line 02", "history line 03"}

	if err, errCode = shis.MakeTree(&words, sigTreeHead); err != nil || errCode != 0 {
		t.Error("sig tree err")
	}

	if err, errCode = shis.HistoryScanner(&histories, sigTreeHead, invertFlag, dupAllowFlag, &histLines, &maxRec, andScanFlag, latestSigNo); err != nil || errCode != 0 {
		t.Error("scan history scan err")
	}

}
