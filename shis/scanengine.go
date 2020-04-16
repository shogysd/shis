package shis

type DataCV struct {
	// data is single character string, it's data
	Data string
	// termination flag
	Term int
	// next node (this flag is used only when end flag ia false)
	NextNode *DataCV
	// parallel node (this flag is used only when end flag ia false)
	ParallelNode *DataCV
	// jump node
	JumpNode *DataCV
}

func NewDataCV(data string) *DataCV {
	cont := new(DataCV)
	cont.Data = data
	cont.Term = -1
	cont.NextNode = nil
	cont.ParallelNode = nil
	cont.JumpNode = nil
	return cont
}

func addSig(sigs *[]string, sigTreeHead *DataCV) (error, int) {

	for sigNo, sig := range *sigs {
		pNode := sigTreeHead
		eNode := NewDataCV("")
		for i, char := range sig {
			nNode := NewDataCV("")
			if pNode.NextNode == nil {
				nNode = NewDataCV(string(char))
				pNode.NextNode = nNode
				pNode = pNode.NextNode
			} else {
				eNode = pNode.NextNode
				for {
					if eNode.Data == string(char) {
						pNode = eNode
						break
					} else if eNode.ParallelNode != nil {
						eNode = eNode.ParallelNode
					} else {
						nNode = NewDataCV(string(char))
						eNode.ParallelNode = nNode
						pNode = eNode.ParallelNode
						break
					}
				}
			}
			if i == len(sig)-1 {
				pNode.Term = sigNo
			}
		}
	}
	return nil, 0
}

func addJumpRoad(sigs *[]string, sigTreeHead *DataCV) (error, int) {
	// make tree (add jump load)
	for _, sig := range *sigs {
		pNode := sigTreeHead
		eNode := NewDataCV("")
		jNode := NewDataCV("")

		for _, char := range sig {
			eNode = pNode.NextNode

			if pNode == sigTreeHead {
				for {
					if eNode.Data == string(char) {
						eNode.JumpNode = sigTreeHead
						pNode = eNode
						break
					} else if eNode.ParallelNode != nil {
						eNode = eNode.ParallelNode
					}
				}
			} else {
				jNode = pNode.JumpNode.NextNode
				for {
					if eNode.Data == string(char) {
						if jNode.Data == eNode.Data { // == string(char)
							eNode.JumpNode = jNode
							pNode = eNode
							break
						} else if jNode.ParallelNode != nil {
							jNode = jNode.ParallelNode
						} else {
							eNode.JumpNode = sigTreeHead
							pNode = eNode
							break
						}
					} else if eNode.ParallelNode != nil {
						eNode = eNode.ParallelNode
					} else {
						eNode.JumpNode = sigTreeHead
						pNode = eNode.NextNode
						break
					}
				}
			}
		}
	}
	return nil, 0
}

func moreSig(sig *string, sigTreeHead *DataCV, sigNo int) (error, int) {

	pNode := sigTreeHead
	eNode := NewDataCV("")

	for i, char := range *sig {
		nNode := NewDataCV("")
		if pNode.NextNode == nil {
			nNode = NewDataCV(string(char))
			pNode.NextNode = nNode
			pNode = pNode.NextNode
		} else {
			eNode = pNode.NextNode
			for {
				if eNode.Data == string(char) {
					pNode = eNode
					break
				} else if eNode.ParallelNode != nil {
					eNode = eNode.ParallelNode
				} else {
					nNode = NewDataCV(string(char))
					eNode.ParallelNode = nNode
					pNode = eNode.ParallelNode
					break
				}
			}
		}
		if i == len(*sig)-1 {
			pNode.Term = sigNo
		}
	}

	return nil, 0
}

func moreJumpRoad(sig *string, sigTreeHead *DataCV) (error, int) {

	pNode := sigTreeHead
	eNode := NewDataCV("")
	jNode := NewDataCV("")

	for _, char := range *sig {
		eNode = pNode.NextNode

		if pNode == sigTreeHead {
			for {
				if eNode.Data == string(char) {
					eNode.JumpNode = sigTreeHead
					pNode = eNode
					break
				} else if eNode.ParallelNode != nil {
					eNode = eNode.ParallelNode
				}
			}
		} else {
			jNode = pNode.JumpNode.NextNode
			for {
				if eNode.Data == string(char) {
					if jNode.Data == eNode.Data { // == string(char)
						eNode.JumpNode = jNode
						pNode = eNode
						break
					} else if jNode.ParallelNode != nil {
						jNode = jNode.ParallelNode
					} else {
						eNode.JumpNode = sigTreeHead
						pNode = eNode
						break
					}
				} else if eNode.ParallelNode != nil {
					eNode = eNode.ParallelNode
				} else {
					eNode.JumpNode = sigTreeHead
					pNode = eNode.NextNode
					break
				}
			}
		}
	}

	return nil, 0
}

func MakeTree(words *[]string, sigTreeHead *DataCV) (error, int) {

	addSig(words, sigTreeHead)
	addJumpRoad(words, sigTreeHead)

	return nil, 0
}

func AddTree(word *string, sigTreeHead *DataCV, latestSigNo int) (error, int) {

	latestSigNo += 1
	moreSig(word, sigTreeHead, latestSigNo)
	moreJumpRoad(word, sigTreeHead)

	return nil, 0
}

func flagsChecker(flags []bool) bool {
	for _, f := range flags {
		if f == false {
			return false
		}
	}
	return true
}

func SearchTree(word *string, sigTreeHead *DataCV, andScanFlag bool, latestSigNo int) (error, int, bool) {
	/*
	* ret arg
	* ( error, int(err code), bool(hit or not hit) )
	 */

	var tapFlags = make([]bool, latestSigNo)

	if sigTreeHead.NextNode == nil {
		// tree is not constructed
		return nil, 0, false
	}

	pNode := sigTreeHead
	eNode := NewDataCV("")

	for _, char := range *word {
		eNode = pNode.NextNode

		for {
			if eNode.Data == string(char) {
				if eNode.Term != -1 {
					tapFlags[eNode.Term] = true
					if !andScanFlag || (andScanFlag && flagsChecker(tapFlags)) {
						return nil, 0, true
					} else {
						break
					}
				} else if eNode.JumpNode != nil && eNode.JumpNode.Term != -1 {
					tapFlags[eNode.JumpNode.Term] = true
					if !andScanFlag || (andScanFlag && flagsChecker(tapFlags)) {
						return nil, 0, true
					} else {
						break
					}
				} else {
					pNode = eNode
					break
				}
			} else if eNode.ParallelNode != nil {
				eNode = eNode.ParallelNode
			} else if pNode.JumpNode != nil {
				pNode = pNode.JumpNode
				eNode = pNode.NextNode
			} else {
				pNode = sigTreeHead
				break
			}
		}
	}
	// searched till the end but did not hit
	return nil, 0, false
}
