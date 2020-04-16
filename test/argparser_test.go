package test

import (
	"../shis"
	"testing"
)

func TestArgParser(t *testing.T) {

	var err error
	var errCode int
	var words []string
	var args = []string{}
	var histFileName = "testHome"
	var invertFlag = false
	var dupAllowFlag = false
	var maxRec = 100
	var andScanFlag = true
	var color = true
	/*
		var histories []string
		var sigTreeHead = shis.NewDataCV("")
		var histLines []shis.HistData
		var latestSigNo = -1
	*/

	args = []string{"arg1", "arg2", "arg3"}
	if err, errCode, words = shis.ArgParser(args, &histFileName, &invertFlag, &dupAllowFlag, &maxRec, &andScanFlag, &color); err != nil {
		t.Error("target list make err")
	} else {

		if errCode != 0 {
			t.Error("target list make err")
		}

		for i, word := range words {
			if word != words[i] {
				t.Error("target list make err")
				break
			}
		}

	}

}
