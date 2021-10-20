package main

import (
	"math/rand"
	"time"

	"github.com/jinzza456/baseball_project/BaseballLib"
)

// 오타 , 카멜, 함수의 기능은 하나만 ,함수와 매서드 중 적합한걸 사용, 매서드명 및 함수명 이 의미있게
// 상수는 메모리의 로스를 줄이기 위해 const 나 var 로 뺴놓는게 좋음
var (
	players = make([]string, 9, 9)
)

func Pitching() int {
	rand.Seed(time.Now().Unix())
	var rnd int
	rnd = rand.Intn(9) + 1
	return rnd
}

func main() {

	BaseballLib.InputPlayersEntry(players)

	BaseballLib.ShowPlayersList(players)

	BaseballLib.GameStartCall()

	c := &BaseballLib.ProgressOfGameCount{0, 0, 0, 0, 0, 0}

	for i := range players {
		for j := 0; j < 3; j++ {
			BaseballLib.CurrentBatter(i, players)

			pitch := Pitching()

			swing := BaseballLib.InputSwing()

			BaseballLib.ShowPitchingResult()

			BaseballLib.InCaseHit(swing, pitch)

			BaseballLib.CompareAllCase(swing, pitch)

			BaseballLib.ShowCurrentGameCount()

			BaseballLib.ShowCurrentScore()

			BaseballLib.ShowRunnerOnBase()

			if c.strikeCount == 3 {
				BaseballLib.IfStrikeCountIsThree()
				break
			} else if c.hitCount == 1 {
				BaseballLib.IfHitCountIsOne()
				break
			}
		}
		if c.outCount == 3 {
			BaseballLib.GameOverCall()
			break
		}
	}

}
