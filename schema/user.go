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

const (
	Minutes = "minutes"
	Hours   = "hours"
)

type Config struct {
	DailyTarget  *int      `json:"dailyTarget"` // in minutes
	DarkMode     *bool     `json:"darkMode"`
	ExceededTime *bool     `json:"exceededTime"`
	TimeLimit    *int      `json:"timeLimit"` // in minutes
	Tags         *[]string `json:"tags"`
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
		field.String("Email").Unique().NotEmpty().StructTag(`json:"email"`),
		field.String("Password").NotEmpty().StructTag(`json:"-"`),
		field.Bool("EmailValid").
			StorageKey("email_valid").
			Default(false).
			StructTag(`json:"emailValid"`),
		field.JSON("Config", &Config{}).StructTag(`json:"config"`),
		field.Bool("TutorialCompleted").
			StorageKey("tutorial_completed").
			Default(false).
			StructTag(`json:"tutorialCompleted"`),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blocks", Block.Type),
		edge.To("codes", VerificationCode.Type),
		edge.To("sessions", Session.Type),
	}
}
