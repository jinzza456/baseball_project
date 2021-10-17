package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	players := make([]string, 9)

	for i := 0; i < len(players); i++ {
		var p string
		fmt.Printf("%d번 선수를 입력해 주세요.\n", i+1)

		fmt.Scanf("%s\n", &p)

		players[i] = p
	}

	fmt.Println(" [최종명단]")
	for i := 0; i < len(players); i++ {
		fmt.Println(i+1, "번 타자", players[i])
	}

	fmt.Println("--------!!playeball!!--------")

	// 게임 시작
	var score *Count
	// score.basecount = 0
	// score.strikecount = 0
	// score.outcount = 0
	// score.score = 0
	for {
		for i := 0; i < len(players); i++ {
			fmt.Println(i+1, "번 타자", players[i])
			fmt.Println("1 ~ 9 까지의 숫자를 입력하여 스윙 하세요!")
			var swing int
			fmt.Scanf("%d/n", &swing)
			score.Compare(swing)

		}

	}

}

func Pitching() int {
	rand.Seed(time.Now().Unix())
	var rnd int
	rnd = rand.Intn(9) + 1
	// fmt.Println(rnd)
	return rnd
}

type Count struct {
	basecount   int
	foulcount   int
	strikecount int
	outcount    int
	score       int
}

func (c *Count) Score() {
	if c.basecount == 4 {
		fmt.Println("------!!!! 득 점 !!!!------")
		c.basecount = 0
		fmt.Println("현재 스코어", c.score, "점")
	}
}

// Piching과 입력값 비교하기
func (c *Count) Compare(swing int) {
	line1 := [3]int{1, 2, 3}
	line2 := [3]int{4, 5, 6}
	line3 := [3]int{7, 8, 9}
	pitch := Pitching()

	for {
		if pitch == swing {
			fmt.Println("------!!!! 안   타 !!!!------")
			c.basecount++
			break
		} else if swing <= 3 {
			for i := 0; i < 3; i++ {
				if line1[i] == pitch {
					fmt.Println("*******!! 파   울 !!*******")
					swing = pitch
					c.strikecount++
					break
				}
			}
			if swing != pitch {
				fmt.Println("*******!! 스트라이크 !!*******")
				c.strikecount++
				break
			}
		} else if swing >= 4 && swing < 7 {
			for i := 0; i < 3; i++ {
				if line2[i] == pitch {
					fmt.Println("*******!! 파   울 !!*******")
					c.strikecount++
					swing = pitch
					break
				}
			}
			if swing != pitch {
				fmt.Println("*******!! 스트라이크 !!*******")
				c.strikecount++
				break
			}

		} else if swing >= 7 && swing < 10 {
			for i := 0; i < 3; i++ {
				if line3[i] == pitch {
					fmt.Println("*******!! 파   울 !!*******")
					c.strikecount++
					swing = pitch
					break
				}
			}
			if swing != pitch {
				fmt.Println("*******!! 스트라이크 !!*******")
				c.strikecount++
				break
			}
		}
		break
	}
}
