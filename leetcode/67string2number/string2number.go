package main

type Stat int
type CharType int

const (
	EMPTY Stat = iota
	SIGN
	NUMBER
	STRING
)
const (
	IS_EMPTY CharType = iota
	IS_SIGN
	IS_NUMBER
	IS_STRING
)

func strToInt(str string) int {
	ruleMap := map[Stat]map[CharType]Stat{
		EMPTY: map[CharType]Stat{
			IS_EMPTY:  EMPTY,
			IS_SIGN:   SIGN,
			IS_NUMBER: NUMBER},
		SIGN: map[CharType]Stat{
			IS_NUMBER: NUMBER,
		},
		NUMBER: map[CharType]Stat{
			IS_NUMBER: NUMBER,
		},
	}
	charStat := EMPTY
	var result int
	for _, charFiled := range str {
		if chatType, ok := ruleMap[charStat]; !ok {
			return result
		} else {
			charStat = getType(charFiled)
		}
	}
}

func getType(c rune) Stat {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER
	case '-':
		return SIGN
	case ' ':
		return EMPTY
	default:
		return STRING
	}
}
func main() {

}
