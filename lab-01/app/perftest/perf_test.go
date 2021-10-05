package perftest

import "testing"

import dist "../distance"


// func BenchmarkLevenshteinMIterative5(b *testing.B)  {
// 	src := []rune("meaow")
// 	dest := []rune("mutes")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative10(b *testing.B)  {
// 	src := []rune("frizziness")
// 	dest := []rune("frizzliest")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative15(b *testing.B)  {
// 	src := []rune("hdwxsjifjwudloh")
// 	dest := []rune("dmunfuidpdnbxoa")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative20(b *testing.B)  {
// 	src := []rune("zyapjbnxuykcwtnmexid")
// 	dest := []rune("ickfleanmgcpakskibcx")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative30(b *testing.B)  {
// 	src := []rune("xeajpurnpjfmdfshnnsykkfjymfype")
// 	dest := []rune("tooczrfaksfbkmdeiwvutrvkdgyeee")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative50(b *testing.B)  {
// 	src := []rune("xtmtsjiwplzkjfzvrwjznzxxjoiwbfksfujjgrdtgswqzwbfyp")
// 	dest := []rune("ujwfhoxzyupjjuxdreijkhhkytcnppjdnbwaphtxdznlmveleu")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMIterative100(b *testing.B)  {
// 	src := []rune("trxkgfazhjlszzfdifjvkludrbdmpqfaeprnxepdpktkbhixpwpupwynajvntixqvxrlookefisvfdkdgrkkmfwxzcrgpnfggkmy")
// 	dest := []rune("iwxxdqzusrswcsluqfndrfqojwuyaovnprgponhtrykjxtsplzlmfeplteodlfdryhqlxdlfuybkaqktmrezgxufncizfbyxeyyl")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }

func BenchmarkLevenshteinMIterative150(b *testing.B)  {
	src := []rune("dvmhztfkjrpmxovlkoeectyvkepmpxogvvwwafkzjyxwbwxrxnhyvsaommudsolmhgixvpgvlnjiqasfxmdqostuvjvmjtjlwudtzhbnttysaqpaptgxhezminifsbfodmtefvzrjcxildyyatoygw")
	dest := []rune("gvpxawxtuvmxrfbhlsxremghefujuukfvwhdguywpzrijnnexdfirgsnhcsbnzrhpjwsuziepmuopumfbvzcesqqcimfxmhodzvaoecssicvlypyznbguvqbvtdpetwmlmrzcwoiiylczftwkotxct")

	for i := 0; i < b.N; i++ {
		dist.LevenshteinMatrixIterative(src, dest)
	}
	
}

// func BenchmarkLevenshteinMIterative200(b *testing.B)  {
// 	src := []rune("qzjwsrmxohygfheugvwxpnqsfqmmajobzpxhmvtppeqzosoctlxmjewtmjkldjcbywlwdpstxhowdutzewspanxawoybobomfoxyseoefibncxmnkiitttunhtwkzzqzsqaespalwffevnmdwigvinkohlddrebmpshdortwxbxqzsdkyiwjthvtlavodtckzzkiamms")
// 	dest := []rune("uhtxecdpbsjedqlcxpoigvxudutdqduysktmdwoelbrmtqvpwpnryjqbgqoayxmktytexuqbtbawsoavsfggkvmaelgiyoetexuvcqlzaqrzkpvhmkqcivlnelfqmjqmmtnuynjyuneibtvrkscgbmxtpffssjbjvovrwfyqrbmxpulntdnxcqsxkdwuynqxorzwxroz")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinMatrixIterative(src, dest)
// 	}
	
// }


