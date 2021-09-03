// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldSalt holds the string denoting the salt field in the database.
	FieldSalt = "salt"
	// FieldAlgorithm holds the string denoting the algorithm field in the database.
	FieldAlgorithm = "algorithm"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldUsername,
	FieldHash,
	FieldSalt,
	FieldAlgorithm,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultAlgorithm holds the default value on creation for the "algorithm" field.
	DefaultAlgorithm int32
)