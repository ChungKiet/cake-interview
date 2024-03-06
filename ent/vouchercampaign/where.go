// Code generated by ent, DO NOT EDIT.

package vouchercampaign

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ChungKiet/cake-interview/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldID, id))
}

// CampaignID applies equality check predicate on the "campaign_id" field. It's identical to CampaignIDEQ.
func CampaignID(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldCampaignID, v))
}

// VoucherID applies equality check predicate on the "voucher_id" field. It's identical to VoucherIDEQ.
func VoucherID(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldVoucherID, v))
}

// StartDay applies equality check predicate on the "start_day" field. It's identical to StartDayEQ.
func StartDay(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldStartDay, v))
}

// ValidTo applies equality check predicate on the "valid_to" field. It's identical to ValidToEQ.
func ValidTo(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldValidTo, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// CampaignIDEQ applies the EQ predicate on the "campaign_id" field.
func CampaignIDEQ(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldCampaignID, v))
}

// CampaignIDNEQ applies the NEQ predicate on the "campaign_id" field.
func CampaignIDNEQ(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldCampaignID, v))
}

// CampaignIDIn applies the In predicate on the "campaign_id" field.
func CampaignIDIn(vs ...int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldCampaignID, vs...))
}

// CampaignIDNotIn applies the NotIn predicate on the "campaign_id" field.
func CampaignIDNotIn(vs ...int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldCampaignID, vs...))
}

// CampaignIDGT applies the GT predicate on the "campaign_id" field.
func CampaignIDGT(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldCampaignID, v))
}

// CampaignIDGTE applies the GTE predicate on the "campaign_id" field.
func CampaignIDGTE(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldCampaignID, v))
}

// CampaignIDLT applies the LT predicate on the "campaign_id" field.
func CampaignIDLT(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldCampaignID, v))
}

// CampaignIDLTE applies the LTE predicate on the "campaign_id" field.
func CampaignIDLTE(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldCampaignID, v))
}

// VoucherIDEQ applies the EQ predicate on the "voucher_id" field.
func VoucherIDEQ(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldVoucherID, v))
}

// VoucherIDNEQ applies the NEQ predicate on the "voucher_id" field.
func VoucherIDNEQ(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldVoucherID, v))
}

// VoucherIDIn applies the In predicate on the "voucher_id" field.
func VoucherIDIn(vs ...int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldVoucherID, vs...))
}

// VoucherIDNotIn applies the NotIn predicate on the "voucher_id" field.
func VoucherIDNotIn(vs ...int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldVoucherID, vs...))
}

// VoucherIDGT applies the GT predicate on the "voucher_id" field.
func VoucherIDGT(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldVoucherID, v))
}

// VoucherIDGTE applies the GTE predicate on the "voucher_id" field.
func VoucherIDGTE(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldVoucherID, v))
}

// VoucherIDLT applies the LT predicate on the "voucher_id" field.
func VoucherIDLT(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldVoucherID, v))
}

// VoucherIDLTE applies the LTE predicate on the "voucher_id" field.
func VoucherIDLTE(v int64) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldVoucherID, v))
}

// StartDayEQ applies the EQ predicate on the "start_day" field.
func StartDayEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldStartDay, v))
}

// StartDayNEQ applies the NEQ predicate on the "start_day" field.
func StartDayNEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldStartDay, v))
}

// StartDayIn applies the In predicate on the "start_day" field.
func StartDayIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldStartDay, vs...))
}

// StartDayNotIn applies the NotIn predicate on the "start_day" field.
func StartDayNotIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldStartDay, vs...))
}

// StartDayGT applies the GT predicate on the "start_day" field.
func StartDayGT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldStartDay, v))
}

// StartDayGTE applies the GTE predicate on the "start_day" field.
func StartDayGTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldStartDay, v))
}

// StartDayLT applies the LT predicate on the "start_day" field.
func StartDayLT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldStartDay, v))
}

// StartDayLTE applies the LTE predicate on the "start_day" field.
func StartDayLTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldStartDay, v))
}

// ValidToEQ applies the EQ predicate on the "valid_to" field.
func ValidToEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldValidTo, v))
}

// ValidToNEQ applies the NEQ predicate on the "valid_to" field.
func ValidToNEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldValidTo, v))
}

// ValidToIn applies the In predicate on the "valid_to" field.
func ValidToIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldValidTo, vs...))
}

// ValidToNotIn applies the NotIn predicate on the "valid_to" field.
func ValidToNotIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldValidTo, vs...))
}

// ValidToGT applies the GT predicate on the "valid_to" field.
func ValidToGT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldValidTo, v))
}

// ValidToGTE applies the GTE predicate on the "valid_to" field.
func ValidToGTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldValidTo, v))
}

// ValidToLT applies the LT predicate on the "valid_to" field.
func ValidToLT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldValidTo, v))
}

// ValidToLTE applies the LTE predicate on the "valid_to" field.
func ValidToLTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldValidTo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VoucherCampaign) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VoucherCampaign) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VoucherCampaign) predicate.VoucherCampaign {
	return predicate.VoucherCampaign(sql.NotPredicates(p))
}
