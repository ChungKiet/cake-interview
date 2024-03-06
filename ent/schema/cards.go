package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Cards.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("type"),
		field.Float("value"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return nil
}
