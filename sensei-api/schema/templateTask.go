package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TemplateTask holds the schema definition for the TemplateTask entity.
type TemplateTask struct {
	ent.Schema
}

func (TemplateTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the TemplateTask.
func (TemplateTask) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weekDay"),
	}
}

// Edges of the TemplateTask.
func (TemplateTask) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("activity", Activity.Type).Ref("templateTasks").Required().Unique(),
		edge.From("template", Template.Type).Ref("templateTasks").Required().Unique(),
	}
}
