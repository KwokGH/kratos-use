package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

var (
	BaseFields = []ent.Field{
		field.String("id").
			MaxLen(60).
			NotEmpty().
			Unique().
			Immutable().DefaultFunc(func() string {
			id, _ := uuid.NewUUID()
			return id.String()
		}),
		field.Int64("created_at").Default(0),
		field.Int64("updated_at").Default(0),
		//field.Int64("deleted_at").Default(0),
	}
)
