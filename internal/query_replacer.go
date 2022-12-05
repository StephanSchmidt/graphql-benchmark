package internal

import ()

func FindClosingBrace(s []rune, pos int) int {
	level := 0
	for pos < len(s) {
		if s[pos] == ')' {
			if level == 0 {
				return pos
			} else {
				level--
			}
		} else if s[pos] == '(' {
			level++
		}
		pos = pos + 1
	}
	return -1
}

func Replacer(q string) string {
	newString := ""
	qR := []rune(q)
	i := 0
	for i < len(qR) {
		if qR[i] == rune('$') && qR[i+1] == rune('$') && qR[i+2] == rune('(') {
			pos := FindClosingBrace(qR, i+3)
			if pos == -1 {
				return q
			}
			newString = newString + "select json_build_object(" + string(qR[i+3:pos]) + ")"
			//			newString = newString + "( select json_agg(row_to_json(t)) from " + string(qR[i+3:pos]) + ") t )"
			i = pos
		} else if qR[i] == rune('$') && qR[i+1] == rune('(') {
			pos := FindClosingBrace(qR, i+3)
			if pos == -1 {
				return q
			}
			newString = newString + "( select json_agg(row_to_json(t)) from (" + string(qR[i+3:pos]) + ") t )"
			i = pos
		} else {
			newString = newString + string(qR[i])
		}
		i++
	}

	return newString
}
