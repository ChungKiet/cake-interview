package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TimeMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Nillable(),
		field.String("account"),
		field.String("full_name").NotEmpty(),
		field.String("phone_numer").Nillable(),
		field.String("email").Nillable(),
		field.String("user_name").Nillable(),
		field.String("password").NotEmpty(),
		field.Time("birthday"),
		field.Time("latest_login"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
