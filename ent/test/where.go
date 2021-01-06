// Code generated by entc, DO NOT EDIT.

package test

import (
	"align/ent/predicate"

	"github.com/facebook/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Field1 applies equality check predicate on the "field1" field. It's identical to Field1EQ.
func Field1(v bool) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField1), v))
	})
}

// Field2 applies equality check predicate on the "field2" field. It's identical to Field2EQ.
func Field2(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField2), v))
	})
}

// Field3 applies equality check predicate on the "field3" field. It's identical to Field3EQ.
func Field3(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField3), v))
	})
}

// Field1EQ applies the EQ predicate on the "field1" field.
func Field1EQ(v bool) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField1), v))
	})
}

// Field1NEQ applies the NEQ predicate on the "field1" field.
func Field1NEQ(v bool) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldField1), v))
	})
}

// Field2EQ applies the EQ predicate on the "field2" field.
func Field2EQ(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField2), v))
	})
}

// Field2NEQ applies the NEQ predicate on the "field2" field.
func Field2NEQ(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldField2), v))
	})
}

// Field2In applies the In predicate on the "field2" field.
func Field2In(vs ...float64) predicate.Test {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Test(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldField2), v...))
	})
}

// Field2NotIn applies the NotIn predicate on the "field2" field.
func Field2NotIn(vs ...float64) predicate.Test {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Test(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldField2), v...))
	})
}

// Field2GT applies the GT predicate on the "field2" field.
func Field2GT(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldField2), v))
	})
}

// Field2GTE applies the GTE predicate on the "field2" field.
func Field2GTE(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldField2), v))
	})
}

// Field2LT applies the LT predicate on the "field2" field.
func Field2LT(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldField2), v))
	})
}

// Field2LTE applies the LTE predicate on the "field2" field.
func Field2LTE(v float64) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldField2), v))
	})
}

// Field3EQ applies the EQ predicate on the "field3" field.
func Field3EQ(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldField3), v))
	})
}

// Field3NEQ applies the NEQ predicate on the "field3" field.
func Field3NEQ(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldField3), v))
	})
}

// Field3In applies the In predicate on the "field3" field.
func Field3In(vs ...int32) predicate.Test {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Test(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldField3), v...))
	})
}

// Field3NotIn applies the NotIn predicate on the "field3" field.
func Field3NotIn(vs ...int32) predicate.Test {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Test(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldField3), v...))
	})
}

// Field3GT applies the GT predicate on the "field3" field.
func Field3GT(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldField3), v))
	})
}

// Field3GTE applies the GTE predicate on the "field3" field.
func Field3GTE(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldField3), v))
	})
}

// Field3LT applies the LT predicate on the "field3" field.
func Field3LT(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldField3), v))
	})
}

// Field3LTE applies the LTE predicate on the "field3" field.
func Field3LTE(v int32) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldField3), v))
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Test) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Test) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
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
func Not(p predicate.Test) predicate.Test {
	return predicate.Test(func(s *sql.Selector) {
		p(s.Not())
	})
}
