//
// Package ctxvaluetype shows that context key type is checked when we are
// trying to retrieve it later.
//
// "context" package will use equal operator `==` to compare keys, and it will
// compare key types also, this will result in false result even if the types
// are aliases.
//
package ctxvaluetype

const strKey string = "strKey"

type ctxKey string

var typedKey ctxKey = "typeKey"