// func BenchmarkLevenshteinMemoRec5(b *testing.B)  {
// 	src := []rune("meaow")
// 	dest := []rune("mutes")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec10(b *testing.B)  {
// 	src := []rune("frizziness")
// 	dest := []rune("frizzliest")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec15(b *testing.B)  {
// 	src := []rune("hdwxsjifjwudloh")
// 	dest := []rune("dmunfuidpdnbxoa")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec20(b *testing.B)  {
// 	src := []rune("zyapjbnxuykcwtnmexid")
// 	dest := []rune("ickfleanmgcpakskibcx")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec30(b *testing.B)  {
// 	src := []rune("xeajpurnpjfmdfshnnsykkfjymfype")
// 	dest := []rune("tooczrfaksfbkmdeiwvutrvkdgyeee")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec50(b *testing.B)  {
// 	src := []rune("xtmtsjiwplzkjfzvrwjznzxxjoiwbfksfujjgrdtgswqzwbfyp")
// 	dest := []rune("ujwfhoxzyupjjuxdreijkhhkytcnppjdnbwaphtxdznlmveleu")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinMemoRec100(b *testing.B)  {
// 	src := []rune("trxkgfazhjlszzfdifjvkludrbdmpqfaeprnxepdpktkbhixpwpupwynajvntixqvxrlookefisvfdkdgrkkmfwxzcrgpnfggkmy")
// 	dest := []rune("iwxxdqzusrswcsluqfndrfqojwuyaovnprgponhtrykjxtsplzlmfeplteodlfdryhqlxdlfuybkaqktmrezgxufncizfbyxeyyl")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

func BenchmarkLevenshteinMemoRec150(b *testing.B)  {
	src := []rune("dvmhztfkjrpmxovlkoeectyvkepmpxogvvwwafkzjyxwbwxrxnhyvsaommudsolmhgixvpgvlnjiqasfxmdqostuvjvmjtjlwudtzhbnttysaqpaptgxhezminifsbfodmtefvzrjcxildyyatoygw")
	dest := []rune("gvpxawxtuvmxrfbhlsxremghefujuukfvwhdguywpzrijnnexdfirgsnhcsbnzrhpjwsuziepmuopumfbvzcesqqcimfxmhodzvaoecssicvlypyznbguvqbvtdpetwmlmrzcwoiiylczftwkotxct")

	for i := 0; i < b.N; i++ {
		dist.LevenshteinPartialMemo(src, dest)
	}
	
}

// func BenchmarkLevenshteinMemoRec200(b *testing.B)  {
// 	src := []rune("qzjwsrmxohygfheugvwxpnqsfqmmajobzpxhmvtppeqzosoctlxmjewtmjkldjcbywlwdpstxhowdutzewspanxawoybobomfoxyseoefibncxmnkiitttunhtwkzzqzsqaespalwffevnmdwigvinkohlddrebmpshdortwxbxqzsdkyiwjthvtlavodtckzzkiamms")
// 	dest := []rune("uhtxecdpbsjedqlcxpoigvxudutdqduysktmdwoelbrmtqvpwpnryjqbgqoayxmktytexuqbtbawsoavsfggkvmaelgiyoetexuvcqlzaqrzkpvhmkqcivlnelfqmjqmmtnuynjyuneibtvrkscgbmxtpffssjbjvovrwfyqrbmxpulntdnxcqsxkdwuynqxorzwxroz")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinPartialMemo(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative5(b *testing.B)  {
// 	src := []rune("meaow")
// 	dest := []rune("mutes")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative10(b *testing.B)  {
// 	src := []rune("frizziness")
// 	dest := []rune("frizzliest")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative15(b *testing.B)  {
// 	src := []rune("hdwxsjifjwudloh")
// 	dest := []rune("dmunfuidpdnbxoa")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative20(b *testing.B)  {
// 	src := []rune("zyapjbnxuykcwtnmexid")
// 	dest := []rune("ickfleanmgcpakskibcx")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}



	
// }

