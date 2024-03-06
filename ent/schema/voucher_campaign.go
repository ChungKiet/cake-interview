package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// VoucherCampaign holds the schema definition for the VoucherCampaign entity.
type VoucherCampaign struct {
	ent.Schema
}

// Fields of the Voucher.
func (VoucherCampaign) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("campaign_id"),
		field.Int64("voucher_id"),
		field.Time("start_day"),
		field.Time("valid_to"),
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the VoucherCampaign.
func (VoucherCampaign) Edges() []ent.Edge {
	return nil
}
