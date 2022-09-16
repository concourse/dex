// Code generated by ent, DO NOT EDIT.

package devicetoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DeviceCode applies equality check predicate on the "device_code" field. It's identical to DeviceCodeEQ.
func DeviceCode(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceCode), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// LastRequest applies equality check predicate on the "last_request" field. It's identical to LastRequestEQ.
func LastRequest(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastRequest), v))
	})
}

// PollInterval applies equality check predicate on the "poll_interval" field. It's identical to PollIntervalEQ.
func PollInterval(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPollInterval), v))
	})
}

// CodeChallenge applies equality check predicate on the "code_challenge" field. It's identical to CodeChallengeEQ.
func CodeChallenge(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeMethod applies equality check predicate on the "code_challenge_method" field. It's identical to CodeChallengeMethodEQ.
func CodeChallengeMethod(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodeChallengeMethod), v))
	})
}

// DeviceCodeEQ applies the EQ predicate on the "device_code" field.
func DeviceCodeEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeNEQ applies the NEQ predicate on the "device_code" field.
func DeviceCodeNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeIn applies the In predicate on the "device_code" field.
func DeviceCodeIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeviceCode), v...))
	})
}

// DeviceCodeNotIn applies the NotIn predicate on the "device_code" field.
func DeviceCodeNotIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeviceCode), v...))
	})
}

// DeviceCodeGT applies the GT predicate on the "device_code" field.
func DeviceCodeGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeGTE applies the GTE predicate on the "device_code" field.
func DeviceCodeGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeLT applies the LT predicate on the "device_code" field.
func DeviceCodeLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeLTE applies the LTE predicate on the "device_code" field.
func DeviceCodeLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeContains applies the Contains predicate on the "device_code" field.
func DeviceCodeContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeHasPrefix applies the HasPrefix predicate on the "device_code" field.
func DeviceCodeHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeHasSuffix applies the HasSuffix predicate on the "device_code" field.
func DeviceCodeHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeEqualFold applies the EqualFold predicate on the "device_code" field.
func DeviceCodeEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDeviceCode), v))
	})
}

// DeviceCodeContainsFold applies the ContainsFold predicate on the "device_code" field.
func DeviceCodeContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDeviceCode), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldToken), v))
	})
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldToken), v))
	})
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...[]byte) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldToken), v...))
	})
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...[]byte) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldToken), v...))
	})
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldToken), v))
	})
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldToken), v))
	})
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldToken), v))
	})
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v []byte) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldToken), v))
	})
}

// TokenIsNil applies the IsNil predicate on the "token" field.
func TokenIsNil() predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldToken)))
	})
}

// TokenNotNil applies the NotNil predicate on the "token" field.
func TokenNotNil() predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldToken)))
	})
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExpiry), v))
	})
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExpiry), v))
	})
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldExpiry), v...))
	})
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldExpiry), v...))
	})
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExpiry), v))
	})
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExpiry), v))
	})
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExpiry), v))
	})
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExpiry), v))
	})
}

// LastRequestEQ applies the EQ predicate on the "last_request" field.
func LastRequestEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLastRequest), v))
	})
}

// LastRequestNEQ applies the NEQ predicate on the "last_request" field.
func LastRequestNEQ(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLastRequest), v))
	})
}

// LastRequestIn applies the In predicate on the "last_request" field.
func LastRequestIn(vs ...time.Time) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLastRequest), v...))
	})
}

// LastRequestNotIn applies the NotIn predicate on the "last_request" field.
func LastRequestNotIn(vs ...time.Time) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLastRequest), v...))
	})
}

// LastRequestGT applies the GT predicate on the "last_request" field.
func LastRequestGT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLastRequest), v))
	})
}

// LastRequestGTE applies the GTE predicate on the "last_request" field.
func LastRequestGTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLastRequest), v))
	})
}

// LastRequestLT applies the LT predicate on the "last_request" field.
func LastRequestLT(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLastRequest), v))
	})
}

// LastRequestLTE applies the LTE predicate on the "last_request" field.
func LastRequestLTE(v time.Time) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLastRequest), v))
	})
}

// PollIntervalEQ applies the EQ predicate on the "poll_interval" field.
func PollIntervalEQ(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPollInterval), v))
	})
}

