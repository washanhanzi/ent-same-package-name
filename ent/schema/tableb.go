package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	bEntity "github.com/washanhanzi/ent-same-package-name/b/entity"
)

// TableB holds the schema definition for the TableB entity.
type TableB struct {
	ent.Schema
}

// Fields of the TableB.
func (TableB) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("b", bEntity.B{}),
	}
}

// Edges of the TableB.
func (TableB) Edges() []ent.Edge {
	return nil
}
