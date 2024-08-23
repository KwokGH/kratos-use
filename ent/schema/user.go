package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("account").Default("").Comment("用户账号"),
		field.String("password").Default("").Comment("密码"),
		field.String("password_salt").Default("").Comment("加盐"),
		field.String("name").Default("unknown").Comment("用户名称"),
		field.String("mobile").Default("").Comment("手机号"),
	}

	return append(BaseFields, fields...)
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		index.Fields("account", "updated_at").
			Unique(),
		index.Fields("mobile", "updated_at").
			Unique(),
	}
}
