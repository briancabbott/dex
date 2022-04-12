// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/briancabbott/dex/storage/ent/db/authcode"
	"github.com/briancabbott/dex/storage/ent/db/predicate"
)

// AuthCodeUpdate is the builder for updating AuthCode entities.
type AuthCodeUpdate struct {
	config
	hooks    []Hook
	mutation *AuthCodeMutation
}

// Where appends a list predicates to the AuthCodeUpdate builder.
func (acu *AuthCodeUpdate) Where(ps ...predicate.AuthCode) *AuthCodeUpdate {
	acu.mutation.Where(ps...)
	return acu
}

// SetClientID sets the "client_id" field.
func (acu *AuthCodeUpdate) SetClientID(s string) *AuthCodeUpdate {
	acu.mutation.SetClientID(s)
	return acu
}

// SetScopes sets the "scopes" field.
func (acu *AuthCodeUpdate) SetScopes(s []string) *AuthCodeUpdate {
	acu.mutation.SetScopes(s)
	return acu
}

// ClearScopes clears the value of the "scopes" field.
func (acu *AuthCodeUpdate) ClearScopes() *AuthCodeUpdate {
	acu.mutation.ClearScopes()
	return acu
}

// SetNonce sets the "nonce" field.
func (acu *AuthCodeUpdate) SetNonce(s string) *AuthCodeUpdate {
	acu.mutation.SetNonce(s)
	return acu
}

// SetRedirectURI sets the "redirect_uri" field.
func (acu *AuthCodeUpdate) SetRedirectURI(s string) *AuthCodeUpdate {
	acu.mutation.SetRedirectURI(s)
	return acu
}

// SetClaimsUserID sets the "claims_user_id" field.
func (acu *AuthCodeUpdate) SetClaimsUserID(s string) *AuthCodeUpdate {
	acu.mutation.SetClaimsUserID(s)
	return acu
}

// SetClaimsUsername sets the "claims_username" field.
func (acu *AuthCodeUpdate) SetClaimsUsername(s string) *AuthCodeUpdate {
	acu.mutation.SetClaimsUsername(s)
	return acu
}

