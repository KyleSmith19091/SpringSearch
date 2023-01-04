package invertedindex

import (
	"bytes"
	"encoding/gob"

	"springsearch/document"
	"springsearch/tokeniser"
)

type InvertedIndex struct {
    IndexStore map[string][]*document.Document
}

func NewInvertedIndex() *InvertedIndex {
    return &InvertedIndex {
        IndexStore: make(map[string][]*document.Document),
    }
}

func (i *InvertedIndex) AddDocument(term string, doc *document.Document) {
    i.IndexStore[term] = append(i.IndexStore[term], doc)
}

func (i *InvertedIndex) AddDocumentFromTokens(terms []*tokeniser.Token, doc *document.Document) {
    for _, term := range terms {
        i.IndexStore[term.Content] = append(i.IndexStore[term.Content], doc)
    }
}

func (i *InvertedIndex) GetDocument(term string) ([]*document.Document, bool) {
    doc, found := i.IndexStore[term]
    return doc, found
}

func keys(docMap map[*document.Document]uint64) []*document.Document {
    keys := make([]*document.Document, 0, len(docMap))
    for k := range docMap {
        keys = append(keys, k)
    }

    return keys
}

func (i *InvertedIndex) GetDocuments(terms []string) ([]*document.Document) {
    docOccurences := make(map[*document.Document]uint64) 

    for _, term := range terms {
        docArr, found := i.IndexStore[term] 
        if found {
            for _, doc := range docArr {
                docOccurences[doc] += 1                
            } 
        } 
    }

    return keys(docOccurences)
}

func (i *InvertedIndex) Serialize() ([]byte, error) {
    buf := new(bytes.Buffer)
    encoder := gob.NewEncoder(buf)
    err := encoder.Encode(i)

    if err != nil {
        return []byte{}, err
    }

    return buf.Bytes(), err
}

func Deserialize(data []byte) (*InvertedIndex, error) {
    i := &InvertedIndex{}
    decoder := gob.NewDecoder(bytes.NewBuffer(data))
    err := decoder.Decode(i)

    if err != nil {
        return nil, err
    }

    return i, nil
}
