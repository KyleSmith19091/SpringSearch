package cursor

import (
	"log"
	"testing"

	"github.com/KyleSmith19091/SpringSearch/result"
)

func GetResultData(res *result.SearchResult) (string,error) {
    return res.Title, nil
}

func createResultArray(data []string) ([]result.SearchResult) {
    var results []result.SearchResult
    for idx, val := range data {
        results = append(results, result.SearchResult{Path: uint64(idx), Title: val})
    }

    return results
}

func TestEmptyCursor(t *testing.T) {
    emptyResult := createResultArray([]string{})
    c := NewCursor(GetResultData, emptyResult)

    if c.NumResults() != 0 {
        log.Fatalf("Incorrect number of results: %v\n", c.NumResults())
        return
    }

    if c.HasNext() {
        log.Fatal("Cursor should not have next")
        return
    }

    _, err := c.Next()

    if err == nil {
        log.Fatal("Next did not return index out of bounds error")
        return
    }

}

func TestCursorIteration(t *testing.T) {
    resStrings := []string {
        "Hello",
        "Kyle",
        "Smith",
        "Is",
        "Cool",
    }
    results := createResultArray(resStrings)
    c := NewCursor(GetResultData, results)

    if c.NumResults() != uint64(len(resStrings)) {
        log.Fatalf("Number of results don't match input %v != %v", c.NumResults(), len(resStrings))
        return
    }

    if !c.HasNext() {
        log.Fatal("Cursor should have next")
        return
    }

    counter := 0
    for c.HasNext() {
        val, err := c.Next()
    
        if err != nil {
            log.Fatalf("Unexpcted error reading from Next(): %v\n", err)
            return
        }

        if val != resStrings[counter] {
            log.Fatalf("Value from Next() and input data don't match %v != %v \n", val, resStrings[counter])
        }

        counter += 1
    }
}
