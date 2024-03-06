// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ChungKiet/cake-interview/ent/campaign"
)

// Campaign is the model entity for the Campaign schema.
type Campaign struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// CurrentSlot holds the value of the "current_slot" field.
	CurrentSlot int64 `json:"current_slot,omitempty"`
	// StartDay holds the value of the "start_day" field.
	StartDay time.Time `json:"start_day,omitempty"`
	// EndDay holds the value of the "end_day" field.
	EndDay time.Time `json:"end_day,omitempty"`
	// TotalSlot holds the value of the "total_slot" field.
	TotalSlot int64 `json:"total_slot,omitempty"`
	// IsFull holds the value of the "is_full" field.
	IsFull bool `json:"is_full,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Campaign) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case campaign.FieldIsFull:
			values[i] = new(sql.NullBool)
		case campaign.FieldID, campaign.FieldCurrentSlot, campaign.FieldTotalSlot:
			values[i] = new(sql.NullInt64)
		case campaign.FieldName:
			values[i] = new(sql.NullString)
		case campaign.FieldStartDay, campaign.FieldEndDay, campaign.FieldCreatedAt, campaign.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Campaign fields.
func (c *Campaign) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case campaign.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case campaign.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case campaign.FieldCurrentSlot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field current_slot", values[i])
			} else if value.Valid {
				c.CurrentSlot = value.Int64
			}
		case campaign.FieldStartDay:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_day", values[i])
			} else if value.Valid {
				c.StartDay = value.Time
			}
		case campaign.FieldEndDay:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_day", values[i])
			} else if value.Valid {
				c.EndDay = value.Time
			}
		case campaign.FieldTotalSlot:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_slot", values[i])
			} else if value.Valid {
				c.TotalSlot = value.Int64
			}
		case campaign.FieldIsFull:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_full", values[i])
			} else if value.Valid {
				c.IsFull = value.Bool
			}
		case campaign.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case campaign.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Campaign.
// This includes values selected through modifiers, order, etc.
func (c *Campaign) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Campaign.
// Note that you need to call Campaign.Unwrap() before calling this method if this Campaign
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Campaign) Update() *CampaignUpdateOne {
	return NewCampaignClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Campaign entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Campaign) Unwrap() *Campaign {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Campaign is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Campaign) String() string {
	var builder strings.Builder
	builder.WriteString("Campaign(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("current_slot=")
	builder.WriteString(fmt.Sprintf("%v", c.CurrentSlot))
	builder.WriteString(", ")
	builder.WriteString("start_day=")
	builder.WriteString(c.StartDay.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_day=")
	builder.WriteString(c.EndDay.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("total_slot=")
	builder.WriteString(fmt.Sprintf("%v", c.TotalSlot))
	builder.WriteString(", ")
	builder.WriteString("is_full=")
	builder.WriteString(fmt.Sprintf("%v", c.IsFull))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Campaigns is a parsable slice of Campaign.
type Campaigns []*Campaign