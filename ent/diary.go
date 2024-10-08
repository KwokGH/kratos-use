// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"kratos-use/ent/diary"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Diary is the model entity for the Diary schema.
type Diary struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt int64 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// BelongAt holds the value of the "belong_at" field.
	BelongAt int64 `json:"belong_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag          string `json:"tag,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Diary) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case diary.FieldDeletedAt, diary.FieldCreatedAt, diary.FieldUpdatedAt, diary.FieldBelongAt:
			values[i] = new(sql.NullInt64)
		case diary.FieldID, diary.FieldTitle, diary.FieldContent, diary.FieldUserID, diary.FieldTag:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Diary fields.
func (d *Diary) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case diary.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				d.ID = value.String
			}
		case diary.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				d.DeletedAt = value.Int64
			}
		case diary.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				d.CreatedAt = value.Int64
			}
		case diary.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				d.UpdatedAt = value.Int64
			}
		case diary.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				d.Title = value.String
			}
		case diary.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				d.Content = value.String
			}
		case diary.FieldBelongAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field belong_at", values[i])
			} else if value.Valid {
				d.BelongAt = value.Int64
			}
		case diary.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				d.UserID = value.String
			}
		case diary.FieldTag:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tag", values[i])
			} else if value.Valid {
				d.Tag = value.String
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Diary.
// This includes values selected through modifiers, order, etc.
func (d *Diary) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// Update returns a builder for updating this Diary.
// Note that you need to call Diary.Unwrap() before calling this method if this Diary
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Diary) Update() *DiaryUpdateOne {
	return NewDiaryClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Diary entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Diary) Unwrap() *Diary {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Diary is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Diary) String() string {
	var builder strings.Builder
	builder.WriteString("Diary(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", d.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", d.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", d.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(d.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(d.Content)
	builder.WriteString(", ")
	builder.WriteString("belong_at=")
	builder.WriteString(fmt.Sprintf("%v", d.BelongAt))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(d.UserID)
	builder.WriteString(", ")
	builder.WriteString("tag=")
	builder.WriteString(d.Tag)
	builder.WriteByte(')')
	return builder.String()
}

// Diaries is a parsable slice of Diary.
type Diaries []*Diary
