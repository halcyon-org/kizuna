// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/externalinformation"
	"github.com/halcyon-org/kizuna/ent/koyodata"
	"github.com/halcyon-org/kizuna/ent/koyoinformation"
	"github.com/halcyon-org/kizuna/ent/predicate"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// KoyoInformationUpdate is the builder for updating KoyoInformation entities.
type KoyoInformationUpdate struct {
	config
	hooks    []Hook
	mutation *KoyoInformationMutation
}

// Where appends a list predicates to the KoyoInformationUpdate builder.
func (kiu *KoyoInformationUpdate) Where(ps ...predicate.KoyoInformation) *KoyoInformationUpdate {
	kiu.mutation.Where(ps...)
	return kiu
}

// SetName sets the "name" field.
func (kiu *KoyoInformationUpdate) SetName(s string) *KoyoInformationUpdate {
	kiu.mutation.SetName(s)
	return kiu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableName(s *string) *KoyoInformationUpdate {
	if s != nil {
		kiu.SetName(*s)
	}
	return kiu
}

// SetDescription sets the "description" field.
func (kiu *KoyoInformationUpdate) SetDescription(s string) *KoyoInformationUpdate {
	kiu.mutation.SetDescription(s)
	return kiu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableDescription(s *string) *KoyoInformationUpdate {
	if s != nil {
		kiu.SetDescription(*s)
	}
	return kiu
}

// SetParams sets the "params" field.
func (kiu *KoyoInformationUpdate) SetParams(m map[string]string) *KoyoInformationUpdate {
	kiu.mutation.SetParams(m)
	return kiu
}

// SetScales sets the "scales" field.
func (kiu *KoyoInformationUpdate) SetScales(f []float64) *KoyoInformationUpdate {
	kiu.mutation.SetScales(f)
	return kiu
}

// AppendScales appends f to the "scales" field.
func (kiu *KoyoInformationUpdate) AppendScales(f []float64) *KoyoInformationUpdate {
	kiu.mutation.AppendScales(f)
	return kiu
}

// SetVersion sets the "version" field.
func (kiu *KoyoInformationUpdate) SetVersion(s string) *KoyoInformationUpdate {
	kiu.mutation.SetVersion(s)
	return kiu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableVersion(s *string) *KoyoInformationUpdate {
	if s != nil {
		kiu.SetVersion(*s)
	}
	return kiu
}

// SetLicense sets the "license" field.
func (kiu *KoyoInformationUpdate) SetLicense(s string) *KoyoInformationUpdate {
	kiu.mutation.SetLicense(s)
	return kiu
}

// SetNillableLicense sets the "license" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableLicense(s *string) *KoyoInformationUpdate {
	if s != nil {
		kiu.SetLicense(*s)
	}
	return kiu
}

// SetDataType sets the "data_type" field.
func (kiu *KoyoInformationUpdate) SetDataType(kt koyoinformation.DataType) *KoyoInformationUpdate {
	kiu.mutation.SetDataType(kt)
	return kiu
}

// SetNillableDataType sets the "data_type" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableDataType(kt *koyoinformation.DataType) *KoyoInformationUpdate {
	if kt != nil {
		kiu.SetDataType(*kt)
	}
	return kiu
}

// SetLastEntryAt sets the "last_entry_at" field.
func (kiu *KoyoInformationUpdate) SetLastEntryAt(t time.Time) *KoyoInformationUpdate {
	kiu.mutation.SetLastEntryAt(t)
	return kiu
}

// SetNillableLastEntryAt sets the "last_entry_at" field if the given value is not nil.
func (kiu *KoyoInformationUpdate) SetNillableLastEntryAt(t *time.Time) *KoyoInformationUpdate {
	if t != nil {
		kiu.SetLastEntryAt(*t)
	}
	return kiu
}

// SetUpdatedAt sets the "updated_at" field.
func (kiu *KoyoInformationUpdate) SetUpdatedAt(t time.Time) *KoyoInformationUpdate {
	kiu.mutation.SetUpdatedAt(t)
	return kiu
}

// AddExternalIDs adds the "externals" edge to the ExternalInformation entity by IDs.
func (kiu *KoyoInformationUpdate) AddExternalIDs(ids ...pulid.ID) *KoyoInformationUpdate {
	kiu.mutation.AddExternalIDs(ids...)
	return kiu
}

// AddExternals adds the "externals" edges to the ExternalInformation entity.
func (kiu *KoyoInformationUpdate) AddExternals(e ...*ExternalInformation) *KoyoInformationUpdate {
	ids := make([]pulid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return kiu.AddExternalIDs(ids...)
}

// AddDatumIDs adds the "data" edge to the KoyoData entity by IDs.
func (kiu *KoyoInformationUpdate) AddDatumIDs(ids ...pulid.ID) *KoyoInformationUpdate {
	kiu.mutation.AddDatumIDs(ids...)
	return kiu
}

// AddData adds the "data" edges to the KoyoData entity.
func (kiu *KoyoInformationUpdate) AddData(k ...*KoyoData) *KoyoInformationUpdate {
	ids := make([]pulid.ID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kiu.AddDatumIDs(ids...)
}

// Mutation returns the KoyoInformationMutation object of the builder.
func (kiu *KoyoInformationUpdate) Mutation() *KoyoInformationMutation {
	return kiu.mutation
}

// ClearExternals clears all "externals" edges to the ExternalInformation entity.
func (kiu *KoyoInformationUpdate) ClearExternals() *KoyoInformationUpdate {
	kiu.mutation.ClearExternals()
	return kiu
}

// RemoveExternalIDs removes the "externals" edge to ExternalInformation entities by IDs.
func (kiu *KoyoInformationUpdate) RemoveExternalIDs(ids ...pulid.ID) *KoyoInformationUpdate {
	kiu.mutation.RemoveExternalIDs(ids...)
	return kiu
}

// RemoveExternals removes "externals" edges to ExternalInformation entities.
func (kiu *KoyoInformationUpdate) RemoveExternals(e ...*ExternalInformation) *KoyoInformationUpdate {
	ids := make([]pulid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return kiu.RemoveExternalIDs(ids...)
}

// ClearData clears all "data" edges to the KoyoData entity.
func (kiu *KoyoInformationUpdate) ClearData() *KoyoInformationUpdate {
	kiu.mutation.ClearData()
	return kiu
}

// RemoveDatumIDs removes the "data" edge to KoyoData entities by IDs.
func (kiu *KoyoInformationUpdate) RemoveDatumIDs(ids ...pulid.ID) *KoyoInformationUpdate {
	kiu.mutation.RemoveDatumIDs(ids...)
	return kiu
}

// RemoveData removes "data" edges to KoyoData entities.
func (kiu *KoyoInformationUpdate) RemoveData(k ...*KoyoData) *KoyoInformationUpdate {
	ids := make([]pulid.ID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kiu.RemoveDatumIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (kiu *KoyoInformationUpdate) Save(ctx context.Context) (int, error) {
	kiu.defaults()
	return withHooks(ctx, kiu.sqlSave, kiu.mutation, kiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kiu *KoyoInformationUpdate) SaveX(ctx context.Context) int {
	affected, err := kiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (kiu *KoyoInformationUpdate) Exec(ctx context.Context) error {
	_, err := kiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kiu *KoyoInformationUpdate) ExecX(ctx context.Context) {
	if err := kiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kiu *KoyoInformationUpdate) defaults() {
	if _, ok := kiu.mutation.UpdatedAt(); !ok {
		v := koyoinformation.UpdateDefaultUpdatedAt()
		kiu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kiu *KoyoInformationUpdate) check() error {
	if v, ok := kiu.mutation.Name(); ok {
		if err := koyoinformation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.name": %w`, err)}
		}
	}
	if v, ok := kiu.mutation.Version(); ok {
		if err := koyoinformation.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.version": %w`, err)}
		}
	}
	if v, ok := kiu.mutation.License(); ok {
		if err := koyoinformation.LicenseValidator(v); err != nil {
			return &ValidationError{Name: "license", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.license": %w`, err)}
		}
	}
	if v, ok := kiu.mutation.DataType(); ok {
		if err := koyoinformation.DataTypeValidator(v); err != nil {
			return &ValidationError{Name: "data_type", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.data_type": %w`, err)}
		}
	}
	return nil
}

func (kiu *KoyoInformationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := kiu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(koyoinformation.Table, koyoinformation.Columns, sqlgraph.NewFieldSpec(koyoinformation.FieldID, field.TypeString))
	if ps := kiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kiu.mutation.Name(); ok {
		_spec.SetField(koyoinformation.FieldName, field.TypeString, value)
	}
	if value, ok := kiu.mutation.Description(); ok {
		_spec.SetField(koyoinformation.FieldDescription, field.TypeString, value)
	}
	if value, ok := kiu.mutation.Params(); ok {
		_spec.SetField(koyoinformation.FieldParams, field.TypeJSON, value)
	}
	if value, ok := kiu.mutation.Scales(); ok {
		_spec.SetField(koyoinformation.FieldScales, field.TypeJSON, value)
	}
	if value, ok := kiu.mutation.AppendedScales(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, koyoinformation.FieldScales, value)
		})
	}
	if value, ok := kiu.mutation.Version(); ok {
		_spec.SetField(koyoinformation.FieldVersion, field.TypeString, value)
	}
	if value, ok := kiu.mutation.License(); ok {
		_spec.SetField(koyoinformation.FieldLicense, field.TypeString, value)
	}
	if value, ok := kiu.mutation.DataType(); ok {
		_spec.SetField(koyoinformation.FieldDataType, field.TypeEnum, value)
	}
	if value, ok := kiu.mutation.LastEntryAt(); ok {
		_spec.SetField(koyoinformation.FieldLastEntryAt, field.TypeTime, value)
	}
	if value, ok := kiu.mutation.UpdatedAt(); ok {
		_spec.SetField(koyoinformation.FieldUpdatedAt, field.TypeTime, value)
	}
	if kiu.mutation.ExternalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiu.mutation.RemovedExternalsIDs(); len(nodes) > 0 && !kiu.mutation.ExternalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiu.mutation.ExternalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kiu.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiu.mutation.RemovedDataIDs(); len(nodes) > 0 && !kiu.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiu.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, kiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{koyoinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	kiu.mutation.done = true
	return n, nil
}

// KoyoInformationUpdateOne is the builder for updating a single KoyoInformation entity.
type KoyoInformationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *KoyoInformationMutation
}

// SetName sets the "name" field.
func (kiuo *KoyoInformationUpdateOne) SetName(s string) *KoyoInformationUpdateOne {
	kiuo.mutation.SetName(s)
	return kiuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableName(s *string) *KoyoInformationUpdateOne {
	if s != nil {
		kiuo.SetName(*s)
	}
	return kiuo
}

// SetDescription sets the "description" field.
func (kiuo *KoyoInformationUpdateOne) SetDescription(s string) *KoyoInformationUpdateOne {
	kiuo.mutation.SetDescription(s)
	return kiuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableDescription(s *string) *KoyoInformationUpdateOne {
	if s != nil {
		kiuo.SetDescription(*s)
	}
	return kiuo
}

// SetParams sets the "params" field.
func (kiuo *KoyoInformationUpdateOne) SetParams(m map[string]string) *KoyoInformationUpdateOne {
	kiuo.mutation.SetParams(m)
	return kiuo
}

// SetScales sets the "scales" field.
func (kiuo *KoyoInformationUpdateOne) SetScales(f []float64) *KoyoInformationUpdateOne {
	kiuo.mutation.SetScales(f)
	return kiuo
}

// AppendScales appends f to the "scales" field.
func (kiuo *KoyoInformationUpdateOne) AppendScales(f []float64) *KoyoInformationUpdateOne {
	kiuo.mutation.AppendScales(f)
	return kiuo
}

// SetVersion sets the "version" field.
func (kiuo *KoyoInformationUpdateOne) SetVersion(s string) *KoyoInformationUpdateOne {
	kiuo.mutation.SetVersion(s)
	return kiuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableVersion(s *string) *KoyoInformationUpdateOne {
	if s != nil {
		kiuo.SetVersion(*s)
	}
	return kiuo
}

// SetLicense sets the "license" field.
func (kiuo *KoyoInformationUpdateOne) SetLicense(s string) *KoyoInformationUpdateOne {
	kiuo.mutation.SetLicense(s)
	return kiuo
}

// SetNillableLicense sets the "license" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableLicense(s *string) *KoyoInformationUpdateOne {
	if s != nil {
		kiuo.SetLicense(*s)
	}
	return kiuo
}

// SetDataType sets the "data_type" field.
func (kiuo *KoyoInformationUpdateOne) SetDataType(kt koyoinformation.DataType) *KoyoInformationUpdateOne {
	kiuo.mutation.SetDataType(kt)
	return kiuo
}

// SetNillableDataType sets the "data_type" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableDataType(kt *koyoinformation.DataType) *KoyoInformationUpdateOne {
	if kt != nil {
		kiuo.SetDataType(*kt)
	}
	return kiuo
}

// SetLastEntryAt sets the "last_entry_at" field.
func (kiuo *KoyoInformationUpdateOne) SetLastEntryAt(t time.Time) *KoyoInformationUpdateOne {
	kiuo.mutation.SetLastEntryAt(t)
	return kiuo
}

// SetNillableLastEntryAt sets the "last_entry_at" field if the given value is not nil.
func (kiuo *KoyoInformationUpdateOne) SetNillableLastEntryAt(t *time.Time) *KoyoInformationUpdateOne {
	if t != nil {
		kiuo.SetLastEntryAt(*t)
	}
	return kiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (kiuo *KoyoInformationUpdateOne) SetUpdatedAt(t time.Time) *KoyoInformationUpdateOne {
	kiuo.mutation.SetUpdatedAt(t)
	return kiuo
}

// AddExternalIDs adds the "externals" edge to the ExternalInformation entity by IDs.
func (kiuo *KoyoInformationUpdateOne) AddExternalIDs(ids ...pulid.ID) *KoyoInformationUpdateOne {
	kiuo.mutation.AddExternalIDs(ids...)
	return kiuo
}

// AddExternals adds the "externals" edges to the ExternalInformation entity.
func (kiuo *KoyoInformationUpdateOne) AddExternals(e ...*ExternalInformation) *KoyoInformationUpdateOne {
	ids := make([]pulid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return kiuo.AddExternalIDs(ids...)
}

// AddDatumIDs adds the "data" edge to the KoyoData entity by IDs.
func (kiuo *KoyoInformationUpdateOne) AddDatumIDs(ids ...pulid.ID) *KoyoInformationUpdateOne {
	kiuo.mutation.AddDatumIDs(ids...)
	return kiuo
}

// AddData adds the "data" edges to the KoyoData entity.
func (kiuo *KoyoInformationUpdateOne) AddData(k ...*KoyoData) *KoyoInformationUpdateOne {
	ids := make([]pulid.ID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kiuo.AddDatumIDs(ids...)
}

// Mutation returns the KoyoInformationMutation object of the builder.
func (kiuo *KoyoInformationUpdateOne) Mutation() *KoyoInformationMutation {
	return kiuo.mutation
}

// ClearExternals clears all "externals" edges to the ExternalInformation entity.
func (kiuo *KoyoInformationUpdateOne) ClearExternals() *KoyoInformationUpdateOne {
	kiuo.mutation.ClearExternals()
	return kiuo
}

// RemoveExternalIDs removes the "externals" edge to ExternalInformation entities by IDs.
func (kiuo *KoyoInformationUpdateOne) RemoveExternalIDs(ids ...pulid.ID) *KoyoInformationUpdateOne {
	kiuo.mutation.RemoveExternalIDs(ids...)
	return kiuo
}

// RemoveExternals removes "externals" edges to ExternalInformation entities.
func (kiuo *KoyoInformationUpdateOne) RemoveExternals(e ...*ExternalInformation) *KoyoInformationUpdateOne {
	ids := make([]pulid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return kiuo.RemoveExternalIDs(ids...)
}

// ClearData clears all "data" edges to the KoyoData entity.
func (kiuo *KoyoInformationUpdateOne) ClearData() *KoyoInformationUpdateOne {
	kiuo.mutation.ClearData()
	return kiuo
}

// RemoveDatumIDs removes the "data" edge to KoyoData entities by IDs.
func (kiuo *KoyoInformationUpdateOne) RemoveDatumIDs(ids ...pulid.ID) *KoyoInformationUpdateOne {
	kiuo.mutation.RemoveDatumIDs(ids...)
	return kiuo
}

// RemoveData removes "data" edges to KoyoData entities.
func (kiuo *KoyoInformationUpdateOne) RemoveData(k ...*KoyoData) *KoyoInformationUpdateOne {
	ids := make([]pulid.ID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kiuo.RemoveDatumIDs(ids...)
}

// Where appends a list predicates to the KoyoInformationUpdate builder.
func (kiuo *KoyoInformationUpdateOne) Where(ps ...predicate.KoyoInformation) *KoyoInformationUpdateOne {
	kiuo.mutation.Where(ps...)
	return kiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (kiuo *KoyoInformationUpdateOne) Select(field string, fields ...string) *KoyoInformationUpdateOne {
	kiuo.fields = append([]string{field}, fields...)
	return kiuo
}

// Save executes the query and returns the updated KoyoInformation entity.
func (kiuo *KoyoInformationUpdateOne) Save(ctx context.Context) (*KoyoInformation, error) {
	kiuo.defaults()
	return withHooks(ctx, kiuo.sqlSave, kiuo.mutation, kiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (kiuo *KoyoInformationUpdateOne) SaveX(ctx context.Context) *KoyoInformation {
	node, err := kiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (kiuo *KoyoInformationUpdateOne) Exec(ctx context.Context) error {
	_, err := kiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kiuo *KoyoInformationUpdateOne) ExecX(ctx context.Context) {
	if err := kiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kiuo *KoyoInformationUpdateOne) defaults() {
	if _, ok := kiuo.mutation.UpdatedAt(); !ok {
		v := koyoinformation.UpdateDefaultUpdatedAt()
		kiuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kiuo *KoyoInformationUpdateOne) check() error {
	if v, ok := kiuo.mutation.Name(); ok {
		if err := koyoinformation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.name": %w`, err)}
		}
	}
	if v, ok := kiuo.mutation.Version(); ok {
		if err := koyoinformation.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.version": %w`, err)}
		}
	}
	if v, ok := kiuo.mutation.License(); ok {
		if err := koyoinformation.LicenseValidator(v); err != nil {
			return &ValidationError{Name: "license", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.license": %w`, err)}
		}
	}
	if v, ok := kiuo.mutation.DataType(); ok {
		if err := koyoinformation.DataTypeValidator(v); err != nil {
			return &ValidationError{Name: "data_type", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.data_type": %w`, err)}
		}
	}
	return nil
}

func (kiuo *KoyoInformationUpdateOne) sqlSave(ctx context.Context) (_node *KoyoInformation, err error) {
	if err := kiuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(koyoinformation.Table, koyoinformation.Columns, sqlgraph.NewFieldSpec(koyoinformation.FieldID, field.TypeString))
	id, ok := kiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "KoyoInformation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := kiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, koyoinformation.FieldID)
		for _, f := range fields {
			if !koyoinformation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != koyoinformation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := kiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := kiuo.mutation.Name(); ok {
		_spec.SetField(koyoinformation.FieldName, field.TypeString, value)
	}
	if value, ok := kiuo.mutation.Description(); ok {
		_spec.SetField(koyoinformation.FieldDescription, field.TypeString, value)
	}
	if value, ok := kiuo.mutation.Params(); ok {
		_spec.SetField(koyoinformation.FieldParams, field.TypeJSON, value)
	}
	if value, ok := kiuo.mutation.Scales(); ok {
		_spec.SetField(koyoinformation.FieldScales, field.TypeJSON, value)
	}
	if value, ok := kiuo.mutation.AppendedScales(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, koyoinformation.FieldScales, value)
		})
	}
	if value, ok := kiuo.mutation.Version(); ok {
		_spec.SetField(koyoinformation.FieldVersion, field.TypeString, value)
	}
	if value, ok := kiuo.mutation.License(); ok {
		_spec.SetField(koyoinformation.FieldLicense, field.TypeString, value)
	}
	if value, ok := kiuo.mutation.DataType(); ok {
		_spec.SetField(koyoinformation.FieldDataType, field.TypeEnum, value)
	}
	if value, ok := kiuo.mutation.LastEntryAt(); ok {
		_spec.SetField(koyoinformation.FieldLastEntryAt, field.TypeTime, value)
	}
	if value, ok := kiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(koyoinformation.FieldUpdatedAt, field.TypeTime, value)
	}
	if kiuo.mutation.ExternalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiuo.mutation.RemovedExternalsIDs(); len(nodes) > 0 && !kiuo.mutation.ExternalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiuo.mutation.ExternalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.ExternalsTable,
			Columns: []string{koyoinformation.ExternalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(externalinformation.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if kiuo.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiuo.mutation.RemovedDataIDs(); len(nodes) > 0 && !kiuo.mutation.DataCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := kiuo.mutation.DataIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   koyoinformation.DataTable,
			Columns: []string{koyoinformation.DataColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(koyodata.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &KoyoInformation{config: kiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, kiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{koyoinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	kiuo.mutation.done = true
	return _node, nil
}