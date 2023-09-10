// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sensei/ent/activity"
	"sensei/ent/task"
	"sensei/ent/templatetask"
	"sensei/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ActivityCreate is the builder for creating a Activity entity.
type ActivityCreate struct {
	config
	mutation *ActivityMutation
	hooks    []Hook
}

// SetCreationDate sets the "creationDate" field.
func (ac *ActivityCreate) SetCreationDate(t time.Time) *ActivityCreate {
	ac.mutation.SetCreationDate(t)
	return ac
}

// SetNillableCreationDate sets the "creationDate" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableCreationDate(t *time.Time) *ActivityCreate {
	if t != nil {
		ac.SetCreationDate(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *ActivityCreate) SetName(s string) *ActivityCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *ActivityCreate) SetDescription(s string) *ActivityCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableDescription(s *string) *ActivityCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetIcon sets the "icon" field.
func (ac *ActivityCreate) SetIcon(s string) *ActivityCreate {
	ac.mutation.SetIcon(s)
	return ac
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableIcon(s *string) *ActivityCreate {
	if s != nil {
		ac.SetIcon(*s)
	}
	return ac
}

// SetSize sets the "size" field.
func (ac *ActivityCreate) SetSize(i int) *ActivityCreate {
	ac.mutation.SetSize(i)
	return ac
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableSize(i *int) *ActivityCreate {
	if i != nil {
		ac.SetSize(*i)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *ActivityCreate) SetID(u uuid.UUID) *ActivityCreate {
	ac.mutation.SetID(u)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *ActivityCreate) SetNillableID(u *uuid.UUID) *ActivityCreate {
	if u != nil {
		ac.SetID(*u)
	}
	return ac
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ac *ActivityCreate) SetUserID(id uuid.UUID) *ActivityCreate {
	ac.mutation.SetUserID(id)
	return ac
}

// SetUser sets the "user" edge to the User entity.
func (ac *ActivityCreate) SetUser(u *User) *ActivityCreate {
	return ac.SetUserID(u.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ac *ActivityCreate) AddTaskIDs(ids ...uuid.UUID) *ActivityCreate {
	ac.mutation.AddTaskIDs(ids...)
	return ac
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ac *ActivityCreate) AddTasks(t ...*Task) *ActivityCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTaskIDs(ids...)
}

// AddTemplateTaskIDs adds the "templateTasks" edge to the TemplateTask entity by IDs.
func (ac *ActivityCreate) AddTemplateTaskIDs(ids ...uuid.UUID) *ActivityCreate {
	ac.mutation.AddTemplateTaskIDs(ids...)
	return ac
}

// AddTemplateTasks adds the "templateTasks" edges to the TemplateTask entity.
func (ac *ActivityCreate) AddTemplateTasks(t ...*TemplateTask) *ActivityCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ac.AddTemplateTaskIDs(ids...)
}

// Mutation returns the ActivityMutation object of the builder.
func (ac *ActivityCreate) Mutation() *ActivityMutation {
	return ac.mutation
}

// Save creates the Activity in the database.
func (ac *ActivityCreate) Save(ctx context.Context) (*Activity, error) {
	ac.defaults()
	return withHooks[*Activity, ActivityMutation](ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *ActivityCreate) SaveX(ctx context.Context) *Activity {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *ActivityCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *ActivityCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *ActivityCreate) defaults() {
	if _, ok := ac.mutation.CreationDate(); !ok {
		v := activity.DefaultCreationDate()
		ac.mutation.SetCreationDate(v)
	}
	if _, ok := ac.mutation.Size(); !ok {
		v := activity.DefaultSize
		ac.mutation.SetSize(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := activity.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *ActivityCreate) check() error {
	if _, ok := ac.mutation.CreationDate(); !ok {
		return &ValidationError{Name: "creationDate", err: errors.New(`ent: missing required field "Activity.creationDate"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Activity.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := activity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Activity.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Activity.size"`)}
	}
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Activity.user"`)}
	}
	return nil
}

func (ac *ActivityCreate) sqlSave(ctx context.Context) (*Activity, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *ActivityCreate) createSpec() (*Activity, *sqlgraph.CreateSpec) {
	var (
		_node = &Activity{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(activity.Table, sqlgraph.NewFieldSpec(activity.FieldID, field.TypeUUID))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ac.mutation.CreationDate(); ok {
		_spec.SetField(activity.FieldCreationDate, field.TypeTime, value)
		_node.CreationDate = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(activity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(activity.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.Icon(); ok {
		_spec.SetField(activity.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := ac.mutation.Size(); ok {
		_spec.SetField(activity.FieldSize, field.TypeInt, value)
		_node.Size = value
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   activity.UserTable,
			Columns: []string{activity.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_activities = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activity.TasksTable,
			Columns: []string{activity.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TemplateTasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   activity.TemplateTasksTable,
			Columns: []string{activity.TemplateTasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(templatetask.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ActivityCreateBulk is the builder for creating many Activity entities in bulk.
type ActivityCreateBulk struct {
	config
	builders []*ActivityCreate
}

// Save creates the Activity entities in the database.
func (acb *ActivityCreateBulk) Save(ctx context.Context) ([]*Activity, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Activity, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ActivityMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *ActivityCreateBulk) SaveX(ctx context.Context) []*Activity {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *ActivityCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *ActivityCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
