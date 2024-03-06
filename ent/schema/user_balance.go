package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// UserBalance holds the schema definition for the UserBalance entity.
type UserBalance struct {
	ent.Schema
}

// Fields of the UserBalance.
func (UserBalance) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id"),
		field.Float("balance"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the UserBalance.
func (UserBalance) Edges() []ent.Edge {
	return nil
}