// SetClaimsEmail sets the "claims_email" field.
func (acu *AuthCodeUpdate) SetClaimsEmail(s string) *AuthCodeUpdate {
	acu.mutation.SetClaimsEmail(s)
	return acu
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (acu *AuthCodeUpdate) SetClaimsEmailVerified(b bool) *AuthCodeUpdate {
	acu.mutation.SetClaimsEmailVerified(b)
	return acu
}

// SetClaimsGroups sets the "claims_groups" field.
func (acu *AuthCodeUpdate) SetClaimsGroups(s []string) *AuthCodeUpdate {
	acu.mutation.SetClaimsGroups(s)
	return acu
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (acu *AuthCodeUpdate) ClearClaimsGroups() *AuthCodeUpdate {
	acu.mutation.ClearClaimsGroups()
	return acu
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (acu *AuthCodeUpdate) SetClaimsPreferredUsername(s string) *AuthCodeUpdate {
	acu.mutation.SetClaimsPreferredUsername(s)
	return acu
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableClaimsPreferredUsername(s *string) *AuthCodeUpdate {
	if s != nil {
		acu.SetClaimsPreferredUsername(*s)
	}
	return acu
}

// SetConnectorID sets the "connector_id" field.
func (acu *AuthCodeUpdate) SetConnectorID(s string) *AuthCodeUpdate {
	acu.mutation.SetConnectorID(s)
	return acu
}

// SetConnectorData sets the "connector_data" field.
func (acu *AuthCodeUpdate) SetConnectorData(b []byte) *AuthCodeUpdate {
	acu.mutation.SetConnectorData(b)
	return acu
}

// ClearConnectorData clears the value of the "connector_data" field.
func (acu *AuthCodeUpdate) ClearConnectorData() *AuthCodeUpdate {
	acu.mutation.ClearConnectorData()
	return acu
}

// SetExpiry sets the "expiry" field.
func (acu *AuthCodeUpdate) SetExpiry(t time.Time) *AuthCodeUpdate {
	acu.mutation.SetExpiry(t)
	return acu
}

// SetCodeChallenge sets the "code_challenge" field.
func (acu *AuthCodeUpdate) SetCodeChallenge(s string) *AuthCodeUpdate {
	acu.mutation.SetCodeChallenge(s)
	return acu
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableCodeChallenge(s *string) *AuthCodeUpdate {
	if s != nil {
		acu.SetCodeChallenge(*s)
	}
	return acu
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (acu *AuthCodeUpdate) SetCodeChallengeMethod(s string) *AuthCodeUpdate {
	acu.mutation.SetCodeChallengeMethod(s)
	return acu
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (acu *AuthCodeUpdate) SetNillableCodeChallengeMethod(s *string) *AuthCodeUpdate {
	if s != nil {
		acu.SetCodeChallengeMethod(*s)
	}
	return acu
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acu *AuthCodeUpdate) Mutation() *AuthCodeMutation {
	return acu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (acu *AuthCodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(acu.hooks) == 0 {
		if err = acu.check(); err != nil {
			return 0, err
		}
		affected, err = acu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acu.check(); err != nil {
				return 0, err
			}
			acu.mutation = mutation
			affected, err = acu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(acu.hooks) - 1; i >= 0; i-- {
			if acu.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = acu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (acu *AuthCodeUpdate) SaveX(ctx context.Context) int {
	affected, err := acu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (acu *AuthCodeUpdate) Exec(ctx context.Context) error {
	_, err := acu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acu *AuthCodeUpdate) ExecX(ctx context.Context) {
	if err := acu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acu *AuthCodeUpdate) check() error {
	if v, ok := acu.mutation.ClientID(); ok {
		if err := authcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.client_id": %w`, err)}
		}
	}
	if v, ok := acu.mutation.Nonce(); ok {
		if err := authcode.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`db: validator failed for field "AuthCode.nonce": %w`, err)}
		}
	}
	if v, ok := acu.mutation.RedirectURI(); ok {
		if err := authcode.RedirectURIValidator(v); err != nil {
			return &ValidationError{Name: "redirect_uri", err: fmt.Errorf(`db: validator failed for field "AuthCode.redirect_uri": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ClaimsUserID(); ok {
		if err := authcode.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_user_id": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ClaimsUsername(); ok {
		if err := authcode.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_username": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ClaimsEmail(); ok {
		if err := authcode.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_email": %w`, err)}
		}
	}
	if v, ok := acu.mutation.ConnectorID(); ok {
		if err := authcode.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.connector_id": %w`, err)}
		}
	}
	return nil
}

func (acu *AuthCodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authcode.Table,
			Columns: authcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authcode.FieldID,
			},
		},
	}
	if ps := acu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acu.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClientID,
		})
	}
	if value, ok := acu.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldScopes,
		})
	}
	if acu.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldScopes,
		})
	}
	if value, ok := acu.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldNonce,
		})
	}
	if value, ok := acu.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldRedirectURI,
		})
	}
	if value, ok := acu.mutation.ClaimsUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsUserID,
		})
	}
	if value, ok := acu.mutation.ClaimsUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsUsername,
		})
	}
	if value, ok := acu.mutation.ClaimsEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsEmail,
		})
	}
	if value, ok := acu.mutation.ClaimsEmailVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authcode.FieldClaimsEmailVerified,
		})
	}
	if value, ok := acu.mutation.ClaimsGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldClaimsGroups,
		})
	}
	if acu.mutation.ClaimsGroupsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldClaimsGroups,
		})
	}
	if value, ok := acu.mutation.ClaimsPreferredUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsPreferredUsername,
		})
	}
	if value, ok := acu.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldConnectorID,
		})
	}
	if value, ok := acu.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: authcode.FieldConnectorData,
		})
	}
	if acu.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: authcode.FieldConnectorData,
		})
	}
	if value, ok := acu.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldExpiry,
		})
	}
	if value, ok := acu.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallenge,
		})
	}
	if value, ok := acu.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallengeMethod,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, acu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AuthCodeUpdateOne is the builder for updating a single AuthCode entity.
type AuthCodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthCodeMutation
}

// SetClientID sets the "client_id" field.
func (acuo *AuthCodeUpdateOne) SetClientID(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClientID(s)
	return acuo
}

