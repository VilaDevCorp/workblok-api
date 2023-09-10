// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sensei/ent/activity"
	"sensei/ent/predicate"
	"sensei/ent/task"
	"sensei/ent/templatetask"
	"sensei/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ActivityUpdate is the builder for updating Activity entities.
type ActivityUpdate struct {
	config
	hooks    []Hook
	mutation *ActivityMutation
}

// Where appends a list predicates to the ActivityUpdate builder.
func (au *ActivityUpdate) Where(ps ...predicate.Activity) *ActivityUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ActivityUpdate) SetName(s string) *ActivityUpdate {
	au.mutation.SetName(s)
	return au
}

// SetDescription sets the "description" field.
func (au *ActivityUpdate) SetDescription(s string) *ActivityUpdate {
	au.mutation.SetDescription(s)
	return au
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableDescription(s *string) *ActivityUpdate {
	if s != nil {
		au.SetDescription(*s)
	}
	return au
}

// ClearDescription clears the value of the "description" field.
func (au *ActivityUpdate) ClearDescription() *ActivityUpdate {
	au.mutation.ClearDescription()
	return au
}

// SetIcon sets the "icon" field.
func (au *ActivityUpdate) SetIcon(s string) *ActivityUpdate {
	au.mutation.SetIcon(s)
	return au
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableIcon(s *string) *ActivityUpdate {
	if s != nil {
		au.SetIcon(*s)
	}
	return au
}

// ClearIcon clears the value of the "icon" field.
func (au *ActivityUpdate) ClearIcon() *ActivityUpdate {
	au.mutation.ClearIcon()
	return au
}

// SetSize sets the "size" field.
func (au *ActivityUpdate) SetSize(i int) *ActivityUpdate {
	au.mutation.ResetSize()
	au.mutation.SetSize(i)
	return au
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (au *ActivityUpdate) SetNillableSize(i *int) *ActivityUpdate {
	if i != nil {
		au.SetSize(*i)
	}
	return au
}

// AddSize adds i to the "size" field.
func (au *ActivityUpdate) AddSize(i int) *ActivityUpdate {
	au.mutation.AddSize(i)
	return au
}

// SetUserID sets the "user" edge to the User entity by ID.
func (au *ActivityUpdate) SetUserID(id uuid.UUID) *ActivityUpdate {
	au.mutation.SetUserID(id)
	return au
}

// SetUser sets the "user" edge to the User entity.
func (au *ActivityUpdate) SetUser(u *User) *ActivityUpdate {
	return au.SetUserID(u.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (au *ActivityUpdate) AddTaskIDs(ids ...uuid.UUID) *ActivityUpdate {
	au.mutation.AddTaskIDs(ids...)
	return au
}

// AddTasks adds the "tasks" edges to the Task entity.
func (au *ActivityUpdate) AddTasks(t ...*Task) *ActivityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTaskIDs(ids...)
}

// AddTemplateTaskIDs adds the "templateTasks" edge to the TemplateTask entity by IDs.
func (au *ActivityUpdate) AddTemplateTaskIDs(ids ...uuid.UUID) *ActivityUpdate {
	au.mutation.AddTemplateTaskIDs(ids...)
	return au
}

// AddTemplateTasks adds the "templateTasks" edges to the TemplateTask entity.
func (au *ActivityUpdate) AddTemplateTasks(t ...*TemplateTask) *ActivityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddTemplateTaskIDs(ids...)
}

// Mutation returns the ActivityMutation object of the builder.
func (au *ActivityUpdate) Mutation() *ActivityMutation {
	return au.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (au *ActivityUpdate) ClearUser() *ActivityUpdate {
	au.mutation.ClearUser()
	return au
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (au *ActivityUpdate) ClearTasks() *ActivityUpdate {
	au.mutation.ClearTasks()
	return au
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (au *ActivityUpdate) RemoveTaskIDs(ids ...uuid.UUID) *ActivityUpdate {
	au.mutation.RemoveTaskIDs(ids...)
	return au
}

// RemoveTasks removes "tasks" edges to Task entities.
func (au *ActivityUpdate) RemoveTasks(t ...*Task) *ActivityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTaskIDs(ids...)
}

// ClearTemplateTasks clears all "templateTasks" edges to the TemplateTask entity.
func (au *ActivityUpdate) ClearTemplateTasks() *ActivityUpdate {
	au.mutation.ClearTemplateTasks()
	return au
}

// RemoveTemplateTaskIDs removes the "templateTasks" edge to TemplateTask entities by IDs.
func (au *ActivityUpdate) RemoveTemplateTaskIDs(ids ...uuid.UUID) *ActivityUpdate {
	au.mutation.RemoveTemplateTaskIDs(ids...)
	return au
}

// RemoveTemplateTasks removes "templateTasks" edges to TemplateTask entities.
func (au *ActivityUpdate) RemoveTemplateTasks(t ...*TemplateTask) *ActivityUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveTemplateTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ActivityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ActivityMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ActivityUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ActivityUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ActivityUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ActivityUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := activity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Activity.name": %w`, err)}
		}
	}
	if _, ok := au.mutation.UserID(); au.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Activity.user"`)
	}
	return nil
}

func (au *ActivityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(activity.Table, activity.Columns, sqlgraph.NewFieldSpec(activity.FieldID, field.TypeUUID))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(activity.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Description(); ok {
		_spec.SetField(activity.FieldDescription, field.TypeString, value)
	}
	if au.mutation.DescriptionCleared() {
		_spec.ClearField(activity.FieldDescription, field.TypeString)
	}
	if value, ok := au.mutation.Icon(); ok {
		_spec.SetField(activity.FieldIcon, field.TypeString, value)
	}
	if au.mutation.IconCleared() {
		_spec.ClearField(activity.FieldIcon, field.TypeString)
	}
	if value, ok := au.mutation.Size(); ok {
		_spec.SetField(activity.FieldSize, field.TypeInt, value)
	}
	if value, ok := au.mutation.AddedSize(); ok {
		_spec.AddField(activity.FieldSize, field.TypeInt, value)
	}
	if au.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTasksIDs(); len(nodes) > 0 && !au.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.TemplateTasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedTemplateTasksIDs(); len(nodes) > 0 && !au.mutation.TemplateTasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.TemplateTasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ActivityUpdateOne is the builder for updating a single Activity entity.
type ActivityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ActivityMutation
}

// SetName sets the "name" field.
func (auo *ActivityUpdateOne) SetName(s string) *ActivityUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetDescription sets the "description" field.
func (auo *ActivityUpdateOne) SetDescription(s string) *ActivityUpdateOne {
	auo.mutation.SetDescription(s)
	return auo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableDescription(s *string) *ActivityUpdateOne {
	if s != nil {
		auo.SetDescription(*s)
	}
	return auo
}

// ClearDescription clears the value of the "description" field.
func (auo *ActivityUpdateOne) ClearDescription() *ActivityUpdateOne {
	auo.mutation.ClearDescription()
	return auo
}

// SetIcon sets the "icon" field.
func (auo *ActivityUpdateOne) SetIcon(s string) *ActivityUpdateOne {
	auo.mutation.SetIcon(s)
	return auo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableIcon(s *string) *ActivityUpdateOne {
	if s != nil {
		auo.SetIcon(*s)
	}
	return auo
}

// ClearIcon clears the value of the "icon" field.
func (auo *ActivityUpdateOne) ClearIcon() *ActivityUpdateOne {
	auo.mutation.ClearIcon()
	return auo
}

// SetSize sets the "size" field.
func (auo *ActivityUpdateOne) SetSize(i int) *ActivityUpdateOne {
	auo.mutation.ResetSize()
	auo.mutation.SetSize(i)
	return auo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (auo *ActivityUpdateOne) SetNillableSize(i *int) *ActivityUpdateOne {
	if i != nil {
		auo.SetSize(*i)
	}
	return auo
}

// AddSize adds i to the "size" field.
func (auo *ActivityUpdateOne) AddSize(i int) *ActivityUpdateOne {
	auo.mutation.AddSize(i)
	return auo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (auo *ActivityUpdateOne) SetUserID(id uuid.UUID) *ActivityUpdateOne {
	auo.mutation.SetUserID(id)
	return auo
}

// SetUser sets the "user" edge to the User entity.
func (auo *ActivityUpdateOne) SetUser(u *User) *ActivityUpdateOne {
	return auo.SetUserID(u.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (auo *ActivityUpdateOne) AddTaskIDs(ids ...uuid.UUID) *ActivityUpdateOne {
	auo.mutation.AddTaskIDs(ids...)
	return auo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (auo *ActivityUpdateOne) AddTasks(t ...*Task) *ActivityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTaskIDs(ids...)
}

// AddTemplateTaskIDs adds the "templateTasks" edge to the TemplateTask entity by IDs.
func (auo *ActivityUpdateOne) AddTemplateTaskIDs(ids ...uuid.UUID) *ActivityUpdateOne {
	auo.mutation.AddTemplateTaskIDs(ids...)
	return auo
}

// AddTemplateTasks adds the "templateTasks" edges to the TemplateTask entity.
func (auo *ActivityUpdateOne) AddTemplateTasks(t ...*TemplateTask) *ActivityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddTemplateTaskIDs(ids...)
}

// Mutation returns the ActivityMutation object of the builder.
func (auo *ActivityUpdateOne) Mutation() *ActivityMutation {
	return auo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (auo *ActivityUpdateOne) ClearUser() *ActivityUpdateOne {
	auo.mutation.ClearUser()
	return auo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (auo *ActivityUpdateOne) ClearTasks() *ActivityUpdateOne {
	auo.mutation.ClearTasks()
	return auo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (auo *ActivityUpdateOne) RemoveTaskIDs(ids ...uuid.UUID) *ActivityUpdateOne {
	auo.mutation.RemoveTaskIDs(ids...)
	return auo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (auo *ActivityUpdateOne) RemoveTasks(t ...*Task) *ActivityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTaskIDs(ids...)
}

// ClearTemplateTasks clears all "templateTasks" edges to the TemplateTask entity.
func (auo *ActivityUpdateOne) ClearTemplateTasks() *ActivityUpdateOne {
	auo.mutation.ClearTemplateTasks()
	return auo
}

// RemoveTemplateTaskIDs removes the "templateTasks" edge to TemplateTask entities by IDs.
func (auo *ActivityUpdateOne) RemoveTemplateTaskIDs(ids ...uuid.UUID) *ActivityUpdateOne {
	auo.mutation.RemoveTemplateTaskIDs(ids...)
	return auo
}

// RemoveTemplateTasks removes "templateTasks" edges to TemplateTask entities.
func (auo *ActivityUpdateOne) RemoveTemplateTasks(t ...*TemplateTask) *ActivityUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveTemplateTaskIDs(ids...)
}

// Where appends a list predicates to the ActivityUpdate builder.
func (auo *ActivityUpdateOne) Where(ps ...predicate.Activity) *ActivityUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ActivityUpdateOne) Select(field string, fields ...string) *ActivityUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Activity entity.
func (auo *ActivityUpdateOne) Save(ctx context.Context) (*Activity, error) {
	return withHooks[*Activity, ActivityMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ActivityUpdateOne) SaveX(ctx context.Context) *Activity {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ActivityUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ActivityUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ActivityUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := activity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Activity.name": %w`, err)}
		}
	}
	if _, ok := auo.mutation.UserID(); auo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Activity.user"`)
	}
	return nil
}

func (auo *ActivityUpdateOne) sqlSave(ctx context.Context) (_node *Activity, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(activity.Table, activity.Columns, sqlgraph.NewFieldSpec(activity.FieldID, field.TypeUUID))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Activity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, activity.FieldID)
		for _, f := range fields {
			if !activity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != activity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(activity.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Description(); ok {
		_spec.SetField(activity.FieldDescription, field.TypeString, value)
	}
	if auo.mutation.DescriptionCleared() {
		_spec.ClearField(activity.FieldDescription, field.TypeString)
	}
	if value, ok := auo.mutation.Icon(); ok {
		_spec.SetField(activity.FieldIcon, field.TypeString, value)
	}
	if auo.mutation.IconCleared() {
		_spec.ClearField(activity.FieldIcon, field.TypeString)
	}
	if value, ok := auo.mutation.Size(); ok {
		_spec.SetField(activity.FieldSize, field.TypeInt, value)
	}
	if value, ok := auo.mutation.AddedSize(); ok {
		_spec.AddField(activity.FieldSize, field.TypeInt, value)
	}
	if auo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !auo.mutation.TasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.TemplateTasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedTemplateTasksIDs(); len(nodes) > 0 && !auo.mutation.TemplateTasksCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.TemplateTasksIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Activity{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{activity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
