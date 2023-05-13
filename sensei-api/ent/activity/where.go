// Code generated by ent, DO NOT EDIT.

package activity

import (
	"sensei/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Activity {
	return predicate.Activity(sql.FieldLTE(FieldID, id))
}

// CreationDate applies equality check predicate on the "creationDate" field. It's identical to CreationDateEQ.
func CreationDate(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldCreationDate, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldName, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldIcon, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldSize, v))
}

// CreationDateEQ applies the EQ predicate on the "creationDate" field.
func CreationDateEQ(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldCreationDate, v))
}

// CreationDateNEQ applies the NEQ predicate on the "creationDate" field.
func CreationDateNEQ(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldNEQ(FieldCreationDate, v))
}

// CreationDateIn applies the In predicate on the "creationDate" field.
func CreationDateIn(vs ...time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldIn(FieldCreationDate, vs...))
}

// CreationDateNotIn applies the NotIn predicate on the "creationDate" field.
func CreationDateNotIn(vs ...time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldNotIn(FieldCreationDate, vs...))
}

// CreationDateGT applies the GT predicate on the "creationDate" field.
func CreationDateGT(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldGT(FieldCreationDate, v))
}

// CreationDateGTE applies the GTE predicate on the "creationDate" field.
func CreationDateGTE(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldGTE(FieldCreationDate, v))
}

// CreationDateLT applies the LT predicate on the "creationDate" field.
func CreationDateLT(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldLT(FieldCreationDate, v))
}

// CreationDateLTE applies the LTE predicate on the "creationDate" field.
func CreationDateLTE(v time.Time) predicate.Activity {
	return predicate.Activity(sql.FieldLTE(FieldCreationDate, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Activity {
	return predicate.Activity(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Activity {
	return predicate.Activity(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Activity {
	return predicate.Activity(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Activity {
	return predicate.Activity(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Activity {
	return predicate.Activity(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Activity {
	return predicate.Activity(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Activity {
	return predicate.Activity(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Activity {
	return predicate.Activity(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Activity {
	return predicate.Activity(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Activity {
	return predicate.Activity(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Activity {
	return predicate.Activity(sql.FieldContainsFold(FieldName, v))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Activity {
	return predicate.Activity(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Activity {
	return predicate.Activity(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Activity {
	return predicate.Activity(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Activity {
	return predicate.Activity(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Activity {
	return predicate.Activity(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Activity {
	return predicate.Activity(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Activity {
	return predicate.Activity(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Activity {
	return predicate.Activity(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Activity {
	return predicate.Activity(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Activity {
	return predicate.Activity(sql.FieldHasSuffix(FieldIcon, v))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Activity {
	return predicate.Activity(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Activity {
	return predicate.Activity(sql.FieldContainsFold(FieldIcon, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.Activity {
	return predicate.Activity(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.Activity {
	return predicate.Activity(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.Activity {
	return predicate.Activity(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.Activity {
	return predicate.Activity(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.Activity {
	return predicate.Activity(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.Activity {
	return predicate.Activity(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.Activity {
	return predicate.Activity(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.Activity {
	return predicate.Activity(sql.FieldLTE(FieldSize, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TasksTable, TasksColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
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
func Not(p predicate.Activity) predicate.Activity {
	return predicate.Activity(func(s *sql.Selector) {
		p(s.Not())
	})
}
