// Code generated by ent, DO NOT EDIT.

package externalinformation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the externalinformation type in the database.
	Label = "external_information"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldLicense holds the string denoting the license field in the database.
	FieldLicense = "license"
	// FieldLicenseDescription holds the string denoting the license_description field in the database.
	FieldLicenseDescription = "license_description"
	// FieldFirstEntryAt holds the string denoting the first_entry_at field in the database.
	FieldFirstEntryAt = "first_entry_at"
	// FieldLastUpdatedAt holds the string denoting the last_updated_at field in the database.
	FieldLastUpdatedAt = "last_updated_at"
	// Table holds the table name of the externalinformation in the database.
	Table = "external_informations"
)

// Columns holds all SQL columns for externalinformation fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldLicense,
	FieldLicenseDescription,
	FieldFirstEntryAt,
	FieldLastUpdatedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "external_informations"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"koyo_information_externals",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// LicenseValidator is a validator for the "license" field. It is called by the builders before save.
	LicenseValidator func(string) error
	// DefaultFirstEntryAt holds the default value on creation for the "first_entry_at" field.
	DefaultFirstEntryAt func() time.Time
	// DefaultLastUpdatedAt holds the default value on creation for the "last_updated_at" field.
	DefaultLastUpdatedAt func() time.Time
	// UpdateDefaultLastUpdatedAt holds the default value on update for the "last_updated_at" field.
	UpdateDefaultLastUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// OrderOption defines the ordering options for the ExternalInformation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByLicense orders the results by the license field.
func ByLicense(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicense, opts...).ToFunc()
}

// ByLicenseDescription orders the results by the license_description field.
func ByLicenseDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicenseDescription, opts...).ToFunc()
}

// ByFirstEntryAt orders the results by the first_entry_at field.
func ByFirstEntryAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstEntryAt, opts...).ToFunc()
}

// ByLastUpdatedAt orders the results by the last_updated_at field.
func ByLastUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdatedAt, opts...).ToFunc()
}