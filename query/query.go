package query

import "springsearch/document"

type Query interface {
    GetTerm() string
    Execute([]string, []*document.Document) []*document.Document
}

