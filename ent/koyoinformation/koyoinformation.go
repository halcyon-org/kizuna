// Code generated by ent, DO NOT EDIT.

package koyoinformation

import (
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

const (
	// Label holds the string label denoting the koyoinformation type in the database.
	Label = "koyo_information"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldParams holds the string denoting the params field in the database.
	FieldParams = "params"
	// FieldScales holds the string denoting the scales field in the database.
	FieldScales = "scales"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// FieldLicense holds the string denoting the license field in the database.
	FieldLicense = "license"
	// FieldDataType holds the string denoting the data_type field in the database.
	FieldDataType = "data_type"
	// FieldFirstEntryAt holds the string denoting the first_entry_at field in the database.
	FieldFirstEntryAt = "first_entry_at"
	// FieldLastEntryAt holds the string denoting the last_entry_at field in the database.
	FieldLastEntryAt = "last_entry_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeExternals holds the string denoting the externals edge name in mutations.
	EdgeExternals = "externals"
	// EdgeData holds the string denoting the data edge name in mutations.
	EdgeData = "data"
	// Table holds the table name of the koyoinformation in the database.
	Table = "koyo_informations"
	// ExternalsTable is the table that holds the externals relation/edge.
	ExternalsTable = "external_informations"
	// ExternalsInverseTable is the table name for the ExternalInformation entity.
	// It exists in this package in order to avoid circular dependency with the "externalinformation" package.
	ExternalsInverseTable = "external_informations"
	// ExternalsColumn is the table column denoting the externals relation/edge.
	ExternalsColumn = "koyo_information_externals"
	// DataTable is the table that holds the data relation/edge.
	DataTable = "koyo_data"
	// DataInverseTable is the table name for the KoyoData entity.
	// It exists in this package in order to avoid circular dependency with the "koyodata" package.
	DataInverseTable = "koyo_data"
	// DataColumn is the table column denoting the data relation/edge.
	DataColumn = "koyo_information_data"
)

// Columns holds all SQL columns for koyoinformation fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldParams,
	FieldScales,
	FieldVersion,
	FieldLicense,
	FieldDataType,
	FieldFirstEntryAt,
	FieldLastEntryAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// VersionValidator is a validator for the "version" field. It is called by the builders before save.
	VersionValidator func(string) error
	// LicenseValidator is a validator for the "license" field. It is called by the builders before save.
	LicenseValidator func(string) error
	// DefaultFirstEntryAt holds the default value on creation for the "first_entry_at" field.
	DefaultFirstEntryAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() pulid.ID
)

// DataType defines the type for the "data_type" enum field.
type DataType string

// DataType values.
const (
	DataTypeUnspecified DataType = "unspecified"
	DataTypeImage       DataType = "image"
	DataTypeCsv         DataType = "csv"
	DataTypeJSON        DataType = "json"
)

func (dt DataType) String() string {
	return string(dt)
}

// DataTypeValidator is a validator for the "data_type" field enum values. It is called by the builders before save.
func DataTypeValidator(dt DataType) error {
	switch dt {
	case DataTypeUnspecified, DataTypeImage, DataTypeCsv, DataTypeJSON:
		return nil
	default:
		return fmt.Errorf("koyoinformation: invalid enum value for data_type field: %q", dt)
	}
}

// OrderOption defines the ordering options for the KoyoInformation queries.
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

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}

// ByLicense orders the results by the license field.
func ByLicense(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLicense, opts...).ToFunc()
}

// ByDataType orders the results by the data_type field.
func ByDataType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDataType, opts...).ToFunc()
}

// ByFirstEntryAt orders the results by the first_entry_at field.
func ByFirstEntryAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstEntryAt, opts...).ToFunc()
}

// ByLastEntryAt orders the results by the last_entry_at field.
func ByLastEntryAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastEntryAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByExternalsCount orders the results by externals count.
func ByExternalsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newExternalsStep(), opts...)
	}
}

// ByExternals orders the results by externals terms.
func ByExternals(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExternalsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDataCount orders the results by data count.
func ByDataCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDataStep(), opts...)
	}
}

// ByData orders the results by data terms.
func ByData(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDataStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newExternalsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExternalsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ExternalsTable, ExternalsColumn),
	)
}
func newDataStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DataInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DataTable, DataColumn),
	)
}