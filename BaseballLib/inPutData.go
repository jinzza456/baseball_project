package BaseballLib

import "fmt"

func InputPlayersEntry(s []string) {
	players := s
	for i := range players {
		var p string
		fmt.Printf("%d번 선수를 입력해 주세요.\n", i+1)

		fmt.Scanf("%s\n", &p)

		players[i] = p
	}
}

func InputSwing() int {
	var swing int
	fmt.Println("1 ~ 9 까지의 숫자를 입력하여 스윙 하세요!")
	fmt.Scanf("%d/n", &swing)
	return swing
}
