// Code generated by ent, DO NOT EDIT.

package campaign

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ChungKiet/cake-interview/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldName, v))
}

// CurrentSlot applies equality check predicate on the "current_slot" field. It's identical to CurrentSlotEQ.
func CurrentSlot(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCurrentSlot, v))
}

// StartDay applies equality check predicate on the "start_day" field. It's identical to StartDayEQ.
func StartDay(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStartDay, v))
}

// EndDay applies equality check predicate on the "end_day" field. It's identical to EndDayEQ.
func EndDay(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldEndDay, v))
}

// TotalSlot applies equality check predicate on the "total_slot" field. It's identical to TotalSlotEQ.
func TotalSlot(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldTotalSlot, v))
}

// IsFull applies equality check predicate on the "is_full" field. It's identical to IsFullEQ.
func IsFull(v bool) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldIsFull, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Campaign {
	return predicate.Campaign(sql.FieldContainsFold(FieldName, v))
}

// CurrentSlotEQ applies the EQ predicate on the "current_slot" field.
func CurrentSlotEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCurrentSlot, v))
}

// CurrentSlotNEQ applies the NEQ predicate on the "current_slot" field.
func CurrentSlotNEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldCurrentSlot, v))
}

// CurrentSlotIn applies the In predicate on the "current_slot" field.
func CurrentSlotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldCurrentSlot, vs...))
}

// CurrentSlotNotIn applies the NotIn predicate on the "current_slot" field.
func CurrentSlotNotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldCurrentSlot, vs...))
}

// CurrentSlotGT applies the GT predicate on the "current_slot" field.
func CurrentSlotGT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldCurrentSlot, v))
}

// CurrentSlotGTE applies the GTE predicate on the "current_slot" field.
func CurrentSlotGTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldCurrentSlot, v))
}

// CurrentSlotLT applies the LT predicate on the "current_slot" field.
func CurrentSlotLT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldCurrentSlot, v))
}

// CurrentSlotLTE applies the LTE predicate on the "current_slot" field.
func CurrentSlotLTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldCurrentSlot, v))
}

// StartDayEQ applies the EQ predicate on the "start_day" field.
func StartDayEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldStartDay, v))
}

// StartDayNEQ applies the NEQ predicate on the "start_day" field.
func StartDayNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldStartDay, v))
}

// StartDayIn applies the In predicate on the "start_day" field.
func StartDayIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldStartDay, vs...))
}

// StartDayNotIn applies the NotIn predicate on the "start_day" field.
func StartDayNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldStartDay, vs...))
}

// StartDayGT applies the GT predicate on the "start_day" field.
func StartDayGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldStartDay, v))
}

// StartDayGTE applies the GTE predicate on the "start_day" field.
func StartDayGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldStartDay, v))
}

// StartDayLT applies the LT predicate on the "start_day" field.
func StartDayLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldStartDay, v))
}

// StartDayLTE applies the LTE predicate on the "start_day" field.
func StartDayLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldStartDay, v))
}

// EndDayEQ applies the EQ predicate on the "end_day" field.
func EndDayEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldEndDay, v))
}

// EndDayNEQ applies the NEQ predicate on the "end_day" field.
func EndDayNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldEndDay, v))
}

// EndDayIn applies the In predicate on the "end_day" field.
func EndDayIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldEndDay, vs...))
}

// EndDayNotIn applies the NotIn predicate on the "end_day" field.
func EndDayNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldEndDay, vs...))
}

// EndDayGT applies the GT predicate on the "end_day" field.
func EndDayGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldEndDay, v))
}

// EndDayGTE applies the GTE predicate on the "end_day" field.
func EndDayGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldEndDay, v))
}

// EndDayLT applies the LT predicate on the "end_day" field.
func EndDayLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldEndDay, v))
}

// EndDayLTE applies the LTE predicate on the "end_day" field.
func EndDayLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldEndDay, v))
}

// TotalSlotEQ applies the EQ predicate on the "total_slot" field.
func TotalSlotEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldTotalSlot, v))
}

// TotalSlotNEQ applies the NEQ predicate on the "total_slot" field.
func TotalSlotNEQ(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldTotalSlot, v))
}

// TotalSlotIn applies the In predicate on the "total_slot" field.
func TotalSlotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldTotalSlot, vs...))
}

// TotalSlotNotIn applies the NotIn predicate on the "total_slot" field.
func TotalSlotNotIn(vs ...int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldTotalSlot, vs...))
}

// TotalSlotGT applies the GT predicate on the "total_slot" field.
func TotalSlotGT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldTotalSlot, v))
}

// TotalSlotGTE applies the GTE predicate on the "total_slot" field.
func TotalSlotGTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldTotalSlot, v))
}

// TotalSlotLT applies the LT predicate on the "total_slot" field.
func TotalSlotLT(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldTotalSlot, v))
}

// TotalSlotLTE applies the LTE predicate on the "total_slot" field.
func TotalSlotLTE(v int64) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldTotalSlot, v))
}

// IsFullEQ applies the EQ predicate on the "is_full" field.
func IsFullEQ(v bool) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldIsFull, v))
}

// IsFullNEQ applies the NEQ predicate on the "is_full" field.
func IsFullNEQ(v bool) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldIsFull, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Campaign {
	return predicate.Campaign(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Campaign) predicate.Campaign {
	return predicate.Campaign(sql.NotPredicates(p))
}
