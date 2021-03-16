//
// Package structalloc demonstrates that order of fields in a structure
// directly affects how much memory will be allocated for an instance of this structure.
//
// This is caused by https://en.wikipedia.org/wiki/Data_structure_alignment
//
package structalloc

// A struct will allocate 24 bytes
type A struct {
	Bool  bool  // 1 byte				[X-------] 8 bytes
	Int64 int64 // 8 bytes				[XXXXXXXX] 8 bytes
	Int32 int32 // 4 bytes				[XXXX----] 8 bytes
}

// B struct will allocate 16 bytes
type B struct {
	Int64 int64 // 8 bytes				[XXXXXXXX] 8 bytes
	Bool  bool  // 1 byte				[X---XXXX] 8 bytes
	Int32 int32 // 4 bytes
}
