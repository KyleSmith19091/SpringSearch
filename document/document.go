package document

import (
	"bytes"
	"encoding/gob"

	termvector "github.com/KyleSmith19091/SpringSearch/termVector"
	"github.com/KyleSmith19091/SpringSearch/tokeniser"
)

type Document struct {
    Title string
    Path uint64
    TermVector *termvector.TermVector
}

func NewDocument(tokeniser *tokeniser.Tokeniser, title string, body string, path uint64) *Document {
    tokens := tokeniser.Tokenise(body)
    tv := termvector.NewTermVector()

    for _, token := range tokens {
        tv.AddTerm(token.Content, token.Pos)
    }

    return &Document {
        Title: title,
        Path: path,
        TermVector: tv,
    }
}

func NewDocumentFromTokens(title string, tokens []*tokeniser.Token, path uint64) *Document {
    tv := termvector.NewTermVector()

    for _, token := range tokens {
        tv.AddTerm(token.Content, token.Pos)
    }

    return &Document {
        Title: title,
        Path: path,
        TermVector: tv,
    }
}

// Might be used in future development v
func (doc *Document) Serialize() ([]byte, error) {
    buf := new(bytes.Buffer)
    encoder := gob.NewEncoder(buf)
    err := encoder.Encode(doc)

    if err != nil {
        return []byte{}, err
    }

    return buf.Bytes(), err
}

func Deserialize(data []byte) (*Document, error) {
    doc := &Document{}   
    dec := gob.NewDecoder(bytes.NewBuffer(data))
    err := dec.Decode(doc)

    if err != nil {
        return nil, err
    }

    return doc, nil
}
