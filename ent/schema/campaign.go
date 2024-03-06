package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Campaign holds the schema definition for the Campaign entity.
type Campaign struct {
	ent.Schema
}

// Fields of the Cards.
func (Campaign) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("current_slot"),
		field.Time("start_day"),
		field.Time("end_day"),
		field.Int64("total_slot"),
		field.Bool("is_full"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Campaign.
func (Campaign) Edges() []ent.Edge {
	return nil
}
