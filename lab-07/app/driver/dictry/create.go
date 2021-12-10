package dictry

import (
	"encoding/csv"
	"os"
	"strconv"
)

func FromCSV(filename string) (Dictionary, error)  {
	var dct Dictionary
	dct.Size = DSIZE

	dct.entries = make([]Entry, dct.Size)
	path := "../meta/" + filename;

	stream, err := os.Open(path)
	if err != nil {
		return dct, err
	}

	lines, err := csv.NewReader(stream).ReadAll()
    if err != nil {
        return dct, err
    }
	for i, line := range lines {
		dct.entries[i].Key = line[0]
		dct.entries[i].Value, _ = strconv.Atoi(line[1])
	}

	return dct, nil
}

func (dct Dictionary) SegmentByAlphabet() SgtDictionary {
	var (
		sdict   SgtDictionary
		cur_ind int = 0
		d Dictionary
	)

	sdict.dicts = make([]Dictionary, 0)

	dct.sortAlphab()

	for (cur_ind != len(dct.entries)) {
		d, cur_ind = getEntry(cur_ind, dct);
		sdict.dicts = append(sdict.dicts, d)
	}

	sdict.sortSegments()
	return sdict
}

func getEntry(start int, dct Dictionary) (Dictionary, int) {
	seg := make([]Entry, 0)
	i := start

	letter := dct.entries[start].Key[0]
	for i = start; i < len(dct.entries) && dct.entries[i].Key[0] == letter; i++ {
		seg = append(seg, dct.entries[i])
	}

	d := Dictionary {len(seg), seg}
	return d, i
}

func (dct Dictionary) sortAlphab() {
	for i := 0; i < len(dct.entries); i++ {
		for j := 1; j < len(dct.entries) - i; j++ {
			if dct.entries[j].Key < dct.entries[j - 1].Key {
				dct.entries[j], dct.entries[j - 1] = dct.entries[j - 1], dct.entries[j]
			}
		}
	}
}

func (sdct SgtDictionary) sortSegments() {
	for i := 0; i < len(sdct.dicts); i++ {
		for j := 1; j < len(sdct.dicts) - i; j++ {
			if sdct.dicts[j].Size > sdct.dicts[j - 1].Size {
				sdct.dicts[j], sdct.dicts[j - 1] = sdct.dicts[j - 1], sdct.dicts[j]
			}
		}
	}
}
