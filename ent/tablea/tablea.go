// Code generated by entc, DO NOT EDIT.

package tablea

const (
	// Label holds the string label denoting the tablea type in the database.
	Label = "tablea"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldA holds the string denoting the a field in the database.
	FieldA = "a"
	// Table holds the table name of the tablea in the database.
	Table = "table_as"
)

// Columns holds all SQL columns for tablea fields.
var Columns = []string{
	FieldID,
	FieldA,
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
