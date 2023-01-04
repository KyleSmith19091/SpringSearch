package tokeniser

import (
	"testing"
)

func check(t *testing.T, tokens []*Token, expectedTokens []string) {

    for idx, token := range tokens {
        if token.Content != expectedTokens[idx] {
            t.Errorf("%v: unexpected tokens", t.Name())
        }
    }
}

func TestTokeniseSimple(t *testing.T) {
    tokeniser := NewTokeniser(TokeniserOptions{
        Seperators: []rune{' ', '-', ','},
    })

    tokens := tokeniser.Tokenise("Hello, world")

    if len(tokens) != 2 {
        t.Errorf("%v: wrong token length", t.Name())
        return
    }

    expected := []string{
        "Hello",
        "world",
    }

    check(t, tokens, expected)
}

func TestTokensieNumbers(t *testing.T) {
    tokeniser := NewTokeniser(TokeniserOptions{
        Seperators: []rune{' ', '-', ','},
    })

    body := "   Hello, world 1234"

    tokens := tokeniser.Tokenise(body)

    if len(tokens) != 3 {
        t.Errorf("%v: wrong token length", t.Name())
        return
    }

    expected := []string{
        "Hello",
        "world",
        "1234",
    }

    check(t, tokens, expected)

}
