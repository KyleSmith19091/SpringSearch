package db

import (
	"log"

	"springsearch/cursor"
	"springsearch/document"
	invertedindex "springsearch/invertedIndex"
	"springsearch/result"
	"springsearch/tokeniser"
    "springsearch/query"

	walLog "github.com/tidwall/wal"
)

type DB struct {
    invertedIndex *invertedindex.InvertedIndex
    tokeniser *tokeniser.Tokeniser
    wal *walLog.Log
    data *walLog.Log
    currIdx uint64
}

func NewDB(walPath string, dataPath string) (*DB, error) {
    wal, err := walLog.Open(walPath, nil)

    if err != nil {
        return nil, err
    }

    currIdx, err := wal.LastIndex()

    if err != nil {
        return nil, err
    }

    data, err := walLog.Open(dataPath, nil)

    if err != nil {
        return nil, err
    }

    db := &DB{
        currIdx: currIdx,
        wal: wal,
        data: data,
        invertedIndex: invertedindex.NewInvertedIndex(),
        tokeniser: tokeniser.NewTokeniser(tokeniser.TokeniserOptions{
            Seperators: []rune{'-','!',',','.','/','*','^','&','%','$',' ', '\'', '"'},
        }),
    }

    applyLog(db)

    return db, nil
}

func (db *DB) Search(q query.Query) *cursor.Cursor {
    tokens := db.tokeniser.Split(q.GetTerm())
    docs := db.invertedIndex.GetDocuments(tokens)

    var docData []result.SearchResult

    for _, doc := range docs {
        docData = append(docData, result.SearchResult{Title: doc.Title, Path: doc.Path})
    }

    return cursor.NewCursor(db.GetResultData, docData) 
}

func (db *DB) GetResultData(res *result.SearchResult) (string,error) {
    data, err := db.data.Read(res.Path)

    if err != nil {
        return "", err
    }

    return string(data), nil
}
 
func applyLog(db *DB) {
    // Get index of last log entry
    lastIdx, err := db.wal.LastIndex()

    // Error reading last log entry index
    if err != nil {
        log.Fatal(err)
        return
    }

    // If the log is empty just continue
    if lastIdx == 0 {
        return
    }

    // Read most up to date index
    data, err := db.wal.Read(lastIdx)

    // Error reading log entry
    if err != nil {
        log.Fatal(err)
        return
    }

    // Convert byte data to index instance
    i, err := invertedindex.Deserialize(data)

    if err != nil {
        log.Fatal(err)
        return
    }

    db.invertedIndex = i

    log.Println("Done Rebuilding Index")
}

func (db *DB) writeInsertToLog() {
    docData, err := db.invertedIndex.Serialize()

    if err != nil {
        return
    }

    db.wal.Write(db.currIdx+1, docData)
    db.currIdx += 1
}

func (db *DB) writeData(body string) (uint64, error) {
    lastIdx, err := db.data.LastIndex()
    
    if err != nil {
        return 0, err
    }

    db.data.Write(lastIdx+1, []byte(body))

    return lastIdx+1, nil 
}

func (db *DB) Insert(title string, body string) {
    idx, err := db.writeData(body)

    if err != nil {
        log.Printf("Error inserting document %s\n", title)
        return
    }

    tokens := db.tokeniser.Tokenise(body)

    doc := document.NewDocumentFromTokens(title, tokens, idx) 

    db.invertedIndex.AddDocumentFromTokens(tokens, doc)

    db.writeInsertToLog()
}
