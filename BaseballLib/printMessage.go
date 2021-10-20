package BaseballLib

import "fmt"

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
	fmt.Println("---------!! 플레이 볼 !!---------")
}

func GameOverCall() {
	fmt.Println("---- 게 --- 임 ---- 종 --- 료 ----")
}

func ShowRunnerOnBase(c *ProgressOfGameCount) {
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

func ShowCurrentScore(c *ProgressOfGameCount) {
	fmt.Printf("현재 스코어 %d\n", c.score)
}

func ShowPitchingResult(p int) {
	fmt.Println("투구 내용:", p)
}

func ShowPlayersList(p []int) {
	players := p
	fmt.Println(" [최종명단]")
	for i := range players {
		fmt.Println(i+1, "번 타자", players[i])
	}
	fmt.Println()
}

func CurrentBatter(i int, p []int) {
	players := p
	fmt.Println(i+1, "번 타자", players[i])
}

func ShowCurrentGameCount(c *ProgressOfGameCount) {
	fmt.Printf("%d 스트라이크 / %d 아웃 \n", c.strikeCount, c.outCount)
}
