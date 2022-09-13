// Code generated by entc, DO NOT EDIT.

package connector

import (
	"entgo.io/ent/dialect/sql"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ResourceVersion applies equality check predicate on the "resource_version" field. It's identical to ResourceVersionEQ.
func ResourceVersion(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceVersion), v))
	})
}

// Config applies equality check predicate on the "config" field. It's identical to ConfigEQ.
func Config(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ResourceVersionEQ applies the EQ predicate on the "resource_version" field.
func ResourceVersionEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionNEQ applies the NEQ predicate on the "resource_version" field.
func ResourceVersionNEQ(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionIn applies the In predicate on the "resource_version" field.
func ResourceVersionIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldResourceVersion), v...))
	})
}

// ResourceVersionNotIn applies the NotIn predicate on the "resource_version" field.
func ResourceVersionNotIn(vs ...string) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldResourceVersion), v...))
	})
}

// ResourceVersionGT applies the GT predicate on the "resource_version" field.
func ResourceVersionGT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionGTE applies the GTE predicate on the "resource_version" field.
func ResourceVersionGTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionLT applies the LT predicate on the "resource_version" field.
func ResourceVersionLT(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionLTE applies the LTE predicate on the "resource_version" field.
func ResourceVersionLTE(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionContains applies the Contains predicate on the "resource_version" field.
func ResourceVersionContains(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionHasPrefix applies the HasPrefix predicate on the "resource_version" field.
func ResourceVersionHasPrefix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionHasSuffix applies the HasSuffix predicate on the "resource_version" field.
func ResourceVersionHasSuffix(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionEqualFold applies the EqualFold predicate on the "resource_version" field.
func ResourceVersionEqualFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResourceVersion), v))
	})
}

// ResourceVersionContainsFold applies the ContainsFold predicate on the "resource_version" field.
func ResourceVersionContainsFold(v string) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResourceVersion), v))
	})
}

// ConfigEQ applies the EQ predicate on the "config" field.
func ConfigEQ(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConfig), v))
	})
}

// ConfigNEQ applies the NEQ predicate on the "config" field.
func ConfigNEQ(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConfig), v))
	})
}

// ConfigIn applies the In predicate on the "config" field.
func ConfigIn(vs ...[]byte) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConfig), v...))
	})
}

// ConfigNotIn applies the NotIn predicate on the "config" field.
func ConfigNotIn(vs ...[]byte) predicate.Connector {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Connector(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConfig), v...))
	})
}

// ConfigGT applies the GT predicate on the "config" field.
func ConfigGT(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConfig), v))
	})
}

// ConfigGTE applies the GTE predicate on the "config" field.
func ConfigGTE(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConfig), v))
	})
}

// ConfigLT applies the LT predicate on the "config" field.
func ConfigLT(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConfig), v))
	})
}

// ConfigLTE applies the LTE predicate on the "config" field.
func ConfigLTE(v []byte) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConfig), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
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
func Not(p predicate.Connector) predicate.Connector {
	return predicate.Connector(func(s *sql.Selector) {
		p(s.Not())
	})
}
