package tokeniser

import (
	"log"
	"strings"
)

type Token struct {
    Pos uint64
    Content string
}

type TokeniserOptions struct {
    Seperators []rune
}

type Tokeniser struct {
    options TokeniserOptions
}

func NewTokeniser(options TokeniserOptions) *Tokeniser {
    return &Tokeniser {
        options: options,
    }
}

func (t *Tokeniser) splitFunc(c rune) bool {
    for _, s := range t.options.Seperators {
        if(c == s) { // Is c a seperator
            return true
        }
    }
    return false
}

func (t *Tokeniser) IsSeperator(c rune) bool {
    for _, s := range t.options.Seperators {
        if(c == s) { // Is c a seperator
            return true
        }
    }

    return false
}

func (t *Tokeniser) Tokenise(body string) []*Token {
    var tokens []*Token
	buffer := ""
    var i uint64
    i = 0

	// Iterate through the string
	for _, c := range body {
		// Check if the character is a separator
		if t.IsSeperator(c) {
			if buffer != "" {
                tokens = append(tokens, &Token{Pos: i - uint64(len(buffer)), Content: buffer})
				buffer = ""
			}
		} else {
			buffer += string(c)
		}
        i += 1
	}

    log.Println(i)

    if buffer != "" {
        tokens = append(tokens, &Token{Pos: uint64(i), Content: buffer})
	}

    return tokens
}

func (t *Tokeniser) Split(body string) []string {
    terms := strings.FieldsFunc(body, t.splitFunc)
    return terms
} 
