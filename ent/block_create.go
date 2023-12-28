// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"workblok/ent/block"
	"workblok/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlockCreate is the builder for creating a Block entity.
type BlockCreate struct {
	config
	mutation *BlockMutation
	hooks    []Hook
}

// SetCreationDate sets the "creationDate" field.
func (bc *BlockCreate) SetCreationDate(t time.Time) *BlockCreate {
	bc.mutation.SetCreationDate(t)
	return bc
}

// SetNillableCreationDate sets the "creationDate" field if the given value is not nil.
func (bc *BlockCreate) SetNillableCreationDate(t *time.Time) *BlockCreate {
	if t != nil {
		bc.SetCreationDate(*t)
	}
	return bc
}

// SetFinishDate sets the "finishDate" field.
func (bc *BlockCreate) SetFinishDate(t time.Time) *BlockCreate {
	bc.mutation.SetFinishDate(t)
	return bc
}

// SetNillableFinishDate sets the "finishDate" field if the given value is not nil.
func (bc *BlockCreate) SetNillableFinishDate(t *time.Time) *BlockCreate {
	if t != nil {
		bc.SetFinishDate(*t)
	}
	return bc
}

// SetTargetMinutes sets the "targetMinutes" field.
func (bc *BlockCreate) SetTargetMinutes(i int) *BlockCreate {
	bc.mutation.SetTargetMinutes(i)
	return bc
}

// SetNillableTargetMinutes sets the "targetMinutes" field if the given value is not nil.
func (bc *BlockCreate) SetNillableTargetMinutes(i *int) *BlockCreate {
	if i != nil {
		bc.SetTargetMinutes(*i)
	}
	return bc
}

// SetDistractionMinutes sets the "distractionMinutes" field.
func (bc *BlockCreate) SetDistractionMinutes(i int) *BlockCreate {
	bc.mutation.SetDistractionMinutes(i)
	return bc
}

// SetNillableDistractionMinutes sets the "distractionMinutes" field if the given value is not nil.
func (bc *BlockCreate) SetNillableDistractionMinutes(i *int) *BlockCreate {
	if i != nil {
		bc.SetDistractionMinutes(*i)
	}
	return bc
}

// SetTag sets the "tag" field.
func (bc *BlockCreate) SetTag(s string) *BlockCreate {
	bc.mutation.SetTag(s)
	return bc
}

// SetID sets the "id" field.
func (bc *BlockCreate) SetID(u uuid.UUID) *BlockCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BlockCreate) SetNillableID(u *uuid.UUID) *BlockCreate {
	if u != nil {
		bc.SetID(*u)
	}
	return bc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (bc *BlockCreate) SetUserID(id uuid.UUID) *BlockCreate {
	bc.mutation.SetUserID(id)
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BlockCreate) SetUser(u *User) *BlockCreate {
	return bc.SetUserID(u.ID)
}

// Mutation returns the BlockMutation object of the builder.
func (bc *BlockCreate) Mutation() *BlockMutation {
	return bc.mutation
}

// Save creates the Block in the database.
func (bc *BlockCreate) Save(ctx context.Context) (*Block, error) {
	bc.defaults()
	return withHooks[*Block, BlockMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockCreate) SaveX(ctx context.Context) *Block {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlockCreate) defaults() {
	if _, ok := bc.mutation.CreationDate(); !ok {
		v := block.DefaultCreationDate()
		bc.mutation.SetCreationDate(v)
	}
	if _, ok := bc.mutation.TargetMinutes(); !ok {
		v := block.DefaultTargetMinutes
		bc.mutation.SetTargetMinutes(v)
	}
	if _, ok := bc.mutation.DistractionMinutes(); !ok {
		v := block.DefaultDistractionMinutes
		bc.mutation.SetDistractionMinutes(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := block.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockCreate) check() error {
	if _, ok := bc.mutation.CreationDate(); !ok {
		return &ValidationError{Name: "creationDate", err: errors.New(`ent: missing required field "Block.creationDate"`)}
	}
	if _, ok := bc.mutation.TargetMinutes(); !ok {
		return &ValidationError{Name: "targetMinutes", err: errors.New(`ent: missing required field "Block.targetMinutes"`)}
	}
	if _, ok := bc.mutation.DistractionMinutes(); !ok {
		return &ValidationError{Name: "distractionMinutes", err: errors.New(`ent: missing required field "Block.distractionMinutes"`)}
	}
	if _, ok := bc.mutation.Tag(); !ok {
		return &ValidationError{Name: "tag", err: errors.New(`ent: missing required field "Block.tag"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Block.user"`)}
	}
	return nil
}

func (bc *BlockCreate) sqlSave(ctx context.Context) (*Block, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlockCreate) createSpec() (*Block, *sqlgraph.CreateSpec) {
	var (
		_node = &Block{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(block.Table, sqlgraph.NewFieldSpec(block.FieldID, field.TypeUUID))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.CreationDate(); ok {
		_spec.SetField(block.FieldCreationDate, field.TypeTime, value)
		_node.CreationDate = value
	}
	if value, ok := bc.mutation.FinishDate(); ok {
		_spec.SetField(block.FieldFinishDate, field.TypeTime, value)
		_node.FinishDate = &value
	}
	if value, ok := bc.mutation.TargetMinutes(); ok {
		_spec.SetField(block.FieldTargetMinutes, field.TypeInt, value)
		_node.TargetMinutes = value
	}
	if value, ok := bc.mutation.DistractionMinutes(); ok {
		_spec.SetField(block.FieldDistractionMinutes, field.TypeInt, value)
		_node.DistractionMinutes = value
	}
	if value, ok := bc.mutation.Tag(); ok {
		_spec.SetField(block.FieldTag, field.TypeString, value)
		_node.Tag = &value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   block.UserTable,
			Columns: []string{block.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_blocks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlockCreateBulk is the builder for creating many Block entities in bulk.
type BlockCreateBulk struct {
	config
	builders []*BlockCreate
}

// Save creates the Block entities in the database.
func (bcb *BlockCreateBulk) Save(ctx context.Context) ([]*Block, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Block, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockCreateBulk) SaveX(ctx context.Context) []*Block {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
