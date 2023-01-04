package result

type SearchResult struct {
    Path uint64
    Title string
}

func NewSearchResult(title string, path uint64) (*SearchResult) {
    return &SearchResult{
        Path: path,
        Title: title,
    }
}
