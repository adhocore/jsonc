package jsonc

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"unicode"
)

// Jsonc is the structure for parsing json with comments
type Jsonc struct {
	comment  int
	commaPos int
	inStr    bool
	index    int
	len      int
}

// New creates Jsonc struct with proper defaults
func New() *Jsonc {
	return &Jsonc{0, -1, false, 0, 0}
}

// Strip strips comments and trailing commas from input byte array
func (j *Jsonc) Strip(jsonb []byte) []byte {
	s := j.StripS(string(jsonb))
	return []byte(s)
}

var crlf = map[string]string{"\n": `\n`, "\t": `\t`, "\r": `\r`}

// StripS strips comments and trailing commas from input string
func (j *Jsonc) StripS(data string) string {
	var oldprev, prev, char, next, s string

	j.doReset()
	j.len = len(data)

	for j.index < j.len {
		oldprev, prev, char, next = j.getSegments(data, prev)
		j.index++

		j.checkTrailComma(char)
		if (char == "]" || char == "}") && j.commaPos > -1 {
			s = j.trimTrailComma(s)
		}
		if j.inStringOrCommentEnd(prev, char, char+next, oldprev) {
			if c, ok := crlf[char]; ok && j.inStr {
				char = c
			}
			s += char
			continue
		}

		wasSingle := j.comment == 1
		if j.hasCommentEnded(char, next) && wasSingle {
			s = strings.TrimRight(s, "\r\n\t ") + char
		}
		if char+next == "*/" {
			j.index++
		}
	}
	return s
}

// Unmarshal strips and parses the json byte array
func (j *Jsonc) Unmarshal(jsonb []byte, v interface{}) error {
	return json.Unmarshal(j.Strip(jsonb), v)
}

// UnmarshalFile strips and parses the json content from file
func (j *Jsonc) UnmarshalFile(file string, v interface{}) error {
	jsonb, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return j.Unmarshal(jsonb, v)
}

func (j *Jsonc) doReset() {
	j.inStr = false
	j.index = 0
	j.commaPos = -1
}

func (j *Jsonc) getSegments(json, old string) (oldprev, prev, char, next string) {
	oldprev = old
	if j.index > 0 {
		prev = json[j.index-1 : j.index]
	}
	char = json[j.index : j.index+1]
	if j.index < j.len-1 {
		next = json[j.index+1 : j.index+2]
	}
	return
}

func (j *Jsonc) checkTrailComma(char string) {
	if char == "," || j.commaPos == -1 {
		if char == "," {
			j.commaPos++
		}
		return
	}

	rchar := []rune(char[0:1])
	if unicode.IsDigit(rchar[0]) || strings.ContainsAny(char, `"tfn{[`) {
		j.commaPos = -1
		return
	}

	if char != "]" && char != "}" {
		j.commaPos++
	}
}

func (j *Jsonc) trimTrailComma(s string) string {
	pos := len(s) - j.commaPos - 1
	s = strings.TrimRight(s[0:pos], ",") + strings.TrimLeft(s[pos:], ",")
	j.commaPos = -1
	return s
}

func (j *Jsonc) inStringOrCommentEnd(prev, char, charnext, oldprev string) bool {
	return j.inString(prev, char, charnext, oldprev) || j.inCommentEnd(charnext)
}

func (j *Jsonc) inString(prev, char, charnext, oldprev string) bool {
	if j.comment == 0 && char == `"` && prev != "\\" {
		j.inStr = !j.inStr
		return j.inStr
	}
	if j.inStr && (charnext == `":` || charnext == `",` || charnext == `"]` || charnext == `"}`) {
		j.inStr = oldprev+prev != "\\\\"
	}
	return j.inStr
}

func (j *Jsonc) inCommentEnd(charnext string) bool {
	if !j.inStr && j.comment == 0 {
		if charnext == "//" {
			j.comment = 1
		}
		if charnext == "/*" {
			j.comment = 2
		}
	}
	return j.comment == 0
}

func (j *Jsonc) hasCommentEnded(char, next string) bool {
	singleEnded := j.comment == 1 && char == "\n"
	if singleEnded || (j.comment == 2 && char+next == "*/") {
		j.comment = 0
		return true
	}
	return false
}
