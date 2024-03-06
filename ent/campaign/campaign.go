// Code generated by ent, DO NOT EDIT.

package campaign

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the campaign type in the database.
	Label = "campaign"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCurrentSlot holds the string denoting the current_slot field in the database.
	FieldCurrentSlot = "current_slot"
	// FieldStartDay holds the string denoting the start_day field in the database.
	FieldStartDay = "start_day"
	// FieldEndDay holds the string denoting the end_day field in the database.
	FieldEndDay = "end_day"
	// FieldTotalSlot holds the string denoting the total_slot field in the database.
	FieldTotalSlot = "total_slot"
	// FieldIsFull holds the string denoting the is_full field in the database.
	FieldIsFull = "is_full"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the campaign in the database.
	Table = "campaigns"
)

// Columns holds all SQL columns for campaign fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCurrentSlot,
	FieldStartDay,
	FieldEndDay,
	FieldTotalSlot,
	FieldIsFull,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Campaign queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCurrentSlot orders the results by the current_slot field.
func ByCurrentSlot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentSlot, opts...).ToFunc()
}

// ByStartDay orders the results by the start_day field.
func ByStartDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDay, opts...).ToFunc()
}

// ByEndDay orders the results by the end_day field.
func ByEndDay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDay, opts...).ToFunc()
}

// ByTotalSlot orders the results by the total_slot field.
func ByTotalSlot(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalSlot, opts...).ToFunc()
}

// ByIsFull orders the results by the is_full field.
func ByIsFull(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsFull, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
