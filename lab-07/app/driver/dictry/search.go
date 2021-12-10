package dictry

import (
	"errors"
)

func (dct Dictionary) BruteForce(skey string) (int, int, error) {
	compars := 0

	for _, entry := range(dct.entries) {
		compars++;
		if (entry.Key == skey) {
			return entry.Value, compars, nil;
		}
	}
	return 0, compars, errors.New("NOT FOUND")
}

func (dct Dictionary) BinarySearch(skey string) (int, int, error) {
	dct.sortAlphab()
	
	r, compars, err := dct.__binaryHelper(skey)
	return r, compars, err
}

func (dct Dictionary) __binaryHelper(skey string) (result int, compars int, err error) {
	var (
		high    int = len(dct.entries)
		mid     int = high / 2
	)

	switch {
	case high == 0:
		result = 0
		err    = errors.New("NOT FOUND")
	case dct.entries[mid].Key > skey:
		dct.entries = dct.entries[:mid]
		result, compars, err = dct.__binaryHelper(skey)
	case dct.entries[mid].Key < skey:
		dct.entries = dct.entries[mid + 1:]
		result, compars, err = dct.__binaryHelper(skey)
	default:
		result = dct.entries[mid].Value
	}
	compars++
	return
}

func (dct Dictionary) SectionedBinary(skey string) (int, int, error) {
	var (
		sec_compars int = 0
		sletter byte = skey[0]
		rvalue  int
		err     error
		compars int
	)

	sdct := dct.SegmentByAlphabet()
	rvalue = 0

	for _, d := range sdct.dicts {
		sec_compars++
		if sletter == d.entries[0].Key[0] {
			rvalue, compars, err = d.BinarySearch(skey)
			break
		}
	}

	return rvalue, compars, err
}