// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"kratos-use/ent/diary"
	"kratos-use/ent/schema"
	"kratos-use/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	diaryMixin := schema.Diary{}.Mixin()
	diaryMixinHooks0 := diaryMixin[0].Hooks()
	diary.Hooks[0] = diaryMixinHooks0[0]
	diaryMixinInters0 := diaryMixin[0].Interceptors()
	diary.Interceptors[0] = diaryMixinInters0[0]
	diaryFields := schema.Diary{}.Fields()
	_ = diaryFields
	// diaryDescCreatedAt is the schema descriptor for created_at field.
	diaryDescCreatedAt := diaryFields[1].Descriptor()
	// diary.DefaultCreatedAt holds the default value on creation for the created_at field.
	diary.DefaultCreatedAt = diaryDescCreatedAt.Default.(int64)
	// diaryDescUpdatedAt is the schema descriptor for updated_at field.
	diaryDescUpdatedAt := diaryFields[2].Descriptor()
	// diary.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	diary.DefaultUpdatedAt = diaryDescUpdatedAt.Default.(int64)
	// diaryDescTitle is the schema descriptor for title field.
	diaryDescTitle := diaryFields[3].Descriptor()
	// diary.DefaultTitle holds the default value on creation for the title field.
	diary.DefaultTitle = diaryDescTitle.Default.(string)
	// diaryDescContent is the schema descriptor for content field.
	diaryDescContent := diaryFields[4].Descriptor()
	// diary.DefaultContent holds the default value on creation for the content field.
	diary.DefaultContent = diaryDescContent.Default.(string)
	// diaryDescBelongAt is the schema descriptor for belong_at field.
	diaryDescBelongAt := diaryFields[5].Descriptor()
	// diary.DefaultBelongAt holds the default value on creation for the belong_at field.
	diary.DefaultBelongAt = diaryDescBelongAt.Default.(int64)
	// diaryDescUserID is the schema descriptor for user_id field.
	diaryDescUserID := diaryFields[6].Descriptor()
	// diary.DefaultUserID holds the default value on creation for the user_id field.
	diary.DefaultUserID = diaryDescUserID.Default.(string)
	// diaryDescTag is the schema descriptor for tag field.
	diaryDescTag := diaryFields[7].Descriptor()
	// diary.DefaultTag holds the default value on creation for the tag field.
	diary.DefaultTag = diaryDescTag.Default.(string)
	// diaryDescID is the schema descriptor for id field.
	diaryDescID := diaryFields[0].Descriptor()
	// diary.DefaultID holds the default value on creation for the id field.
	diary.DefaultID = diaryDescID.Default.(func() string)
	// diary.IDValidator is a validator for the "id" field. It is called by the builders before save.
	diary.IDValidator = func() func(string) error {
		validators := diaryDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(int64)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(int64)
	// userDescAccount is the schema descriptor for account field.
	userDescAccount := userFields[3].Descriptor()
	// user.DefaultAccount holds the default value on creation for the account field.
	user.DefaultAccount = userDescAccount.Default.(string)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[4].Descriptor()
	// user.DefaultPassword holds the default value on creation for the password field.
	user.DefaultPassword = userDescPassword.Default.(string)
	// userDescPasswordSalt is the schema descriptor for password_salt field.
	userDescPasswordSalt := userFields[5].Descriptor()
	// user.DefaultPasswordSalt holds the default value on creation for the password_salt field.
	user.DefaultPasswordSalt = userDescPasswordSalt.Default.(string)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[6].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescMobile is the schema descriptor for mobile field.
	userDescMobile := userFields[7].Descriptor()
	// user.DefaultMobile holds the default value on creation for the mobile field.
	user.DefaultMobile = userDescMobile.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() string)
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = func() func(string) error {
		validators := userDescID.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(id string) error {
			for _, fn := range fns {
				if err := fn(id); err != nil {
					return err
				}
			}
			return nil
		}
	}()
}

const (
	Version = "v0.13.1"                                         // Version of ent codegen.
	Sum     = "h1:uD8QwN1h6SNphdCCzmkMN3feSUzNnVvV/WIkHKMbzOE=" // Sum of ent codegen.
)
