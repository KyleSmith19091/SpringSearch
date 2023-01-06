package query

import "springsearch/document"

type TextQuery struct {
    Term string
}

func (t *TextQuery) GetTerm() string {
    return t.Term
}

func (t *TextQuery) Execute(_ []string, docs []*document.Document) ([]*document.Document) {
    return docs
}