// PollIntervalNEQ applies the NEQ predicate on the "poll_interval" field.
func PollIntervalNEQ(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPollInterval), v))
	})
}

// PollIntervalIn applies the In predicate on the "poll_interval" field.
func PollIntervalIn(vs ...int) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPollInterval), v...))
	})
}

// PollIntervalNotIn applies the NotIn predicate on the "poll_interval" field.
func PollIntervalNotIn(vs ...int) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPollInterval), v...))
	})
}

// PollIntervalGT applies the GT predicate on the "poll_interval" field.
func PollIntervalGT(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPollInterval), v))
	})
}

// PollIntervalGTE applies the GTE predicate on the "poll_interval" field.
func PollIntervalGTE(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPollInterval), v))
	})
}

// PollIntervalLT applies the LT predicate on the "poll_interval" field.
func PollIntervalLT(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPollInterval), v))
	})
}

// PollIntervalLTE applies the LTE predicate on the "poll_interval" field.
func PollIntervalLTE(v int) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPollInterval), v))
	})
}

// CodeChallengeEQ applies the EQ predicate on the "code_challenge" field.
func CodeChallengeEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeNEQ applies the NEQ predicate on the "code_challenge" field.
func CodeChallengeNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeIn applies the In predicate on the "code_challenge" field.
func CodeChallengeIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCodeChallenge), v...))
	})
}

// CodeChallengeNotIn applies the NotIn predicate on the "code_challenge" field.
func CodeChallengeNotIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCodeChallenge), v...))
	})
}

// CodeChallengeGT applies the GT predicate on the "code_challenge" field.
func CodeChallengeGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeGTE applies the GTE predicate on the "code_challenge" field.
func CodeChallengeGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeLT applies the LT predicate on the "code_challenge" field.
func CodeChallengeLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeLTE applies the LTE predicate on the "code_challenge" field.
func CodeChallengeLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeContains applies the Contains predicate on the "code_challenge" field.
func CodeChallengeContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeHasPrefix applies the HasPrefix predicate on the "code_challenge" field.
func CodeChallengeHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeHasSuffix applies the HasSuffix predicate on the "code_challenge" field.
func CodeChallengeHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeEqualFold applies the EqualFold predicate on the "code_challenge" field.
func CodeChallengeEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeContainsFold applies the ContainsFold predicate on the "code_challenge" field.
func CodeChallengeContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCodeChallenge), v))
	})
}

// CodeChallengeMethodEQ applies the EQ predicate on the "code_challenge_method" field.
func CodeChallengeMethodEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodNEQ applies the NEQ predicate on the "code_challenge_method" field.
func CodeChallengeMethodNEQ(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodIn applies the In predicate on the "code_challenge_method" field.
func CodeChallengeMethodIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCodeChallengeMethod), v...))
	})
}

// CodeChallengeMethodNotIn applies the NotIn predicate on the "code_challenge_method" field.
func CodeChallengeMethodNotIn(vs ...string) predicate.DeviceToken {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCodeChallengeMethod), v...))
	})
}

// CodeChallengeMethodGT applies the GT predicate on the "code_challenge_method" field.
func CodeChallengeMethodGT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodGTE applies the GTE predicate on the "code_challenge_method" field.
func CodeChallengeMethodGTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodLT applies the LT predicate on the "code_challenge_method" field.
func CodeChallengeMethodLT(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodLTE applies the LTE predicate on the "code_challenge_method" field.
func CodeChallengeMethodLTE(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodContains applies the Contains predicate on the "code_challenge_method" field.
func CodeChallengeMethodContains(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodHasPrefix applies the HasPrefix predicate on the "code_challenge_method" field.
func CodeChallengeMethodHasPrefix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodHasSuffix applies the HasSuffix predicate on the "code_challenge_method" field.
func CodeChallengeMethodHasSuffix(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodEqualFold applies the EqualFold predicate on the "code_challenge_method" field.
func CodeChallengeMethodEqualFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCodeChallengeMethod), v))
	})
}

// CodeChallengeMethodContainsFold applies the ContainsFold predicate on the "code_challenge_method" field.
func CodeChallengeMethodContainsFold(v string) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCodeChallengeMethod), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
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
func Not(p predicate.DeviceToken) predicate.DeviceToken {
	return predicate.DeviceToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
