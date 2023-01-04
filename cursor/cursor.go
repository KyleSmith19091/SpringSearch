package cursor

import (
	"errors"
	"springsearch/result"
)

type Cursor struct {
    db func(*result.SearchResult) (string,error)
    Current string
    results []result.SearchResult
    currIdx uint64
}

func NewCursor(readFunc func(res *result.SearchResult) (string,error), data []result.SearchResult) (*Cursor) {
    return &Cursor{
        db: readFunc,
        Current: "",                
        results: data,
        currIdx: 0,
    }
}

func (c *Cursor) Next() (string, error) {
    if c.currIdx >= uint64(len(c.results)) {
        return "", errors.New("index out of bounds")
    }

    data, err := c.db(&c.results[c.currIdx]) 
    c.currIdx += 1
    return data, err
}

func (c *Cursor) HasNext() (bool) {
    return c.currIdx < uint64(len(c.results))
}

func (c *Cursor) NumResults() (uint64) {
    return uint64(len(c.results))
}
