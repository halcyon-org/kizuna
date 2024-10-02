// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/halcyon-org/kizuna/ent/externalinformation"
	"github.com/halcyon-org/kizuna/ent/koyodata"
	"github.com/halcyon-org/kizuna/ent/koyoinformation"
	"github.com/halcyon-org/kizuna/ent/schema/pulid"
)

// KoyoInformationCreate is the builder for creating a KoyoInformation entity.
type KoyoInformationCreate struct {
	config
	mutation *KoyoInformationMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (kic *KoyoInformationCreate) SetName(s string) *KoyoInformationCreate {
	kic.mutation.SetName(s)
	return kic
}

// SetDescription sets the "description" field.
func (kic *KoyoInformationCreate) SetDescription(s string) *KoyoInformationCreate {
	kic.mutation.SetDescription(s)
	return kic
}

// SetParams sets the "params" field.
func (kic *KoyoInformationCreate) SetParams(m map[string]string) *KoyoInformationCreate {
	kic.mutation.SetParams(m)
	return kic
}

// SetScales sets the "scales" field.
func (kic *KoyoInformationCreate) SetScales(f []float64) *KoyoInformationCreate {
	kic.mutation.SetScales(f)
	return kic
}

// SetVersion sets the "version" field.
func (kic *KoyoInformationCreate) SetVersion(s string) *KoyoInformationCreate {
	kic.mutation.SetVersion(s)
	return kic
}

// SetLicense sets the "license" field.
func (kic *KoyoInformationCreate) SetLicense(s string) *KoyoInformationCreate {
	kic.mutation.SetLicense(s)
	return kic
}

// SetDataType sets the "data_type" field.
func (kic *KoyoInformationCreate) SetDataType(kt koyoinformation.DataType) *KoyoInformationCreate {
	kic.mutation.SetDataType(kt)
	return kic
}

// SetFirstEntryAt sets the "first_entry_at" field.
func (kic *KoyoInformationCreate) SetFirstEntryAt(t time.Time) *KoyoInformationCreate {
	kic.mutation.SetFirstEntryAt(t)
	return kic
}

// SetNillableFirstEntryAt sets the "first_entry_at" field if the given value is not nil.
func (kic *KoyoInformationCreate) SetNillableFirstEntryAt(t *time.Time) *KoyoInformationCreate {
	if t != nil {
		kic.SetFirstEntryAt(*t)
	}
	return kic
}

// SetLastEntryAt sets the "last_entry_at" field.
func (kic *KoyoInformationCreate) SetLastEntryAt(t time.Time) *KoyoInformationCreate {
	kic.mutation.SetLastEntryAt(t)
	return kic
}

// SetUpdatedAt sets the "updated_at" field.
func (kic *KoyoInformationCreate) SetUpdatedAt(t time.Time) *KoyoInformationCreate {
	kic.mutation.SetUpdatedAt(t)
	return kic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (kic *KoyoInformationCreate) SetNillableUpdatedAt(t *time.Time) *KoyoInformationCreate {
	if t != nil {
		kic.SetUpdatedAt(*t)
	}
	return kic
}

// SetID sets the "id" field.
func (kic *KoyoInformationCreate) SetID(pu pulid.ID) *KoyoInformationCreate {
	kic.mutation.SetID(pu)
	return kic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (kic *KoyoInformationCreate) SetNillableID(pu *pulid.ID) *KoyoInformationCreate {
	if pu != nil {
		kic.SetID(*pu)
	}
	return kic
}

// AddExternalIDs adds the "externals" edge to the ExternalInformation entity by IDs.
func (kic *KoyoInformationCreate) AddExternalIDs(ids ...pulid.ID) *KoyoInformationCreate {
	kic.mutation.AddExternalIDs(ids...)
	return kic
}

// AddExternals adds the "externals" edges to the ExternalInformation entity.
func (kic *KoyoInformationCreate) AddExternals(e ...*ExternalInformation) *KoyoInformationCreate {
	ids := make([]pulid.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return kic.AddExternalIDs(ids...)
}

// AddDatumIDs adds the "data" edge to the KoyoData entity by IDs.
func (kic *KoyoInformationCreate) AddDatumIDs(ids ...pulid.ID) *KoyoInformationCreate {
	kic.mutation.AddDatumIDs(ids...)
	return kic
}

// AddData adds the "data" edges to the KoyoData entity.
func (kic *KoyoInformationCreate) AddData(k ...*KoyoData) *KoyoInformationCreate {
	ids := make([]pulid.ID, len(k))
	for i := range k {
		ids[i] = k[i].ID
	}
	return kic.AddDatumIDs(ids...)
}

// Mutation returns the KoyoInformationMutation object of the builder.
func (kic *KoyoInformationCreate) Mutation() *KoyoInformationMutation {
	return kic.mutation
}

// Save creates the KoyoInformation in the database.
func (kic *KoyoInformationCreate) Save(ctx context.Context) (*KoyoInformation, error) {
	kic.defaults()
	return withHooks(ctx, kic.sqlSave, kic.mutation, kic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (kic *KoyoInformationCreate) SaveX(ctx context.Context) *KoyoInformation {
	v, err := kic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kic *KoyoInformationCreate) Exec(ctx context.Context) error {
	_, err := kic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kic *KoyoInformationCreate) ExecX(ctx context.Context) {
	if err := kic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (kic *KoyoInformationCreate) defaults() {
	if _, ok := kic.mutation.FirstEntryAt(); !ok {
		v := koyoinformation.DefaultFirstEntryAt()
		kic.mutation.SetFirstEntryAt(v)
	}
	if _, ok := kic.mutation.UpdatedAt(); !ok {
		v := koyoinformation.DefaultUpdatedAt()
		kic.mutation.SetUpdatedAt(v)
	}
	if _, ok := kic.mutation.ID(); !ok {
		v := koyoinformation.DefaultID()
		kic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (kic *KoyoInformationCreate) check() error {
	if _, ok := kic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "KoyoInformation.name"`)}
	}
	if v, ok := kic.mutation.Name(); ok {
		if err := koyoinformation.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.name": %w`, err)}
		}
	}
	if _, ok := kic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "KoyoInformation.description"`)}
	}
	if _, ok := kic.mutation.Params(); !ok {
		return &ValidationError{Name: "params", err: errors.New(`ent: missing required field "KoyoInformation.params"`)}
	}
	if _, ok := kic.mutation.Scales(); !ok {
		return &ValidationError{Name: "scales", err: errors.New(`ent: missing required field "KoyoInformation.scales"`)}
	}
	if _, ok := kic.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "KoyoInformation.version"`)}
	}
	if v, ok := kic.mutation.Version(); ok {
		if err := koyoinformation.VersionValidator(v); err != nil {
			return &ValidationError{Name: "version", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.version": %w`, err)}
		}
	}
	if _, ok := kic.mutation.License(); !ok {
		return &ValidationError{Name: "license", err: errors.New(`ent: missing required field "KoyoInformation.license"`)}
	}
	if v, ok := kic.mutation.License(); ok {
		if err := koyoinformation.LicenseValidator(v); err != nil {
			return &ValidationError{Name: "license", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.license": %w`, err)}
		}
	}
	if _, ok := kic.mutation.DataType(); !ok {
		return &ValidationError{Name: "data_type", err: errors.New(`ent: missing required field "KoyoInformation.data_type"`)}
	}
	if v, ok := kic.mutation.DataType(); ok {
		if err := koyoinformation.DataTypeValidator(v); err != nil {
			return &ValidationError{Name: "data_type", err: fmt.Errorf(`ent: validator failed for field "KoyoInformation.data_type": %w`, err)}
		}
	}
	if _, ok := kic.mutation.FirstEntryAt(); !ok {
		return &ValidationError{Name: "first_entry_at", err: errors.New(`ent: missing required field "KoyoInformation.first_entry_at"`)}
	}
	if _, ok := kic.mutation.LastEntryAt(); !ok {
		return &ValidationError{Name: "last_entry_at", err: errors.New(`ent: missing required field "KoyoInformation.last_entry_at"`)}
	}
	if _, ok := kic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "KoyoInformation.updated_at"`)}
	}
	return nil
}

func (kic *KoyoInformationCreate) sqlSave(ctx context.Context) (*KoyoInformation, error) {
	if err := kic.check(); err != nil {
		return nil, err
	}
	_node, _spec := kic.createSpec()
	if err := sqlgraph.CreateNode(ctx, kic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*pulid.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	kic.mutation.id = &_node.ID
	kic.mutation.done = true
	return _node, nil
}

func (kic *KoyoInformationCreate) createSpec() (*KoyoInformation, *sqlgraph.CreateSpec) {
	var (
		_node = &KoyoInformation{config: kic.config}
		_spec = sqlgraph.NewCreateSpec(koyoinformation.Table, sqlgraph.NewFieldSpec(koyoinformation.FieldID, field.TypeString))
	)
	if id, ok := kic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := kic.mutation.Name(); ok {
		_spec.SetField(koyoinformation.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := kic.mutation.Description(); ok {
		_spec.SetField(koyoinformation.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := kic.mutation.Params(); ok {
		_spec.SetField(koyoinformation.FieldParams, field.TypeJSON, value)
		_node.Params = value
	}
	if value, ok := kic.mutation.Scales(); ok {
		_spec.SetField(koyoinformation.FieldScales, field.TypeJSON, value)
		_node.Scales = value
	}
	if value, ok := kic.mutation.Version(); ok {
		_spec.SetField(koyoinformation.FieldVersion, field.TypeString, value)
		_node.Version = value
	}
	if value, ok := kic.mutation.License(); ok {
		_spec.SetField(koyoinformation.FieldLicense, field.TypeString, value)
		_node.License = value
	}
	if value, ok := kic.mutation.DataType(); ok {
		_spec.SetField(koyoinformation.FieldDataType, field.TypeEnum, value)
		_node.DataType = value
	}
	if value, ok := kic.mutation.FirstEntryAt(); ok {
		_spec.SetField(koyoinformation.FieldFirstEntryAt, field.TypeTime, value)
		_node.FirstEntryAt = value
	}
	if value, ok := kic.mutation.LastEntryAt(); ok {
		_spec.SetField(koyoinformation.FieldLastEntryAt, field.TypeTime, value)
		_node.LastEntryAt = value
	}
	if value, ok := kic.mutation.UpdatedAt(); ok {
		_spec.SetField(koyoinformation.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := kic.mutation.ExternalsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := kic.mutation.DataIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// KoyoInformationCreateBulk is the builder for creating many KoyoInformation entities in bulk.
type KoyoInformationCreateBulk struct {
	config
	err      error
	builders []*KoyoInformationCreate
}

// Save creates the KoyoInformation entities in the database.
func (kicb *KoyoInformationCreateBulk) Save(ctx context.Context) ([]*KoyoInformation, error) {
	if kicb.err != nil {
		return nil, kicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(kicb.builders))
	nodes := make([]*KoyoInformation, len(kicb.builders))
	mutators := make([]Mutator, len(kicb.builders))
	for i := range kicb.builders {
		func(i int, root context.Context) {
			builder := kicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*KoyoInformationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, kicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, kicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, kicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (kicb *KoyoInformationCreateBulk) SaveX(ctx context.Context) []*KoyoInformation {
	v, err := kicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (kicb *KoyoInformationCreateBulk) Exec(ctx context.Context) error {
	_, err := kicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (kicb *KoyoInformationCreateBulk) ExecX(ctx context.Context) {
	if err := kicb.Exec(ctx); err != nil {
		panic(err)
	}
}
