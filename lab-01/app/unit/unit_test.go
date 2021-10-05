package unit_test

import "testing"

import dist "../distance"

func TestStringEmpty(t *testing.T) {
	eth := 0
	
	w1 := []rune("")
	w2 := []rune("")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}

func TestString1(t *testing.T) {
	eth := 2
	
	w1 := []rune("book")
	w2 := []rune("back")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}

func TestString2(t *testing.T) {
	eth := 8
	
	w1 := []rune("critical")
	w2 := []rune("colleague")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}


func TestString3(t *testing.T) {
	eth := 5
	
	w1 := []rune("reptile")
	w2 := []rune("perfume")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}


func TestString4(t *testing.T) {
	eth := 5
	
	w1 := []rune("reptile")
	w2 := []rune("perfume")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}

func TestString5(t *testing.T) {
	eth := 4
	
	w1 := []rune("fset")
	w2 := []rune("note")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	eth = 3

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}


func TestString6(t *testing.T) {
	eth := 2
	
	w1 := []rune("bow")
	w2 := []rune("elbow")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	eth = 2

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}

func TestString7(t *testing.T) {
	eth := 0
	
	w1 := []rune("same")
	w2 := []rune("same")
	res := dist.LevenshteinMatrixIterative(w1, w2)

	if (res != eth) {
		t.Fatalf("ERROR(LevInter) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinPartialMemo(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecMemo) : want %d, have %d\n", eth, res)
	}

	res = dist.LevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(LevRecursive) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinIterative(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevIter) : want %d, have %d\n", eth, res)
	}

	res = dist.DamerauLevenshteinRecursive(w1, w2)
	if (res != eth) {
		t.Fatalf("ERROR(DLevRec) : want %d, have %d\n", eth, res)
	}

}