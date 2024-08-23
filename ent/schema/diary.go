package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"kratos-use/ent/schema/dbo"
)

// Diary holds the schema definition for the Diary entity.
type Diary struct {
	ent.Schema
}

// Fields of the Diary.
func (Diary) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("title").Default(""),
		field.Text("content").Default(""),
		field.Int64("belong_at").Default(0),
		field.String("user_id").Default(""),
		field.String("tag").Default("-"),
	}

	return append(BaseFields, fields...)
}

// Edges of the Diary.
func (Diary) Edges() []ent.Edge {
	return nil
}

func (Diary) Indexes() []ent.Index {
	return []ent.Index{
		// 唯一约束索引
		//index.Fields("email", "deleted_at").
		//	Unique(),
	}
}

func (Diary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		dbo.SoftDeleteMixin{},
	}
}
