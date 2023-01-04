package termvector

import (
	"bytes"
	"encoding/gob"
)

type TermVector struct {
    VectorStore map[string][]uint64
}

func NewTermVector() *TermVector {
    return &TermVector {
        VectorStore: make(map[string][]uint64),
    }
}

func (v *TermVector) AddTerm(term string, idx uint64) {
    v.VectorStore[term] = append(v.VectorStore[term], idx)
}

func (v *TermVector) AddTerms(term string, idxs []uint64) {
    v.VectorStore[term] = append(v.VectorStore[term], idxs...)
}

func (v *TermVector) Serialize() ([]byte, error) {
    buf := new(bytes.Buffer)
    encoder := gob.NewEncoder(buf)
    err := encoder.Encode(v.VectorStore)
    
    if err != nil {
        return []byte{}, err
    }
    
    return buf.Bytes(), nil
}

