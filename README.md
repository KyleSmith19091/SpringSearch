# SpringSearch
<div align="center">
  <img width="300px" src="https://user-images.githubusercontent.com/29174023/210623428-3bbca6e5-245f-4b8c-8463-050b46521a07.png" />
</div>

An in-memory text search engine that is durable and persistent. Implementation based on ideas and structures from Elasticsearch and Apache Lucene. Uses
inverted indexes and term vectors to support full text search. Data and indexes are persisted on disk using a Write Ahead Log.

Springsearch is useful for searching Logs, metrics or even for developing a search backend.

## Get Started
First Download the package

```
go get github.com/KyleSmith19091/SpringSearch@v0.1.0
```

```go
import (
    "log"

	"github.com/KyleSmith19091/SpringSearch/db"
    "github.com/KyleSmith19091/SpringSearch/query"
)

func main() {
    // 1st param = Path to index data directory, 2nd param = Path to directory where data should be stored
    db, err := db.NewDB("./mylog", "./data")
    
    if err != nil {
        log.Fatal(err)
    }

    // Perform search query, returns a Cursor
    res := db.Search(&query.TextQuery{Term: "world"})
    log.Printf("%v results found", res.NumResults())
    
    // Iterate over cursor and print data
    for res.HasNext() {
        data, err := res.Next() 
    
        if err != nil {
            break
        }

        log.Println(data)
    }
}
```

## How does this work?
The search engine takes in a query interface(to support multiple types of queries). For the sake of this example
we consider the case of the TextQuery, which has a single attribute called Term. For the sake of this example assume
the index has already been given data to index. Consider we query the engine with the text query "Hello, World!". The engine
takes this string and tokenises the string into a token array ["Hello", "World"].

These tokens are then passed to our inverted index which will look up the documents
that contain these tokens. The index will then pass these documents to the original
query which can then perform further operations on these documents(the text query just returns the given documents).
These documents in turn point to data that is currently on disk, we then  build a cursor structure from the documents.
The cursor ensures that we only load the current document's content in memory.
