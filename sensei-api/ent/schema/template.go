package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Template holds the schema definition for the Template entity.
type Template struct {
	ent.Schema
}

func (Template) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Template.
func (Template) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Template.
func (Template) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("templates").Required().Unique(),
		edge.To("templateTasks", TemplateTask.Type),
	}
}
