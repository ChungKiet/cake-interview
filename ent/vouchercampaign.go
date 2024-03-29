// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ChungKiet/cake-interview/ent/vouchercampaign"
)

// VoucherCampaign is the model entity for the VoucherCampaign schema.
type VoucherCampaign struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CampaignID holds the value of the "campaign_id" field.
	CampaignID int64 `json:"campaign_id,omitempty"`
	// VoucherID holds the value of the "voucher_id" field.
	VoucherID int64 `json:"voucher_id,omitempty"`
	// StartDay holds the value of the "start_day" field.
	StartDay time.Time `json:"start_day,omitempty"`
	// ValidTo holds the value of the "valid_to" field.
	ValidTo time.Time `json:"valid_to,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VoucherCampaign) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vouchercampaign.FieldID, vouchercampaign.FieldCampaignID, vouchercampaign.FieldVoucherID:
			values[i] = new(sql.NullInt64)
		case vouchercampaign.FieldStartDay, vouchercampaign.FieldValidTo, vouchercampaign.FieldCreatedAt, vouchercampaign.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VoucherCampaign fields.
func (vc *VoucherCampaign) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vouchercampaign.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vc.ID = int(value.Int64)
		case vouchercampaign.FieldCampaignID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field campaign_id", values[i])
			} else if value.Valid {
				vc.CampaignID = value.Int64
			}
		case vouchercampaign.FieldVoucherID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field voucher_id", values[i])
			} else if value.Valid {
				vc.VoucherID = value.Int64
			}
		case vouchercampaign.FieldStartDay:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_day", values[i])
			} else if value.Valid {
				vc.StartDay = value.Time
			}
		case vouchercampaign.FieldValidTo:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field valid_to", values[i])
			} else if value.Valid {
				vc.ValidTo = value.Time
			}
		case vouchercampaign.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vc.CreatedAt = value.Time
			}
		case vouchercampaign.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vc.UpdatedAt = value.Time
			}
		default:
			vc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VoucherCampaign.
// This includes values selected through modifiers, order, etc.
func (vc *VoucherCampaign) Value(name string) (ent.Value, error) {
	return vc.selectValues.Get(name)
}

// Update returns a builder for updating this VoucherCampaign.
// Note that you need to call VoucherCampaign.Unwrap() before calling this method if this VoucherCampaign
// was returned from a transaction, and the transaction was committed or rolled back.
func (vc *VoucherCampaign) Update() *VoucherCampaignUpdateOne {
	return NewVoucherCampaignClient(vc.config).UpdateOne(vc)
}

// Unwrap unwraps the VoucherCampaign entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vc *VoucherCampaign) Unwrap() *VoucherCampaign {
	_tx, ok := vc.config.driver.(*txDriver)
	if !ok {
		panic("ent: VoucherCampaign is not a transactional entity")
	}
	vc.config.driver = _tx.drv
	return vc
}

// String implements the fmt.Stringer.
func (vc *VoucherCampaign) String() string {
	var builder strings.Builder
	builder.WriteString("VoucherCampaign(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vc.ID))
	builder.WriteString("campaign_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.CampaignID))
	builder.WriteString(", ")
	builder.WriteString("voucher_id=")
	builder.WriteString(fmt.Sprintf("%v", vc.VoucherID))
	builder.WriteString(", ")
	builder.WriteString("start_day=")
	builder.WriteString(vc.StartDay.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("valid_to=")
	builder.WriteString(vc.ValidTo.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(vc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vc.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// VoucherCampaigns is a parsable slice of VoucherCampaign.
type VoucherCampaigns []*VoucherCampaign
