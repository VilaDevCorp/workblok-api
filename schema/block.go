package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Block holds the schema definition for the Block entity.
type Block struct {
	ent.Schema
}

func (Block) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the Block.
func (Block) Fields() []ent.Field {
	return []ent.Field{
		field.Time("finishDate").Optional().Nillable(),
		field.Int("targetMinutes").Default(5),
		field.Int("distractionMinutes").Default(0),
		field.String("tag").Nillable(),
	}
}

// Edges of the Block.
func (Block) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("blocks").Required().Unique(),
	}
}
