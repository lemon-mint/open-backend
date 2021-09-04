// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/lemon-mint/open-backend/ent/schema"
	"github.com/lemon-mint/open-backend/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAlgorithm is the schema descriptor for algorithm field.
	userDescAlgorithm := userFields[4].Descriptor()
	// user.DefaultAlgorithm holds the default value on creation for the algorithm field.
	user.DefaultAlgorithm = userDescAlgorithm.Default.(int32)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[5].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
}
