// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"workblok/ent/predicate"
	"workblok/ent/user"
	"workblok/ent/verificationcode"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// VerificationCodeUpdate is the builder for updating VerificationCode entities.
type VerificationCodeUpdate struct {
	config
	hooks    []Hook
	mutation *VerificationCodeMutation
}

// Where appends a list predicates to the VerificationCodeUpdate builder.
func (vcu *VerificationCodeUpdate) Where(ps ...predicate.VerificationCode) *VerificationCodeUpdate {
	vcu.mutation.Where(ps...)
	return vcu
}

// SetType sets the "type" field.
func (vcu *VerificationCodeUpdate) SetType(s string) *VerificationCodeUpdate {
	vcu.mutation.SetType(s)
	return vcu
}

// SetCode sets the "code" field.
func (vcu *VerificationCodeUpdate) SetCode(s string) *VerificationCodeUpdate {
	vcu.mutation.SetCode(s)
	return vcu
}

// SetExpireDate sets the "expireDate" field.
func (vcu *VerificationCodeUpdate) SetExpireDate(t time.Time) *VerificationCodeUpdate {
	vcu.mutation.SetExpireDate(t)
	return vcu
}

// SetValid sets the "valid" field.
func (vcu *VerificationCodeUpdate) SetValid(b bool) *VerificationCodeUpdate {
	vcu.mutation.SetValid(b)
	return vcu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vcu *VerificationCodeUpdate) SetUserID(id uuid.UUID) *VerificationCodeUpdate {
	vcu.mutation.SetUserID(id)
	return vcu
}

// SetUser sets the "user" edge to the User entity.
func (vcu *VerificationCodeUpdate) SetUser(u *User) *VerificationCodeUpdate {
	return vcu.SetUserID(u.ID)
}

