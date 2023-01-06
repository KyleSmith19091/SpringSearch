package query

import (
    "github.com/KyleSmith19091/SpringSearch/document"
)

type Query interface {
    GetTerm() string
    Execute([]string, []*document.Document) []*document.Document
}

