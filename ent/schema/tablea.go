package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	aEntity "github.com/washanhanzi/ent-same-package-name/a/entity"
)

// TableA holds the schema definition for the TableA entity.
type TableA struct {
	ent.Schema
}

// Fields of the TableA.
func (TableA) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("a", aEntity.A{}),
	}
}

// Edges of the TableA.
func (TableA) Edges() []ent.Edge {
	return nil
}