// Mutation returns the VerificationCodeMutation object of the builder.
func (vcu *VerificationCodeUpdate) Mutation() *VerificationCodeMutation {
	return vcu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vcu *VerificationCodeUpdate) ClearUser() *VerificationCodeUpdate {
	vcu.mutation.ClearUser()
	return vcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vcu *VerificationCodeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, VerificationCodeMutation](ctx, vcu.sqlSave, vcu.mutation, vcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vcu *VerificationCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := vcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vcu *VerificationCodeUpdate) Exec(ctx context.Context) error {
	_, err := vcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcu *VerificationCodeUpdate) ExecX(ctx context.Context) {
	if err := vcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcu *VerificationCodeUpdate) check() error {
	if v, ok := vcu.mutation.GetType(); ok {
		if err := verificationcode.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "VerificationCode.type": %w`, err)}
		}
	}
	if v, ok := vcu.mutation.Code(); ok {
		if err := verificationcode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "VerificationCode.code": %w`, err)}
		}
	}
	if _, ok := vcu.mutation.UserID(); vcu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "VerificationCode.user"`)
	}
	return nil
}

func (vcu *VerificationCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := vcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(verificationcode.Table, verificationcode.Columns, sqlgraph.NewFieldSpec(verificationcode.FieldID, field.TypeUUID))
	if ps := vcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vcu.mutation.GetType(); ok {
		_spec.SetField(verificationcode.FieldType, field.TypeString, value)
	}
	if value, ok := vcu.mutation.Code(); ok {
		_spec.SetField(verificationcode.FieldCode, field.TypeString, value)
	}
	if value, ok := vcu.mutation.ExpireDate(); ok {
		_spec.SetField(verificationcode.FieldExpireDate, field.TypeTime, value)
	}
	if value, ok := vcu.mutation.Valid(); ok {
		_spec.SetField(verificationcode.FieldValid, field.TypeBool, value)
	}
	if vcu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   verificationcode.UserTable,
			Columns: []string{verificationcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vcu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   verificationcode.UserTable,
			Columns: []string{verificationcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verificationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vcu.mutation.done = true
	return n, nil
}

// VerificationCodeUpdateOne is the builder for updating a single VerificationCode entity.
type VerificationCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VerificationCodeMutation
}

// SetType sets the "type" field.
func (vcuo *VerificationCodeUpdateOne) SetType(s string) *VerificationCodeUpdateOne {
	vcuo.mutation.SetType(s)
	return vcuo
}

// SetCode sets the "code" field.
func (vcuo *VerificationCodeUpdateOne) SetCode(s string) *VerificationCodeUpdateOne {
	vcuo.mutation.SetCode(s)
	return vcuo
}

// SetExpireDate sets the "expireDate" field.
func (vcuo *VerificationCodeUpdateOne) SetExpireDate(t time.Time) *VerificationCodeUpdateOne {
	vcuo.mutation.SetExpireDate(t)
	return vcuo
}

// SetValid sets the "valid" field.
func (vcuo *VerificationCodeUpdateOne) SetValid(b bool) *VerificationCodeUpdateOne {
	vcuo.mutation.SetValid(b)
	return vcuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vcuo *VerificationCodeUpdateOne) SetUserID(id uuid.UUID) *VerificationCodeUpdateOne {
	vcuo.mutation.SetUserID(id)
	return vcuo
}

// SetUser sets the "user" edge to the User entity.
func (vcuo *VerificationCodeUpdateOne) SetUser(u *User) *VerificationCodeUpdateOne {
	return vcuo.SetUserID(u.ID)
}

// Mutation returns the VerificationCodeMutation object of the builder.
func (vcuo *VerificationCodeUpdateOne) Mutation() *VerificationCodeMutation {
	return vcuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vcuo *VerificationCodeUpdateOne) ClearUser() *VerificationCodeUpdateOne {
	vcuo.mutation.ClearUser()
	return vcuo
}

// Where appends a list predicates to the VerificationCodeUpdate builder.
func (vcuo *VerificationCodeUpdateOne) Where(ps ...predicate.VerificationCode) *VerificationCodeUpdateOne {
	vcuo.mutation.Where(ps...)
	return vcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vcuo *VerificationCodeUpdateOne) Select(field string, fields ...string) *VerificationCodeUpdateOne {
	vcuo.fields = append([]string{field}, fields...)
	return vcuo
}

// Save executes the query and returns the updated VerificationCode entity.
func (vcuo *VerificationCodeUpdateOne) Save(ctx context.Context) (*VerificationCode, error) {
	return withHooks[*VerificationCode, VerificationCodeMutation](ctx, vcuo.sqlSave, vcuo.mutation, vcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vcuo *VerificationCodeUpdateOne) SaveX(ctx context.Context) *VerificationCode {
	node, err := vcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vcuo *VerificationCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := vcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcuo *VerificationCodeUpdateOne) ExecX(ctx context.Context) {
	if err := vcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vcuo *VerificationCodeUpdateOne) check() error {
	if v, ok := vcuo.mutation.GetType(); ok {
		if err := verificationcode.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "VerificationCode.type": %w`, err)}
		}
	}
	if v, ok := vcuo.mutation.Code(); ok {
		if err := verificationcode.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "VerificationCode.code": %w`, err)}
		}
	}
	if _, ok := vcuo.mutation.UserID(); vcuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "VerificationCode.user"`)
	}
	return nil
}

func (vcuo *VerificationCodeUpdateOne) sqlSave(ctx context.Context) (_node *VerificationCode, err error) {
	if err := vcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(verificationcode.Table, verificationcode.Columns, sqlgraph.NewFieldSpec(verificationcode.FieldID, field.TypeUUID))
	id, ok := vcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VerificationCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, verificationcode.FieldID)
		for _, f := range fields {
			if !verificationcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != verificationcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vcuo.mutation.GetType(); ok {
		_spec.SetField(verificationcode.FieldType, field.TypeString, value)
	}
	if value, ok := vcuo.mutation.Code(); ok {
		_spec.SetField(verificationcode.FieldCode, field.TypeString, value)
	}
	if value, ok := vcuo.mutation.ExpireDate(); ok {
		_spec.SetField(verificationcode.FieldExpireDate, field.TypeTime, value)
	}
	if value, ok := vcuo.mutation.Valid(); ok {
		_spec.SetField(verificationcode.FieldValid, field.TypeBool, value)
	}
	if vcuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   verificationcode.UserTable,
			Columns: []string{verificationcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vcuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   verificationcode.UserTable,
			Columns: []string{verificationcode.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VerificationCode{config: vcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{verificationcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vcuo.mutation.done = true
	return _node, nil
}
