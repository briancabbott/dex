// Code generated by entc, DO NOT EDIT.

package keys

import (
	"time"

	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/dex/storage/ent/db/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// NextRotation applies equality check predicate on the "next_rotation" field. It's identical to NextRotationEQ.
func NextRotation(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextRotation), v))
	})
}

// NextRotationEQ applies the EQ predicate on the "next_rotation" field.
func NextRotationEQ(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNextRotation), v))
	})
}

// NextRotationNEQ applies the NEQ predicate on the "next_rotation" field.
func NextRotationNEQ(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNextRotation), v))
	})
}

// NextRotationIn applies the In predicate on the "next_rotation" field.
func NextRotationIn(vs ...time.Time) predicate.Keys {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Keys(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNextRotation), v...))
	})
}

// NextRotationNotIn applies the NotIn predicate on the "next_rotation" field.
func NextRotationNotIn(vs ...time.Time) predicate.Keys {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Keys(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNextRotation), v...))
	})
}

// NextRotationGT applies the GT predicate on the "next_rotation" field.
func NextRotationGT(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNextRotation), v))
	})
}

// NextRotationGTE applies the GTE predicate on the "next_rotation" field.
func NextRotationGTE(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNextRotation), v))
	})
}

// NextRotationLT applies the LT predicate on the "next_rotation" field.
func NextRotationLT(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNextRotation), v))
	})
}

// NextRotationLTE applies the LTE predicate on the "next_rotation" field.
func NextRotationLTE(v time.Time) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNextRotation), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Keys) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Keys) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
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
func Not(p predicate.Keys) predicate.Keys {
	return predicate.Keys(func(s *sql.Selector) {
		p(s.Not())
	})
}
