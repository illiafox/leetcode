package easy

import (
	"fmt"
	"strings"
	"testing"
)

func TestNumOfStringsUnified(t *testing.T) {
	functions := map[string]func([]string, string) int{
		"Std":          numOfStrings,
		"Custom":       numOfStringsCustom,
		"Aho-Corasick": numOfStringsAhoCorasick,
	}

	tests := []struct {
		name     string
		word     string
		patterns []string
	}{
		{
			name:     "example 1",
			word:     "abc",
			patterns: []string{"a", "abc", "bc", "d"},
		},
		{
			name:     "example 2",
			word:     "aaaaabbbbb",
			patterns: []string{"a", "b", "c", "a", "a", "a"},
		},
		{
			name: "suite 1",
			word: "wsnnrijyzrhrygxvzmtuxulhjriiikuwwibjpodtbmiyivserr",
			patterns: []string{
				"e", "ulhjriiikuwwi", "odtbmiyiv", "iiku", "o", "nao", "ilrbdhfvndjaxohexcynpgocoqorifjngaurokumcqqkirqp",
				"zmt", "tuxulhjriiikuwwibjpodtbmiyiv", "lhjriiikuwwibjpodtbmiyi", "qjnwwkfqhztiyqayiabjjzyqhumh", "lhjriiikuwwibjpodtbmiyiv",
				"oyaloibpkqpyubftinknjraptsuqgejucpfigc", "fdqvajkfornietcdvxltbktlkg", "bcxslqwrjaabvywpynzfojetmnnwrurimjgwl",
				"ncvvyhbznxrxoonfrzfmecfdnrvikffpvrc", "gxvzmtuxulhjriiikuwwibjpodtbmiyi", "ozuwzgetoddakvjwiwzocpwpiavcyclxhlguhfqtpbk",
				"ktwqukguxobakxvbitkkdemvlmqypd", "gxvzmtuxulhjri", "xulhjriiiku", "sjemhhtbqastmbtbvsfnuygqfypmmi", "gqfmyqbjcvuxsjdaeffwthlelb",
				"kuwwibjpodtbmiyivse", "ydwiabsimbkdwmsvgp", "ixrzoktohtadgblem", "uicfykignmptxodjd", "kecqujgjynq", "sruqpwjevngokrbyi",
				"rijyzrhrygxvzmtuxulhjriii", "gsljxjrmkzbwaqdgmwysnblp", "arkmbbafkziemddeqripnyjoavmoaxtn", "gxhijktottj", "ibjpodtbmiyivse",
				"j", "qzvmjstehwibuqcqlzdam", "ikuwwibjpodt", "mnrp", "pduagikyudhcvdnqoaadqvvmhluywfsstqnawfmq", "vse", "iyiv", "miyivs",
				"qdfrbkhffkksgtjpnbgvqtrnigbnoacmqkepllouhqpgskpupy", "wxbqfmzvmmffwtwjhjpuitcfmknnuwdtamtutiukmpsxzu", "odtbmiyi",
				"hrygxvzmtuxulhjriiiku", "yzrhrygxvzmtuxu", "se", "gixgfjtunbltzcwgkrheavilcvektnsizvrrabmysx", "hjriiikuwwibj", "ahaeorlnaqkk",
				"gvnyr", "tynqzmrvdjaznuji", "jbwybcttaspwsbieawybfqxsiwxulwkouezkhbauadatb", "uwwibjpodt", "gxv", "zrhrygxvzmtuxulhjriiikuwwib",
				"wivvuywmwchszzctw", "lijcroypqr", "mrx", "m", "ygngqghhza", "riiikuwwibjpodtbmiyivs", "tymtjqykindweexrfztizyagijnxntrcrcy",
				"vse", "kntxszudqpmtlrxzspcfivbrwonejzywv", "ygxvzmtuxulh", "wsnnrijyzrhrygxvzmtuxulhjriiik", "zpymedyxskwqddvxlycxvq",
				"qdkm", "gxvzmtuxu", "snnrijyzrhrygxvzmtuxulhjriiikuwwibjpo", "iyivserr", "yzrhrygxvzmt", "rygxvzmtuxulhjrii", "snnrijyzrhrygxvzmtuxul",
				"gxvzmtuxulhjriiikuwwibjpodtbmiyivserr", "kuwwibjpodt", "ibjpodtbmiy", "mtuxulhjriii", "zrhrygxvzmtu", "qxjvgdyuqt", "rygxv",
				"gnwmvzqhrodzqdvdexgupqbzogcpmtfq", "oleibiotmojkqexnnlx", "dtb", "jjividhfhdixcvyduyunoiiipyvi", "podtb", "jpodtbmiyi",
				"iiku", "od", "kuwwibjpodtb", "b", "irugjrzvsfp", "wsnnrijyzrhrygxvzmtuxulhjrii", "jriiikuwwibjpodtbmi",
			},
		},
		{
			name: "suite 2",
			word: "rcfcfheebvoavnjoksfjoclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsardyooxubjamzwc",
			patterns: []string{
				"frchczhlghnjfmkvytbv", "qaco", "kavmvdvqpigexcyuxfvgrpaodqvauobdvlwewjxwosn", "jamzw",
				"ksfjoclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsardyooxubjam", "uzilsldtqkmbhcspxapaslyrfqgqrzycxjzbtrmirymzdmgy",
				"abjusimniemqwsrtwvy", "xhqqqxhaqlbqlulldflrbxtjemijsouzuvitigguhkcwdpegrdublohrthtiiqofimwphejhqsvlbnplgxdknyw",
				"kucupdndnbgmhmuwgrewmywclaunnufkugipgxlzmvdzynedjlgylcofymcnesyabmbhcwvsrynogqqia",
				"cfheebvoavnjoksfjoclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiyd", "ldhlsardyooxubjamzw", "xsfwqiycvwbzwyuycgybejhrheouqhvxtfjtuw",
				"iungqacofiydxinyctphldhlsa", "aoewruozodgtndgenepvqrrhgggwbywwjrvbncfxqvrhptwfwlrorvaepfkkkqorhyocwawptdlxbpzchjq",
				"nfewllmzxfttlrrajlfvvpxpgdhplolanacnrewvtoahommyuoczeemvpqtkue", "tsjjbfqwruintrdbqbvmnmjjifqjwntfmcqoosnznrljsplpmkwk",
				"yctphldhlsardyooxubja", "rjkahbpcykorxsjgdbiypdbijukwqpg", "wfexzlrtgkkmdvtkzhbymssrvsriyfgjmldpzdnzdhuabhpiauxdhhhwnsytmcxygmwgoqj",
				"vfohxrisfk", "ardyooxubjamz", "rlumemurqviqgzenildoihrtxbrjpsxbqqdwasvviizoaxpqvlcdmiqeolzsamfzxlqnc",
				"onddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsardy", "oxu", "lddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctph",
				"donddfrchczhlghnjfmkvytbvfohxrisfkgiungqaco", "vnjoksfjoclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsardyooxub",
				"kllckvxmjsmeusenynulmjdiseqtvvvtdecvvfioxpcjzmqbzatzyiozzzhxwixqalvkcoa", "cgqvkvdywfibsqsgwcotrcgwzbgemhdpuqsnwlnsmmjfnz",
				"fsftymjoqkwyrhbixleopimsfischagjciblhwkbyddmrkajeqbxsstrs", "donddfrchczhlghnjfmkvytbvfohxrisfkgiungqacof", "ofiyd",
				"orzbqcmugopxvnfqjjnthwyeawswqzzjujcxkefef", "ngqacofiydxi", "utudwmwhkkzayhmcfwzgknlduoqeykjcdnqezyvs", "ufjhefsxjfnwkwcx",
				"yuzxvcjqrpocrglizlrlyzwjtyoixqmevpchehsnyxxnhihcmdrnscunlcfppzdzbazymplwmm", "pevhwclpkpzy",
				"gsbuzybtyambwrxtcgwcxutlrshwjochkfkuqhdmgdsxhcnzbg", "zipyftbkgozsagzswolkztgtxzqbdpiupnsgocvispxuthxzupozexsyvhqrzwtumkuxpmyydwcphmsxukkm",
				"wixgkyjbwtelrrqjhyzphrqxrkttpofzbazzgeda", "czhlghnjfmkvytbvfohxrisfkgiungqaco", "enpykrmuzaiglrtearzzezdnwhvwoxofvxhjtytmhhjtncmxgcxgiiixydhkfsvchpvq",
				"sjurtdtunrlmrwxfksutlkfusqcqnvpzlxsvzzpswizacirvfdkeeeohoaasqzxhnnlrpxooaadtdmwum", "fzmjbkfswjggsecmqtpzuyvkjsdslavaxoxkbsjmyxobnwfvtnnohbbk",
				"joclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsar", "joclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsardyooxubjam",
				"pexnitmgzqdpyqwtzdymudxrecdvwbsikvclpjbveirobuu", "isfkgiungqacofiydxinyc", "pqhceljyrwahjbztilyulvihxyyfobcftkqkcewnqihctz",
				"wvjxeyqxalacglnqptcega", "bsoaqtnwbgvewaqdkdcdjyznizipinephiyxwbhxrmholwhdtktewbpurfnewwanbdutubl",
				"ebvoavnjoksfjoclddonddfrchczhlghnjfmkvytbvfohxrisfkgiungqacofiydxinyctphldhlsard", "fohxrisfkgiungqac", "kgiungqacofiydxinyc",
				"fiydxinyctphldhlsardyooxub", "gmwrbcatdwblmlyncloulczqggtehccxbatmkjtjwgaqifffalizqtzicgylykblfxsezwyuwvnfmyxtuywkoloeg",
				"xgbozsaswbtjolvlipdnmnkxdpqacbfunzwyoiskydycdzicfevsguhuhtnsjwyqqmzphygrstx", "ub",
				"bzmejxsmvlqffryjdasdqkygtufkbfkurczwismssiiajatrzpldcfonwkwctpgyztceulhsmdz", "oamdshdvicfwokkj",
				"rlokkopqxgqiwamcachoaykmmdxwbvtskyyvatdkawffvydegotzatizkcirfieq", "zgrypddaqsvyaovhaxkiltqniulwqrksaylkxhrrdreksqymotpphbrf",
				"hl", "l", "zamitevatvx", "xubja", "yokiqkqqvvesqtntizbejzn", "xjnyczudmupnzwemylkcaguhoqbmofzxtpbzdtvagmfiuutkmvwwyuzskkaicfgamafsrastme",
				"cejihcrsmjvbyyyokbzrbdufcxoulnfjqkvoxbgzhtltlb", "xrisfkgiungqacofiydxinyctphldhlsardyooxub", "fjocl",
				"tckecammruujrxbgkedoggxgvsmaaqmpyrguofunhejujvgaddpxbfjbg", "xrsqmktyfgpgwuyuxttfgjkqmgviyaqfbjcpsdppguybxjqipc",
				"csqsbwrbkftlwihbqruftshqhuphhpxxfpfpcvqrqztccmtruidwemaoeimitjyggmosfalfrdzqjphaatssxequg", "c",
				"aazqzihuvcxnwwhqlobwyxelweuihemxqyunokntxrvhajbkgvnpitiwgcwosmdbyzhf", "o", "ooxubjamz", "brvneoekkhbaehyktlkme",
				"tbvfohxrisfkgi", "aqqoftfpfmqswukwnhjbnzmdmmctxmfwnwsvuk", "lghnjfmkvy", "rdqsylebybzggiuqtaovsgkknjriqvcmkrxzonpvqjalpopqs",
				"ofiydxinyctphldhlsardyo", "hlghnjfmkvytbvfohxrisfkg", "lxjlveanvxtuhybljdubrtzkqlnheqjzaemwqjtkgfjyqggwhbhvglqspgkfyivljktfmj",
				"hldhlsa", "u", "zhlghnjfmkvytbvfohxrisfkgiungqacofiy", "phldhlsardyooxubj", "xrisf", "kdpccjurvgfbnxhkqgdgltfouhcyez",
				"fljqxpqgzpxjsazbrfmnxwhdnykkykzgptyahuxfrcokycqafyalhukgjhsuacxqcwzwhzeneafdizhibufh", "cfcfheeb",
				"wmhlkzejyllepomswhjsrlwvceemrunngfxwdrrhsvcfsqmibxglkjhdgwahfhcistwp", "risfkgiungqacofiydxinyctphldhlsardyoox",
				"hldhlsardyooxub", "fiydxinyctphldhlsardyo",
			},
		},
	}

	for name, fn := range functions {
		t.Run(name, func(t *testing.T) {
			for _, tc := range tests {
				t.Run(tc.name, func(t *testing.T) {
					for _, pattern := range tc.patterns {
						count := fn([]string{pattern}, tc.word)

						expected := 0
						if strings.Contains(tc.word, pattern) {
							expected = 1
						}

						if count != expected {
							t.Errorf("discrepancy for pattern '%s': expected %d, got %d", pattern, expected, count)
						}
					}
				})
			}
		})
	}
}

