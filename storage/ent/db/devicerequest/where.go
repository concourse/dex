// Code generated by ent, DO NOT EDIT.

package devicerequest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserCode applies equality check predicate on the "user_code" field. It's identical to UserCodeEQ.
func UserCode(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserCode), v))
	})
}

// DeviceCode applies equality check predicate on the "device_code" field. It's identical to DeviceCodeEQ.
func DeviceCode(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceCode), v))
	})
}

// ClientID applies equality check predicate on the "client_id" field. It's identical to ClientIDEQ.
func ClientID(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientSecret applies equality check predicate on the "client_secret" field. It's identical to ClientSecretEQ.
func ClientSecret(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// UserCodeEQ applies the EQ predicate on the "user_code" field.
func UserCodeEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserCode), v))
	})
}

// UserCodeNEQ applies the NEQ predicate on the "user_code" field.
func UserCodeNEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserCode), v))
	})
}

// UserCodeIn applies the In predicate on the "user_code" field.
func UserCodeIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserCode), v...))
	})
}

// UserCodeNotIn applies the NotIn predicate on the "user_code" field.
func UserCodeNotIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserCode), v...))
	})
}

// UserCodeGT applies the GT predicate on the "user_code" field.
func UserCodeGT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserCode), v))
	})
}

// UserCodeGTE applies the GTE predicate on the "user_code" field.
func UserCodeGTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserCode), v))
	})
}

// UserCodeLT applies the LT predicate on the "user_code" field.
func UserCodeLT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserCode), v))
	})
}

// UserCodeLTE applies the LTE predicate on the "user_code" field.
func UserCodeLTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserCode), v))
	})
}

// UserCodeContains applies the Contains predicate on the "user_code" field.
func UserCodeContains(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserCode), v))
	})
}

// UserCodeHasPrefix applies the HasPrefix predicate on the "user_code" field.
func UserCodeHasPrefix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserCode), v))
	})
}

// UserCodeHasSuffix applies the HasSuffix predicate on the "user_code" field.
func UserCodeHasSuffix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserCode), v))
	})
}

// UserCodeEqualFold applies the EqualFold predicate on the "user_code" field.
func UserCodeEqualFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserCode), v))
	})
}

// UserCodeContainsFold applies the ContainsFold predicate on the "user_code" field.
func UserCodeContainsFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserCode), v))
	})
}

// DeviceCodeEQ applies the EQ predicate on the "device_code" field.
func DeviceCodeEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeNEQ applies the NEQ predicate on the "device_code" field.
func DeviceCodeNEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeIn applies the In predicate on the "device_code" field.
func DeviceCodeIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeviceCode), v...))
	})
}

// DeviceCodeNotIn applies the NotIn predicate on the "device_code" field.
func DeviceCodeNotIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeviceCode), v...))
	})
}

// DeviceCodeGT applies the GT predicate on the "device_code" field.
func DeviceCodeGT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeGTE applies the GTE predicate on the "device_code" field.
func DeviceCodeGTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeLT applies the LT predicate on the "device_code" field.
func DeviceCodeLT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeLTE applies the LTE predicate on the "device_code" field.
func DeviceCodeLTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeContains applies the Contains predicate on the "device_code" field.
func DeviceCodeContains(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeHasPrefix applies the HasPrefix predicate on the "device_code" field.
func DeviceCodeHasPrefix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeHasSuffix applies the HasSuffix predicate on the "device_code" field.
func DeviceCodeHasSuffix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeEqualFold applies the EqualFold predicate on the "device_code" field.
func DeviceCodeEqualFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeContainsFold applies the ContainsFold predicate on the "device_code" field.
func DeviceCodeContainsFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDeviceCode), v))
	})
}

// ClientIDEQ applies the EQ predicate on the "client_id" field.
func ClientIDEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientID), v))
	})
}

// ClientIDNEQ applies the NEQ predicate on the "client_id" field.
func ClientIDNEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientID), v))
	})
}

// ClientIDIn applies the In predicate on the "client_id" field.
func ClientIDIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientID), v...))
	})
}

// ClientIDNotIn applies the NotIn predicate on the "client_id" field.
func ClientIDNotIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientID), v...))
	})
}

// ClientIDGT applies the GT predicate on the "client_id" field.
func ClientIDGT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientID), v))
	})
}

// ClientIDGTE applies the GTE predicate on the "client_id" field.
func ClientIDGTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientID), v))
	})
}

// ClientIDLT applies the LT predicate on the "client_id" field.
func ClientIDLT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientID), v))
	})
}

// ClientIDLTE applies the LTE predicate on the "client_id" field.
func ClientIDLTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientID), v))
	})
}

// ClientIDContains applies the Contains predicate on the "client_id" field.
func ClientIDContains(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientID), v))
	})
}

// ClientIDHasPrefix applies the HasPrefix predicate on the "client_id" field.
func ClientIDHasPrefix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientID), v))
	})
}

// ClientIDHasSuffix applies the HasSuffix predicate on the "client_id" field.
func ClientIDHasSuffix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientID), v))
	})
}

// ClientIDEqualFold applies the EqualFold predicate on the "client_id" field.
func ClientIDEqualFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientID), v))
	})
}

// ClientIDContainsFold applies the ContainsFold predicate on the "client_id" field.
func ClientIDContainsFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientID), v))
	})
}

// ClientSecretEQ applies the EQ predicate on the "client_secret" field.
func ClientSecretEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretNEQ applies the NEQ predicate on the "client_secret" field.
func ClientSecretNEQ(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientSecret), v))
	})
}

// ClientSecretIn applies the In predicate on the "client_secret" field.
func ClientSecretIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretNotIn applies the NotIn predicate on the "client_secret" field.
func ClientSecretNotIn(vs ...string) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientSecret), v...))
	})
}

// ClientSecretGT applies the GT predicate on the "client_secret" field.
func ClientSecretGT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretGTE applies the GTE predicate on the "client_secret" field.
func ClientSecretGTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLT applies the LT predicate on the "client_secret" field.
func ClientSecretLT(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientSecret), v))
	})
}

// ClientSecretLTE applies the LTE predicate on the "client_secret" field.
func ClientSecretLTE(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContains applies the Contains predicate on the "client_secret" field.
func ClientSecretContains(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasPrefix applies the HasPrefix predicate on the "client_secret" field.
func ClientSecretHasPrefix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretHasSuffix applies the HasSuffix predicate on the "client_secret" field.
func ClientSecretHasSuffix(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientSecret), v))
	})
}

// ClientSecretEqualFold applies the EqualFold predicate on the "client_secret" field.
func ClientSecretEqualFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientSecret), v))
	})
}

// ClientSecretContainsFold applies the ContainsFold predicate on the "client_secret" field.
func ClientSecretContainsFold(v string) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientSecret), v))
	})
}

// ScopesIsNil applies the IsNil predicate on the "scopes" field.
func ScopesIsNil() predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScopes)))
	})
}

// ScopesNotNil applies the NotNil predicate on the "scopes" field.
func ScopesNotNil() predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScopes)))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.DeviceRequest {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceRequest) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceRequest) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeviceRequest) predicate.DeviceRequest {
	return predicate.DeviceRequest(func(s *sql.Selector) {
		p(s.Not())
	})
}
