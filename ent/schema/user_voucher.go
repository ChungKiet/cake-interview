package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// UserVoucher holds the schema definition for the UserVoucher entity.
type UserVoucher struct {
	ent.Schema
}

// Fields of the UserVoucher.
func (UserVoucher) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("voucher_campaign_id"),
		field.String("user_id"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the UserVoucher.
func (UserVoucher) Edges() []ent.Edge {
	return nil
}
