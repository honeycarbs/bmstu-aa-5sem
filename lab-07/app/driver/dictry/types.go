package dictry

// const DSIZE = 4798
const DSIZE = 2883

type Entry struct {
	Key   string
	Value int
}

type Dictionary struct {
	Size    int
	entries []Entry
}

type SgtDictionary struct {
	Size    int
	dicts   []Dictionary
}