package mapkeyasstruct

import "testing"

func BenchmarkStringKey(b *testing.B) {
	testObj := make(A)

	testObj["A"] = 0
	testObj["B"] = 1
	testObj["C"] = 2
	testObj["D"] = 3
	testObj["E"] = 4
	testObj["F"] = 5

	for n := 0; n < b.N; n++ {
		_ = testObj["A"]
		_ = testObj["F"]
	}
}

func BenchmarkStructKey(b *testing.B) {
	testObj := make(B)

	testObj[key{"A", 1}] = 0
	testObj[key{"B", 2}] = 1
	testObj[key{"C", 3}] = 2
	testObj[key{"D", 4}] = 3
	testObj[key{"E", 5}] = 4
	testObj[key{"F", 6}] = 5

	for n := 0; n < b.N; n++ {
		_ = testObj[key{"A", 1}]
		_ = testObj[key{"F", 6}]
	}
}
