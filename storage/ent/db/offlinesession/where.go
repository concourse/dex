// Code generated by ent, DO NOT EDIT.

package offlinesession

import (
	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldContainsFold(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldUserID, v))
}

// ConnID applies equality check predicate on the "conn_id" field. It's identical to ConnIDEQ.
func ConnID(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldConnID, v))
}

// Refresh applies equality check predicate on the "refresh" field. It's identical to RefreshEQ.
func Refresh(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldRefresh, v))
}

// ConnectorData applies equality check predicate on the "connector_data" field. It's identical to ConnectorDataEQ.
func ConnectorData(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldConnectorData, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldContainsFold(FieldUserID, v))
}

// ConnIDEQ applies the EQ predicate on the "conn_id" field.
func ConnIDEQ(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldConnID, v))
}

// ConnIDNEQ applies the NEQ predicate on the "conn_id" field.
func ConnIDNEQ(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNEQ(FieldConnID, v))
}

// ConnIDIn applies the In predicate on the "conn_id" field.
func ConnIDIn(vs ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIn(FieldConnID, vs...))
}

// ConnIDNotIn applies the NotIn predicate on the "conn_id" field.
func ConnIDNotIn(vs ...string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotIn(FieldConnID, vs...))
}

// ConnIDGT applies the GT predicate on the "conn_id" field.
func ConnIDGT(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGT(FieldConnID, v))
}

// ConnIDGTE applies the GTE predicate on the "conn_id" field.
func ConnIDGTE(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGTE(FieldConnID, v))
}

// ConnIDLT applies the LT predicate on the "conn_id" field.
func ConnIDLT(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLT(FieldConnID, v))
}

// ConnIDLTE applies the LTE predicate on the "conn_id" field.
func ConnIDLTE(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLTE(FieldConnID, v))
}

// ConnIDContains applies the Contains predicate on the "conn_id" field.
func ConnIDContains(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldContains(FieldConnID, v))
}

// ConnIDHasPrefix applies the HasPrefix predicate on the "conn_id" field.
func ConnIDHasPrefix(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldHasPrefix(FieldConnID, v))
}

// ConnIDHasSuffix applies the HasSuffix predicate on the "conn_id" field.
func ConnIDHasSuffix(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldHasSuffix(FieldConnID, v))
}

// ConnIDEqualFold applies the EqualFold predicate on the "conn_id" field.
func ConnIDEqualFold(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEqualFold(FieldConnID, v))
}

// ConnIDContainsFold applies the ContainsFold predicate on the "conn_id" field.
func ConnIDContainsFold(v string) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldContainsFold(FieldConnID, v))
}

// RefreshEQ applies the EQ predicate on the "refresh" field.
func RefreshEQ(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldRefresh, v))
}

// RefreshNEQ applies the NEQ predicate on the "refresh" field.
func RefreshNEQ(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNEQ(FieldRefresh, v))
}

// RefreshIn applies the In predicate on the "refresh" field.
func RefreshIn(vs ...[]byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIn(FieldRefresh, vs...))
}

// RefreshNotIn applies the NotIn predicate on the "refresh" field.
func RefreshNotIn(vs ...[]byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotIn(FieldRefresh, vs...))
}

// RefreshGT applies the GT predicate on the "refresh" field.
func RefreshGT(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGT(FieldRefresh, v))
}

// RefreshGTE applies the GTE predicate on the "refresh" field.
func RefreshGTE(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGTE(FieldRefresh, v))
}

// RefreshLT applies the LT predicate on the "refresh" field.
func RefreshLT(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLT(FieldRefresh, v))
}

// RefreshLTE applies the LTE predicate on the "refresh" field.
func RefreshLTE(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLTE(FieldRefresh, v))
}

// ConnectorDataEQ applies the EQ predicate on the "connector_data" field.
func ConnectorDataEQ(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldEQ(FieldConnectorData, v))
}

// ConnectorDataNEQ applies the NEQ predicate on the "connector_data" field.
func ConnectorDataNEQ(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNEQ(FieldConnectorData, v))
}

// ConnectorDataIn applies the In predicate on the "connector_data" field.
func ConnectorDataIn(vs ...[]byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIn(FieldConnectorData, vs...))
}

// ConnectorDataNotIn applies the NotIn predicate on the "connector_data" field.
func ConnectorDataNotIn(vs ...[]byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotIn(FieldConnectorData, vs...))
}

// ConnectorDataGT applies the GT predicate on the "connector_data" field.
func ConnectorDataGT(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGT(FieldConnectorData, v))
}

// ConnectorDataGTE applies the GTE predicate on the "connector_data" field.
func ConnectorDataGTE(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldGTE(FieldConnectorData, v))
}

// ConnectorDataLT applies the LT predicate on the "connector_data" field.
func ConnectorDataLT(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLT(FieldConnectorData, v))
}

// ConnectorDataLTE applies the LTE predicate on the "connector_data" field.
func ConnectorDataLTE(v []byte) predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldLTE(FieldConnectorData, v))
}

// ConnectorDataIsNil applies the IsNil predicate on the "connector_data" field.
func ConnectorDataIsNil() predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldIsNull(FieldConnectorData))
}

// ConnectorDataNotNil applies the NotNil predicate on the "connector_data" field.
func ConnectorDataNotNil() predicate.OfflineSession {
	return predicate.OfflineSession(sql.FieldNotNull(FieldConnectorData))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OfflineSession) predicate.OfflineSession {
	return predicate.OfflineSession(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OfflineSession) predicate.OfflineSession {
	return predicate.OfflineSession(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OfflineSession) predicate.OfflineSession {
	return predicate.OfflineSession(sql.NotPredicates(p))
}
