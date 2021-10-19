package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 오타 , 카멜, 함수의 기능은 하나만 ,함수와 매서드 중 적합한걸 사용, 매서드명 및 함수명 이 의미있게
// 상수는 메모리의 로스를 줄이기 위해 const 나 var 로 뺴놓는게 좋음
var (
	players = make([]string, 9, 9)

	line1 = [3]int{1, 2, 3}
	line2 = [3]int{4, 5, 6}
	line3 = [3]int{7, 8, 9}

	c *ProgressOfGameCount
)

type ProgressOfGameCount struct {
	baseCount   int
	foulCount   int
	strikeCount int
	outCount    int
	score       int
	hitCount    int
}

func InputPlayersEntry() {

	for i := range players {
		var p string
		fmt.Printf("%d번 선수를 입력해 주세요.\n", i+1)

		fmt.Scanf("%s\n", &p)

		players[i] = p
	}
}

func ShowPlayersList() {
	fmt.Println(" [최종명단]")
	for i := range players {
		fmt.Println(i+1, "번 타자", players[i])
	}
	fmt.Println()
}

func CurrentBatter(i int) {
	fmt.Println(i+1, "번 타자", players[i])

}

func InputSwing() int {
	var swing int
	fmt.Println("1 ~ 9 까지의 숫자를 입력하여 스윙 하세요!")
	fmt.Scanf("%d/n", &swing)
	return swing
}

func Pitching() int {
	rand.Seed(time.Now().Unix())
	var rnd int
	rnd = rand.Intn(9) + 1
	// fmt.Println(rnd)
	return rnd
}

func ShowPitchingResult() {
	fmt.Println("투구 내용:", Pitching())
}

func (c *ProgressOfGameCount) AddBaseCount() {
	c.baseCount++
}

func (c *ProgressOfGameCount) AddFoulCount() {
	c.foulCount++
}

func (c *ProgressOfGameCount) AddStrikeCount() {
	c.strikeCount++
}

func (c *ProgressOfGameCount) AddOutCount() {
	c.outCount++
}

func (c *ProgressOfGameCount) AddScore() {
	c.score++
}

func (c *ProgressOfGameCount) AddHitCount() {
	c.hitCount++
}

func (c *ProgressOfGameCount) ResetBaseCount() {
	c.baseCount = 0
}

func (c *ProgressOfGameCount) ResetFoulCount() {
	c.foulCount = 0
}

func (c *ProgressOfGameCount) ResetStrikeCount() {
	c.strikeCount = 0
}

func (c *ProgressOfGameCount) ResetHitCount() {
	c.hitCount = 0
}

func OutCall() {
	fmt.Println("!!!****!!!! 아    웃 !!!!****!!!")
}

func StrikeCall() {
	fmt.Println("!!*****!!!! 스트라이크 !!!!*****!!")
}

func FoulCall() {
	fmt.Println("!******!!!!! 파   울 !!!!!******!")
}

func HitCall() {
	fmt.Println("------!!!!!! 안   타 !!!!!!------")
}

func ScoredCall() {
	fmt.Println("-----!!!!!!!! 득 점 !!!!!!!!-----")
}

func GameStartCall() {
	fmt.Println("---------!! Play Ball !!---------")
}

func GameOverCall() {
	fmt.Println("---- 게 --- 임 ---- 종 --- 료 ----")
}

func IfBaseCountIsFour() {

	if c.baseCount == 4 {
		ScoredCall()
		c.AddScore()
		c.ResetBaseCount()
	}
}

func IfFoulCountIsTwo() {

	if c.foulCount == 2 {
		c.AddStrikeCount()
		c.ResetFoulCount()
	}

}

func IfStrikeCountIsThree() {

	if c.strikeCount == 3 {
		OutCall()
		c.AddOutCount()
	}
}

func IfHitCountIsOne() {
	c.ResetHitCount()
}

func ShowCurrentGameCount() {

	fmt.Printf("%d 스트라이크 / %d 아웃 \n", c.strikeCount, c.outCount)
}

func ShowRunnerOnBase() {

	if c.baseCount == 3 {
		fmt.Println("주자 만루")
	} else if c.baseCount == 2 {
		fmt.Println("주자 1, 2 루")
	} else if c.baseCount == 1 {
		fmt.Println("주자 1루")
	} else {
		fmt.Println("주자 없음")
	}
}

func ShowCurrentScore() {

	fmt.Printf("현재 스코어 %d\n", c.score)
}

func InCaseHit(swing, pitch int) {

	if pitch == swing {
		HitCall()
		c.AddBaseCount()
		IfBaseCountIsFour()
	}
}

func InCaseStrike() {

	StrikeCall()
	c.AddStrikeCount()
	IfFoulCountIsTwo()
}

func InCaseSwingLine1(pitch int) {

	for i := range line1 {
		if line1[i] == pitch {
			FoulCall()
			c.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
		break
	}
}

func InCaseSwingLine2(pitch int) {

	for i := range line1 {
		if line2[i] == pitch {
			FoulCall()
			c.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
		break
	}
}

func InCaseSwingLine3(pitch int) {

	for i := range line1 {
		if line3[i] == pitch {
			FoulCall()
			c.AddFoulCount()
			IfFoulCountIsTwo()
			break
		}
		break
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

// 	for {

// 			break
// 		} else if swing <= 3 {
// 			for i := 0; i < 3; i++ {
// 				if line1[i] == pitch {
// 					fmt.Println("*******!! 파   울 !!*******")
// 					swing = pitch
// 					c.foulCount++
// 					c.FoulCount()
// 					break
// 				}
// 			}
// 			if swing != pitch {
// 				fmt.Println("*******!! 스트라이크 !!*******")
// 				c.strikeCount++
// 				c.StrikeCount()
// 				break
// 			}
// 			break
// 		} else if swing >= 4 && swing < 7 {
// 			for i := 0; i < 3; i++ {
// 				if line2[i] == pitch {
// 					fmt.Println("*******!! 파   울 !!*******")
// 					c.foulCount++
// 					c.FoulCount()
// 					swing = pitch
// 					break
// 				}
// 			}
// 			if swing != pitch {
// 				fmt.Println("*******!! 스트라이크 !!*******")
// 				c.strikeCount++
// 				c.StrikeCount()
// 				break
// 			}
// 			break
// 		} else if swing >= 7 && swing < 10 {
// 			for i := 0; i < 3; i++ {
// 				if line3[i] == pitch {
// 					fmt.Println("*******!! 파   울 !!*******")
// 					c.foulCount++
// 					c.FoulCount()
// 					swing = pitch
// 					break
// 				}
// 			}
// 			if swing != pitch {
// 				fmt.Println("*******!! 스트라이크 !!*******")
// 				c.strikeCount++
// 				c.StrikeCount()
// 				break
// 			}
// 			break
// 		}
// 		break
// 	}
// }

func main() {

	InputPlayersEntry()

	ShowPlayersList()

	GameStartCall()

	c = &ProgressOfGameCount{0, 0, 0, 0, 0, 0}

	for i := range players {
		for j := 0; j < 3; j++ {
			CurrentBatter(i)

			pitch := Pitching()

			swing := InputSwing()

			ShowPitchingResult()

			InCaseHit(swing, pitch)

			CompareAllCase(swing, pitch)

			ShowCurrentGameCount()

			ShowCurrentScore()

			ShowRunnerOnBase()

			if c.strikeCount == 3 {
				IfStrikeCountIsThree()
				break
			} else if c.hitCount == 1 {
				IfHitCountIsOne()
				break
			}
		}
		if c.outCount == 3 {
			GameOverCall()
			break
		}
	}

}
