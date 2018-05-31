package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aybabtme/color/brush"
)

func direction(s string) string {
	return fmt.Sprintf("%s", brush.Red(s))
}

func place(s string) string {
	return fmt.Sprintf("%s", brush.Blue(s))
}

func time(s string) string {
	return fmt.Sprintf("%s", brush.Green(s))
}

var (
	prepositions = []string{
		"in",
		"to",
		"at",
		"into",
		"after",
		"before",
		"behind",
		"above",
		"across",
		"against",
		"anti",
		"below",
		"beneath",
		"between",
		"beyond",
		"by",
		"down",
		"during",
		"for",
		"from",
		"inside",
		"near",
		"of",
		"off",
		"on",
		"out",
		"outside",
		"over",
		"past",
		"since",
		"through",
		"throughout",
		"till",
		"toward",
		"under",
		"underneath",
		"until",
		"up",
		"via",
		"within",
		// TODO: more word!
	}
	times = []string{
		"morning",
		"evening",
		"tonight",
		"tommorow",
		"yesterday",
		"today",
		"twilight",
		"night",
		"dusk",
		"dark",
		"dawn",
	}
	unitTimes = []string{
		"night",
		"years",
		"o'clock",
		"months",
		"weeks",
		"days",
		"hours",
		"minutes",
		"seconds",
		"year",
		"month",
		"week",
		"day",
		"hour",
		"minute",
		"second",
	}
	verbs = []string{
		"be",
		"become",
		"begin",
		"break",
		"bring",
		"build",
		"buy",
		"catch",
		"choose",
		"come",
		"cut",
		"do",
		"draw",
		"drink",
		"drive",
		"eat",
		"fall",
		"feel",
		"fight",
		"find",
		"fly",
		"forget",
		"get",
		"give",
		"go",
		"grow",
		"have",
		"hear",
		"hit",
		"hold",
		"keep",
		"know",
		"lay",
		"leave",
		"lend",
		"let",
		"lie",
		"lose",
		"make",
		"mean",
		"meet",
		"pay",
		"put",
		"read",
		"ride",
		"ring",
		"rise",
		"run",
		"say",
		"see",
		"sell",
		"send",
		"shine",
		"show",
		"sing",
		"sink",
		"sit",
		"sleep",
		"speak",
		"spend",
		"stand",
		"steal",
		"swim",
		"take",
		"teach",
		"tell",
		"think",
		"throw",
		"use",
		"understand",
		"wake",
		"wear",
		"win",
		"write",

		"live",
		"finish",
		"play",
	}
	direcs = []string{
		"above",
		"across",
		"behind",
		"beneath",
		"beyond",
		"down",
		"into",
		"out",
	}
)

func main() {
	// read sentence
	si := bufio.NewScanner(os.Stdin)
	// explain color
	fmt.Printf("%s %s %s\n",
		time("time"),
		place("place"),
		direction("direction"))

	for {
		if !si.Scan() {
			break
		}
		s := si.Text()
		// processing sentence
		ss := strings.Split(s, " ")
		pPos := make([]int, 0, 0)
		for i, sss := range ss {
			if contains(sss, prepositions) != -1 {
				pPos = append(pPos, i)
			}
		}

		for _, pos := range pPos {
			// timeならtime(ss[idx])
			// placeならplace(ss[idx])
			// directionならdirection(ss[idx])
			// TODO: それぞれの前置詞について考える
			if contains(ss[pos], prepositions) != -1 {
				ss[pos] = solve(ss, pos)
			}
		}

		// output
		for _, sss := range ss {
			fmt.Print(sss, " ")
		}
		fmt.Println()
	}
}

func contains(s string, ss []string) int {
	for i, sss := range ss {
		if s == sss {
			return i
		}
	}

	// return -1 ... nothing
	return -1
}

func isTime(s []string, n int) bool {
	if contains(s[n], []string{
		"after",
		"before",
		"during"}) != -1 {
		return true
	}
	if len(s) <= n {
		return false
	}
	if len(s) <= n+1 {
		return false
	}
	// 前置詞の次がmorning, eveningなど一字で時間を表すものかどうか
	if contains(s[n+1], times) != -1 {
		return true
	}
	if len(s) <= n+2 {
		return false
	}
	// 前置詞の2語後がyearなどかどうか見る
	if contains(s[n+2], unitTimes) != -1 {
		return true
	}
	return false
}

func solve(s []string, n int) string {
	/*   time      ... contains times array
	 *   direction ... to XXXX (without infinitive)
	 *   place     ... otherwise
	 */
	pre := s[n]

	// 前置詞が最後に来た場合
	if len(s) == n-1 {
		return pre
	}

	// 時間関係かどうか判定
	if isTime(s, n) {
		return time(pre)
	}

	// direction 判定(direcsの中のものはすべて方向のものと見る)
	if contains(pre, direcs) != -1 {
		return direction(pre)
	}

	// それ以外なら、to XXXX(XXXX!=動詞) ならばdirection, それ以外ならplaceとする
	switch pre {
	case "to":
		// to 動詞を弾く(動詞は予め何個かピックアップしてる中しか見ない)
		if contains(s[n+1], verbs) == -1 {
			return direction(pre)
		}
		return pre
	}

	return place(pre)
}
