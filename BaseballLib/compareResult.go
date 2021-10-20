package BaseballLib

import "github.com/jinzza456/baseball_project/BaseballLib"

var (
	line1 = [3]int{1, 2, 3}
	line2 = [3]int{4, 5, 6}
	line3 = [3]int{7, 8, 9}
)

func InCaseHit(swing, pitch int) {
	if pitch == swing {
		BaseballLib.HitCall()
		BaseballLib.AddBaseCount()
		IfBaseCountIsFour()
	}
}

func InCaseStrike() {
	BaseballLib.StrikeCall()
	BaseballLib.AddStrikeCount()
	IfFoulCountIsTwo()
}

func InCaseSwingLine1(pitch int) {
	for i := range line1 {
		if line1[i] == pitch {
			BaseballLib.FoulCall()
			BaseballLib.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
	}
}

func InCaseSwingLine2(pitch int) {
	for i := range line2 {
		if line2[i] == pitch {
			BaseballLib.FoulCall()
			BaseballLib.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
	}
}

func InCaseSwingLine3(pitch int) {
	for i := range line3 {
		if line3[i] == pitch {
			BaseballLib.FoulCall()
			BaseballLib.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
	}
}

func CompareAllCase(pitch, swing int) {
	if swing < 4 {
		InCaseSwingLine1(pitch)
	} else {
		InCaseStrike()
	}
	if swing > 3 && swing < 7 {
		InCaseSwingLine2(pitch)
	} else {
		InCaseStrike()
	}
	if swing > 6 {
		InCaseSwingLine3(pitch)
	} else {
		InCaseStrike()
	}
}