// func BenchmarkDamerauIterative30(b *testing.B)  {
// 	src := []rune("xeajpurnpjfmdfshnnsykkfjymfype")
// 	dest := []rune("tooczrfaksfbkmdeiwvutrvkdgyeee")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative50(b *testing.B)  {
// 	src := []rune("xtmtsjiwplzkjfzvrwjznzxxjoiwbfksfujjgrdtgswqzwbfyp")
// 	dest := []rune("ujwfhoxzyupjjuxdreijkhhkytcnppjdnbwaphtxdznlmveleu")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkDamerauIterative100(b *testing.B)  {
// 	src := []rune("trxkgfazhjlszzfdifjvkludrbdmpqfaeprnxepdpktkbhixpwpupwynajvntixqvxrlookefisvfdkdgrkkmfwxzcrgpnfggkmy")
// 	dest := []rune("iwxxdqzusrswcsluqfndrfqojwuyaovnprgponhtrykjxtsplzlmfeplteodlfdryhqlxdlfuybkaqktmrezgxufncizfbyxeyyl")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

func BenchmarkDamerauIterative150(b *testing.B)  {
	src := []rune("dvmhztfkjrpmxovlkoeectyvkepmpxogvvwwafkzjyxwbwxrxnhyvsaommudsolmhgixvpgvlnjiqasfxmdqostuvjvmjtjlwudtzhbnttysaqpaptgxhezminifsbfodmtefvzrjcxildyyatoygw")
	dest := []rune("gvpxawxtuvmxrfbhlsxremghefujuukfvwhdguywpzrijnnexdfirgsnhcsbnzrhpjwsuziepmuopumfbvzcesqqcimfxmhodzvaoecssicvlypyznbguvqbvtdpetwmlmrzcwoiiylczftwkotxct")

	for i := 0; i < b.N; i++ {
		dist.DamerauLevenshteinIterative(src, dest)
	}
	
}


// func BenchmarkDamerauIterative200(b *testing.B)  {
// 	src := []rune("qzjwsrmxohygfheugvwxpnqsfqmmajobzpxhmvtppeqzosoctlxmjewtmjkldjcbywlwdpstxhowdutzewspanxawoybobomfoxyseoefibncxmnkiitttunhtwkzzqzsqaespalwffevnmdwigvinkohlddrebmpshdortwxbxqzsdkyiwjthvtlavodtckzzkiamms")
// 	dest := []rune("uhtxecdpbsjedqlcxpoigvxudutdqduysktmdwoelbrmtqvpwpnryjqbgqoayxmktytexuqbtbawsoavsfggkvmaelgiyoetexuvcqlzaqrzkpvhmkqcivlnelfqmjqmmtnuynjyuneibtvrkscgbmxtpffssjbjvovrwfyqrbmxpulntdnxcqsxkdwuynqxorzwxroz")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinIterative(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinRecursive5(b *testing.B)  {
// 	src := []rune("meaow")
// 	dest := []rune("mutes")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinRecursive(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinRecursive10(b *testing.B)  {
// 	src := []rune("frizziness")
// 	dest := []rune("frizzliest")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinRecursive(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinRecursive15(b *testing.B)  {
// 	src := []rune("hdwxsjifjwudloh")
// 	dest := []rune("dmunfuidpdnbxoa")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinRecursive(src, dest)
// 	}
	
// }

// func BenchmarkLevenshteinRecursive20(b *testing.B)  {
// 	src := []rune("zyapjbnxuykcwtnmexid")
// 	dest := []rune("ickfleanmgcpakskibcx")

// 	for i := 0; i < b.N; i++ {
// 		dist.LevenshteinRecursive(src, dest)
// 	}
	
// }


// func BenchmarkDamerauLevenshteinRecursive5(b *testing.B)  {
// 	src := []rune("meaow")
// 	dest := []rune("mutes")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinRecursive(src, dest)
// 	}
	
// }

// func BenchmarkDamerauLevenshteinRecursive10(b *testing.B)  {
// 	src := []rune("frizziness")
// 	dest := []rune("frizzliest")

// 	for i := 0; i < b.N; i++ {
// 		dist.DamerauLevenshteinRecursive(src, dest)
// 	}
	
// }

func BenchmarkDamerauLevenshteinRecursive15(b *testing.B)  {
	src := []rune("hdwxsjifjwudloh")
	dest := []rune("dmunfuidpdnbxoa")

	for i := 0; i < b.N; i++ {
		dist.DamerauLevenshteinRecursive(src, dest)
	}
	
}