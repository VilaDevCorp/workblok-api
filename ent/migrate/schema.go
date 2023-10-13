// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlocksColumns holds the columns for the "blocks" table.
	BlocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "finish_date", Type: field.TypeTime, Nullable: true},
		{Name: "target_minutes", Type: field.TypeInt, Default: 5},
		{Name: "distraction_minutes", Type: field.TypeInt, Default: 0},
		{Name: "user_blocks", Type: field.TypeUUID},
	}
	// BlocksTable holds the schema information for the "blocks" table.
	BlocksTable = &schema.Table{
		Name:       "blocks",
		Columns:    BlocksColumns,
		PrimaryKey: []*schema.Column{BlocksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "blocks_users_blocks",
				Columns:    []*schema.Column{BlocksColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "email_valid", Type: field.TypeBool, Default: false},
		{Name: "config", Type: field.TypeJSON, Nullable: true},
		{Name: "tutorial_completed", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// VerificationCodesColumns holds the columns for the "verification_codes" table.
	VerificationCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "creation_date", Type: field.TypeTime},
		{Name: "type", Type: field.TypeString},
		{Name: "code", Type: field.TypeString},
		{Name: "expire_date", Type: field.TypeTime},
		{Name: "valid", Type: field.TypeBool},
		{Name: "user_codes", Type: field.TypeUUID},
	}
	// VerificationCodesTable holds the schema information for the "verification_codes" table.
	VerificationCodesTable = &schema.Table{
		Name:       "verification_codes",
		Columns:    VerificationCodesColumns,
		PrimaryKey: []*schema.Column{VerificationCodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "verification_codes_users_codes",
				Columns:    []*schema.Column{VerificationCodesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlocksTable,
		UsersTable,
		VerificationCodesTable,
	}
)

func init() {
	BlocksTable.ForeignKeys[0].RefTable = UsersTable
	VerificationCodesTable.ForeignKeys[0].RefTable = UsersTable
	VerificationCodesTable.Annotation = &entsql.Annotation{
		Table: "verification_codes",
	}
}