// SetScopes sets the "scopes" field.
func (acuo *AuthCodeUpdateOne) SetScopes(s []string) *AuthCodeUpdateOne {
	acuo.mutation.SetScopes(s)
	return acuo
}

// ClearScopes clears the value of the "scopes" field.
func (acuo *AuthCodeUpdateOne) ClearScopes() *AuthCodeUpdateOne {
	acuo.mutation.ClearScopes()
	return acuo
}

// SetNonce sets the "nonce" field.
func (acuo *AuthCodeUpdateOne) SetNonce(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetNonce(s)
	return acuo
}

// SetRedirectURI sets the "redirect_uri" field.
func (acuo *AuthCodeUpdateOne) SetRedirectURI(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetRedirectURI(s)
	return acuo
}

// SetClaimsUserID sets the "claims_user_id" field.
func (acuo *AuthCodeUpdateOne) SetClaimsUserID(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsUserID(s)
	return acuo
}

// SetClaimsUsername sets the "claims_username" field.
func (acuo *AuthCodeUpdateOne) SetClaimsUsername(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsUsername(s)
	return acuo
}

// SetClaimsEmail sets the "claims_email" field.
func (acuo *AuthCodeUpdateOne) SetClaimsEmail(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsEmail(s)
	return acuo
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (acuo *AuthCodeUpdateOne) SetClaimsEmailVerified(b bool) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsEmailVerified(b)
	return acuo
}

