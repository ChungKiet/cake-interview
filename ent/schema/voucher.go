package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Voucher holds the schema definition for the Voucher entity.
type Voucher struct {
	ent.Schema
}

// Fields of the Voucher.
func (Voucher) Fields() []ent.Field {
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

// Edges of the Voucher.
func (Voucher) Edges() []ent.Edge {
	return nil
}
