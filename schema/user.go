package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		CommonMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("Username").Unique().NotEmpty().StructTag(`json:"username"`),
		field.String("Mail").Unique().NotEmpty().StructTag(`json:"mail"`),
		field.String("Password").NotEmpty().StructTag(`json:"-"`),
		field.Bool("MailValid").StorageKey("mail_valid").Default(false).StructTag(`json:"mailValid"`),
		field.Bool("TutorialCompleted").StorageKey("tutorial_completed").Default(false).StructTag(`json:"tutorialCompleted"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("activities", Activity.Type),
		edge.To("codes", VerificationCode.Type),
	}
}