// SetClaimsGroups sets the "claims_groups" field.
func (acuo *AuthCodeUpdateOne) SetClaimsGroups(s []string) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsGroups(s)
	return acuo
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (acuo *AuthCodeUpdateOne) ClearClaimsGroups() *AuthCodeUpdateOne {
	acuo.mutation.ClearClaimsGroups()
	return acuo
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (acuo *AuthCodeUpdateOne) SetClaimsPreferredUsername(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetClaimsPreferredUsername(s)
	return acuo
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableClaimsPreferredUsername(s *string) *AuthCodeUpdateOne {
	if s != nil {
		acuo.SetClaimsPreferredUsername(*s)
	}
	return acuo
}

// SetConnectorID sets the "connector_id" field.
func (acuo *AuthCodeUpdateOne) SetConnectorID(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetConnectorID(s)
	return acuo
}

// SetConnectorData sets the "connector_data" field.
func (acuo *AuthCodeUpdateOne) SetConnectorData(b []byte) *AuthCodeUpdateOne {
	acuo.mutation.SetConnectorData(b)
	return acuo
}

// ClearConnectorData clears the value of the "connector_data" field.
func (acuo *AuthCodeUpdateOne) ClearConnectorData() *AuthCodeUpdateOne {
	acuo.mutation.ClearConnectorData()
	return acuo
}

// SetExpiry sets the "expiry" field.
func (acuo *AuthCodeUpdateOne) SetExpiry(t time.Time) *AuthCodeUpdateOne {
	acuo.mutation.SetExpiry(t)
	return acuo
}

// SetCodeChallenge sets the "code_challenge" field.
func (acuo *AuthCodeUpdateOne) SetCodeChallenge(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetCodeChallenge(s)
	return acuo
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableCodeChallenge(s *string) *AuthCodeUpdateOne {
	if s != nil {
		acuo.SetCodeChallenge(*s)
	}
	return acuo
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (acuo *AuthCodeUpdateOne) SetCodeChallengeMethod(s string) *AuthCodeUpdateOne {
	acuo.mutation.SetCodeChallengeMethod(s)
	return acuo
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (acuo *AuthCodeUpdateOne) SetNillableCodeChallengeMethod(s *string) *AuthCodeUpdateOne {
	if s != nil {
		acuo.SetCodeChallengeMethod(*s)
	}
	return acuo
}

// Mutation returns the AuthCodeMutation object of the builder.
func (acuo *AuthCodeUpdateOne) Mutation() *AuthCodeMutation {
	return acuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (acuo *AuthCodeUpdateOne) Select(field string, fields ...string) *AuthCodeUpdateOne {
	acuo.fields = append([]string{field}, fields...)
	return acuo
}

// Save executes the query and returns the updated AuthCode entity.
func (acuo *AuthCodeUpdateOne) Save(ctx context.Context) (*AuthCode, error) {
	var (
		err  error
		node *AuthCode
	)
	if len(acuo.hooks) == 0 {
		if err = acuo.check(); err != nil {
			return nil, err
		}
		node, err = acuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthCodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acuo.check(); err != nil {
				return nil, err
			}
			acuo.mutation = mutation
			node, err = acuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(acuo.hooks) - 1; i >= 0; i-- {
			if acuo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = acuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) SaveX(ctx context.Context) *AuthCode {
	node, err := acuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (acuo *AuthCodeUpdateOne) Exec(ctx context.Context) error {
	_, err := acuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acuo *AuthCodeUpdateOne) ExecX(ctx context.Context) {
	if err := acuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acuo *AuthCodeUpdateOne) check() error {
	if v, ok := acuo.mutation.ClientID(); ok {
		if err := authcode.ClientIDValidator(v); err != nil {
			return &ValidationError{Name: "client_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.client_id": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.Nonce(); ok {
		if err := authcode.NonceValidator(v); err != nil {
			return &ValidationError{Name: "nonce", err: fmt.Errorf(`db: validator failed for field "AuthCode.nonce": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.RedirectURI(); ok {
		if err := authcode.RedirectURIValidator(v); err != nil {
			return &ValidationError{Name: "redirect_uri", err: fmt.Errorf(`db: validator failed for field "AuthCode.redirect_uri": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ClaimsUserID(); ok {
		if err := authcode.ClaimsUserIDValidator(v); err != nil {
			return &ValidationError{Name: "claims_user_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_user_id": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ClaimsUsername(); ok {
		if err := authcode.ClaimsUsernameValidator(v); err != nil {
			return &ValidationError{Name: "claims_username", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_username": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ClaimsEmail(); ok {
		if err := authcode.ClaimsEmailValidator(v); err != nil {
			return &ValidationError{Name: "claims_email", err: fmt.Errorf(`db: validator failed for field "AuthCode.claims_email": %w`, err)}
		}
	}
	if v, ok := acuo.mutation.ConnectorID(); ok {
		if err := authcode.ConnectorIDValidator(v); err != nil {
			return &ValidationError{Name: "connector_id", err: fmt.Errorf(`db: validator failed for field "AuthCode.connector_id": %w`, err)}
		}
	}
	return nil
}

func (acuo *AuthCodeUpdateOne) sqlSave(ctx context.Context) (_node *AuthCode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authcode.Table,
			Columns: authcode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authcode.FieldID,
			},
		},
	}
	id, ok := acuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AuthCode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := acuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authcode.FieldID)
		for _, f := range fields {
			if !authcode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != authcode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := acuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := acuo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClientID,
		})
	}
	if value, ok := acuo.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldScopes,
		})
	}
	if acuo.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldScopes,
		})
	}
	if value, ok := acuo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldNonce,
		})
	}
	if value, ok := acuo.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldRedirectURI,
		})
	}
	if value, ok := acuo.mutation.ClaimsUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsUserID,
		})
	}
	if value, ok := acuo.mutation.ClaimsUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsUsername,
		})
	}
	if value, ok := acuo.mutation.ClaimsEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsEmail,
		})
	}
	if value, ok := acuo.mutation.ClaimsEmailVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authcode.FieldClaimsEmailVerified,
		})
	}
	if value, ok := acuo.mutation.ClaimsGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authcode.FieldClaimsGroups,
		})
	}
	if acuo.mutation.ClaimsGroupsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authcode.FieldClaimsGroups,
		})
	}
	if value, ok := acuo.mutation.ClaimsPreferredUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldClaimsPreferredUsername,
		})
	}
	if value, ok := acuo.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldConnectorID,
		})
	}
	if value, ok := acuo.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: authcode.FieldConnectorData,
		})
	}
	if acuo.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: authcode.FieldConnectorData,
		})
	}
	if value, ok := acuo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authcode.FieldExpiry,
		})
	}
	if value, ok := acuo.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallenge,
		})
	}
	if value, ok := acuo.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authcode.FieldCodeChallengeMethod,
		})
	}
	_node = &AuthCode{config: acuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, acuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authcode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
