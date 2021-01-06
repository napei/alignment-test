package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Test holds the schema definition for the Test entity.
type Test struct {
	ent.Schema
}

// Fields of the Test.
func (Test) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("field1"),
		field.Float("field2"),
		field.Int32("field3"),
	}
}

// Edges of the Test.
func (Test) Edges() []ent.Edge {
	return nil
}
