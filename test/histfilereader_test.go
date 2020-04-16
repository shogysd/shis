package test

import (
	"../shis"
	"testing"
)

func TestHistFileReader(t *testing.T) {

	var err error
	var errCode int
	var histories = []string{}
	var histFileName = ""

	histFileName = "./testsamples/sample_history_file"
	err, errCode = shis.HistFileReader(&histories, &histFileName)
	if err != nil || errCode != 0 {
		t.Error("file open fail")
	}

	histFileName = "./testsamples/sample_history_file_ZZZ"
	err, errCode = shis.HistFileReader(&histories, &histFileName)
	if err == nil || errCode == 0 {
		t.Error("file open fail ()")
	}

	histories = []string{}
	histFileName = "./testsamples/sample_history_file"
	err, errCode = shis.HistFileReader(&histories, &histFileName)
	if err != nil || errCode != 0 {
		t.Error("file open fail")
	}
	var hisAnswer = []string{"#0", "command line 01", "#1", "command line 02", "#3", "command line 03"}
	for i, readLine := range histories {
		if readLine != hisAnswer[i] {
			t.Error("read file test err")
			break
		}
	}
}