// goos: darwin
// goarch: arm64
// pkg: solutions/src/easy
// cpu: Apple M1 Pro
// BenchmarkNumOfStringCustomLogic
// BenchmarkNumOfStringCustomLogic-10          	   80956	     13986 ns/op
// BenchmarkNumOfStringsStdContains
// BenchmarkNumOfStringsStdContains-10         	 2193931	       548.7 ns/op
// BenchmarkNumOfStringsAhoCorasick
// BenchmarkNumOfStringsAhoCorasick-10         	  131289	      8422 ns/op
// BenchmarkNumOfStringsLargeCustom
// BenchmarkNumOfStringsLargeCustom-10         	       1	14013557250 ns/op
// BenchmarkNumOfStringsLargeStd
// BenchmarkNumOfStringsLargeStd-10            	       1	3072263708 ns/op
// BenchmarkNumOfStringsLargeAhoCorasick
// BenchmarkNumOfStringsLargeAhoCorasick-10    	       6	 188216576 ns/op

// Conclusion: Well, my "custom" solution sucks.
// The input is small (<= 100 chars) and strings.Contains
// searches substrings one by one pattern much faster
// than all-in-one loop due to vector instructions optimizations,
// though it does lots of memory scans
//
// Aho-Corasick algorithm performed well on large data
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm

