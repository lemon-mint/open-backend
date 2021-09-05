// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lemon-mint/open-backend/ent/group"
	"github.com/lemon-mint/open-backend/ent/resource"
)

// Resource is the model entity for the Resource schema.
type Resource struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Group holds the value of the "group" field.
	Group []string `json:"group,omitempty"`
	// Others holds the value of the "others" field.
	Others []string `json:"others,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ResourceQuery when eager-loading is set.
	Edges           ResourceEdges `json:"edges"`
	group_resources *int
}

// ResourceEdges holds the relations/edges for other nodes in the graph.
type ResourceEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Group `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ResourceEdges) OwnerOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Resource) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case resource.FieldGroup, resource.FieldOthers:
			values[i] = new([]byte)
		case resource.FieldName:
			values[i] = new(sql.NullString)
		case resource.FieldID:
			values[i] = new(uuid.UUID)
		case resource.ForeignKeys[0]: // group_resources
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Resource", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Resource fields.
func (r *Resource) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case resource.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case resource.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case resource.FieldGroup:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field group", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Group); err != nil {
					return fmt.Errorf("unmarshal field group: %w", err)
				}
			}
		case resource.FieldOthers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field others", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Others); err != nil {
					return fmt.Errorf("unmarshal field others: %w", err)
				}
			}
		case resource.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_resources", value)
			} else if value.Valid {
				r.group_resources = new(int)
				*r.group_resources = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Resource entity.
func (r *Resource) QueryOwner() *GroupQuery {
	return (&ResourceClient{config: r.config}).QueryOwner(r)
}

// Update returns a builder for updating this Resource.
// Note that you need to call Resource.Unwrap() before calling this method if this Resource
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Resource) Update() *ResourceUpdateOne {
	return (&ResourceClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Resource entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Resource) Unwrap() *Resource {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Resource is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Resource) String() string {
	var builder strings.Builder
	builder.WriteString("Resource(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", name=")
	builder.WriteString(r.Name)
	builder.WriteString(", group=")
	builder.WriteString(fmt.Sprintf("%v", r.Group))
	builder.WriteString(", others=")
	builder.WriteString(fmt.Sprintf("%v", r.Others))
	builder.WriteByte(')')
	return builder.String()
}

// Resources is a parsable slice of Resource.
type Resources []*Resource

func (r Resources) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}