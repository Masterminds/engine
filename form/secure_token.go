package form

import (
	"github.com/aokoli/goutils"
)

// The name of the token that is automatically placed into a form.
//
// This should never be modified after a form has been generated.
var SecureTokenName = "__token__"

// The length of the security token.
var SecurityTokenLength = 32

// Generate a security token.
//
// This uses SecurityTokenLength to determine the appropriate length. If
// that value is <= 0, this will panic.
func SecurityToken() string {
	tok, err := goutils.RandomAlphaNumeric(SecurityTokenLength)
	if err != nil {
		// We panic because the only way to error out here is to pass in an
		// illegal security token length.
		panic(err)
	}
	return tok
}

// SecurityField returns a Hidden form element initialized with a security
// token.
func SecurityField() Hidden {
	return Hidden{
		Name:  SecureTokenName,
		Value: SecurityToken(),
	}
}