func getBenchmarkData() ([]string, string) {
	word := strings.Repeat("abcdefghij", 10)

	patterns := make([]string, 100)
	for i := range 100 {
		if i%2 == 0 {
			patterns[i] = "abcdefghij" // matching pattern
		} else {
			patterns[i] = "xyz" // non-matching
		}
	}
	return patterns, word
}

func BenchmarkNumOfStringCustomLogic(b *testing.B) {
	patterns, word := getBenchmarkData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStringsCustom(patterns, word)
	}
}

func BenchmarkNumOfStringsStdContains(b *testing.B) {
	patterns, word := getBenchmarkData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStrings(patterns, word)
	}
}

func BenchmarkNumOfStringsAhoCorasick(b *testing.B) {
	patterns, word := getBenchmarkData()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStringsAhoCorasick(patterns, word)
	}
}

func getLargeBenchmarkData() ([]string, string) {
	// Generate a massive 5MB word file (simulating a huge network packet or log stream)
	word := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 200_000)

	// Generate 2,000 distinct patterns to track simultaneously
	patterns := make([]string, 2000)
	for i := range 2000 {
		// A mix of matching strings and non-matching anomalies
		if i%2 == 0 {
			patterns[i] = fmt.Sprintf("xyzabc%d", i) // likely won't match, or matches deep
		} else {
			patterns[i] = "defgh" // Matches frequently throughout
		}
	}

	return patterns, word
}

func BenchmarkNumOfStringsLargeCustom(b *testing.B) {
	patterns, word := getLargeBenchmarkData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStringsCustom(patterns, word)
	}
}

func BenchmarkNumOfStringsLargeStd(b *testing.B) {
	patterns, word := getLargeBenchmarkData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStrings(patterns, word)
	}
}

func BenchmarkNumOfStringsLargeAhoCorasick(b *testing.B) {
	patterns, word := getLargeBenchmarkData()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = numOfStringsAhoCorasick(patterns, word)
	}
}
