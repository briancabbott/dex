// Code generated by entc, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/briancabbott/dex/storage/ent/db/authrequest"
	"github.com/briancabbott/dex/storage/ent/db/predicate"
	"github.com/briancabbott/entgo/dialect/sql"
	"github.com/briancabbott/entgo/dialect/sql/sqlgraph"
	"github.com/briancabbott/entgo/schema/field"
)

// AuthRequestUpdate is the builder for updating AuthRequest entities.
type AuthRequestUpdate struct {
	config
	hooks    []Hook
	mutation *AuthRequestMutation
}

// Where appends a list predicates to the AuthRequestUpdate builder.
func (aru *AuthRequestUpdate) Where(ps ...predicate.AuthRequest) *AuthRequestUpdate {
	aru.mutation.Where(ps...)
	return aru
}

// SetClientID sets the "client_id" field.
func (aru *AuthRequestUpdate) SetClientID(s string) *AuthRequestUpdate {
	aru.mutation.SetClientID(s)
	return aru
}

// SetScopes sets the "scopes" field.
func (aru *AuthRequestUpdate) SetScopes(s []string) *AuthRequestUpdate {
	aru.mutation.SetScopes(s)
	return aru
}

// ClearScopes clears the value of the "scopes" field.
func (aru *AuthRequestUpdate) ClearScopes() *AuthRequestUpdate {
	aru.mutation.ClearScopes()
	return aru
}

// SetResponseTypes sets the "response_types" field.
func (aru *AuthRequestUpdate) SetResponseTypes(s []string) *AuthRequestUpdate {
	aru.mutation.SetResponseTypes(s)
	return aru
}

// ClearResponseTypes clears the value of the "response_types" field.
func (aru *AuthRequestUpdate) ClearResponseTypes() *AuthRequestUpdate {
	aru.mutation.ClearResponseTypes()
	return aru
}

// SetRedirectURI sets the "redirect_uri" field.
func (aru *AuthRequestUpdate) SetRedirectURI(s string) *AuthRequestUpdate {
	aru.mutation.SetRedirectURI(s)
	return aru
}

// SetNonce sets the "nonce" field.
func (aru *AuthRequestUpdate) SetNonce(s string) *AuthRequestUpdate {
	aru.mutation.SetNonce(s)
	return aru
}

// SetState sets the "state" field.
func (aru *AuthRequestUpdate) SetState(s string) *AuthRequestUpdate {
	aru.mutation.SetState(s)
	return aru
}

// SetForceApprovalPrompt sets the "force_approval_prompt" field.
func (aru *AuthRequestUpdate) SetForceApprovalPrompt(b bool) *AuthRequestUpdate {
	aru.mutation.SetForceApprovalPrompt(b)
	return aru
}

// SetLoggedIn sets the "logged_in" field.
func (aru *AuthRequestUpdate) SetLoggedIn(b bool) *AuthRequestUpdate {
	aru.mutation.SetLoggedIn(b)
	return aru
}

// SetClaimsUserID sets the "claims_user_id" field.
func (aru *AuthRequestUpdate) SetClaimsUserID(s string) *AuthRequestUpdate {
	aru.mutation.SetClaimsUserID(s)
	return aru
}

// SetClaimsUsername sets the "claims_username" field.
func (aru *AuthRequestUpdate) SetClaimsUsername(s string) *AuthRequestUpdate {
	aru.mutation.SetClaimsUsername(s)
	return aru
}

// SetClaimsEmail sets the "claims_email" field.
func (aru *AuthRequestUpdate) SetClaimsEmail(s string) *AuthRequestUpdate {
	aru.mutation.SetClaimsEmail(s)
	return aru
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (aru *AuthRequestUpdate) SetClaimsEmailVerified(b bool) *AuthRequestUpdate {
	aru.mutation.SetClaimsEmailVerified(b)
	return aru
}

// SetClaimsGroups sets the "claims_groups" field.
func (aru *AuthRequestUpdate) SetClaimsGroups(s []string) *AuthRequestUpdate {
	aru.mutation.SetClaimsGroups(s)
	return aru
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (aru *AuthRequestUpdate) ClearClaimsGroups() *AuthRequestUpdate {
	aru.mutation.ClearClaimsGroups()
	return aru
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (aru *AuthRequestUpdate) SetClaimsPreferredUsername(s string) *AuthRequestUpdate {
	aru.mutation.SetClaimsPreferredUsername(s)
	return aru
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (aru *AuthRequestUpdate) SetNillableClaimsPreferredUsername(s *string) *AuthRequestUpdate {
	if s != nil {
		aru.SetClaimsPreferredUsername(*s)
	}
	return aru
}

// SetConnectorID sets the "connector_id" field.
func (aru *AuthRequestUpdate) SetConnectorID(s string) *AuthRequestUpdate {
	aru.mutation.SetConnectorID(s)
	return aru
}

// SetConnectorData sets the "connector_data" field.
func (aru *AuthRequestUpdate) SetConnectorData(b []byte) *AuthRequestUpdate {
	aru.mutation.SetConnectorData(b)
	return aru
}

// ClearConnectorData clears the value of the "connector_data" field.
func (aru *AuthRequestUpdate) ClearConnectorData() *AuthRequestUpdate {
	aru.mutation.ClearConnectorData()
	return aru
}

// SetExpiry sets the "expiry" field.
func (aru *AuthRequestUpdate) SetExpiry(t time.Time) *AuthRequestUpdate {
	aru.mutation.SetExpiry(t)
	return aru
}

// SetCodeChallenge sets the "code_challenge" field.
func (aru *AuthRequestUpdate) SetCodeChallenge(s string) *AuthRequestUpdate {
	aru.mutation.SetCodeChallenge(s)
	return aru
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (aru *AuthRequestUpdate) SetNillableCodeChallenge(s *string) *AuthRequestUpdate {
	if s != nil {
		aru.SetCodeChallenge(*s)
	}
	return aru
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (aru *AuthRequestUpdate) SetCodeChallengeMethod(s string) *AuthRequestUpdate {
	aru.mutation.SetCodeChallengeMethod(s)
	return aru
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (aru *AuthRequestUpdate) SetNillableCodeChallengeMethod(s *string) *AuthRequestUpdate {
	if s != nil {
		aru.SetCodeChallengeMethod(*s)
	}
	return aru
}

// Mutation returns the AuthRequestMutation object of the builder.
func (aru *AuthRequestUpdate) Mutation() *AuthRequestMutation {
	return aru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aru *AuthRequestUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aru.hooks) == 0 {
		affected, err = aru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aru.mutation = mutation
			affected, err = aru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aru.hooks) - 1; i >= 0; i-- {
			if aru.hooks[i] == nil {
				return 0, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = aru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aru *AuthRequestUpdate) SaveX(ctx context.Context) int {
	affected, err := aru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aru *AuthRequestUpdate) Exec(ctx context.Context) error {
	_, err := aru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aru *AuthRequestUpdate) ExecX(ctx context.Context) {
	if err := aru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aru *AuthRequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authrequest.Table,
			Columns: authrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authrequest.FieldID,
			},
		},
	}
	if ps := aru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aru.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClientID,
		})
	}
	if value, ok := aru.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldScopes,
		})
	}
	if aru.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldScopes,
		})
	}
	if value, ok := aru.mutation.ResponseTypes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldResponseTypes,
		})
	}
	if aru.mutation.ResponseTypesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldResponseTypes,
		})
	}
	if value, ok := aru.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldRedirectURI,
		})
	}
	if value, ok := aru.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldNonce,
		})
	}
	if value, ok := aru.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldState,
		})
	}
	if value, ok := aru.mutation.ForceApprovalPrompt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldForceApprovalPrompt,
		})
	}
	if value, ok := aru.mutation.LoggedIn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldLoggedIn,
		})
	}
	if value, ok := aru.mutation.ClaimsUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsUserID,
		})
	}
	if value, ok := aru.mutation.ClaimsUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsUsername,
		})
	}
	if value, ok := aru.mutation.ClaimsEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsEmail,
		})
	}
	if value, ok := aru.mutation.ClaimsEmailVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldClaimsEmailVerified,
		})
	}
	if value, ok := aru.mutation.ClaimsGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldClaimsGroups,
		})
	}
	if aru.mutation.ClaimsGroupsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldClaimsGroups,
		})
	}
	if value, ok := aru.mutation.ClaimsPreferredUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsPreferredUsername,
		})
	}
	if value, ok := aru.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldConnectorID,
		})
	}
	if value, ok := aru.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: authrequest.FieldConnectorData,
		})
	}
	if aru.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: authrequest.FieldConnectorData,
		})
	}
	if value, ok := aru.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authrequest.FieldExpiry,
		})
	}
	if value, ok := aru.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldCodeChallenge,
		})
	}
	if value, ok := aru.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldCodeChallengeMethod,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, aru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// AuthRequestUpdateOne is the builder for updating a single AuthRequest entity.
type AuthRequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AuthRequestMutation
}

// SetClientID sets the "client_id" field.
func (aruo *AuthRequestUpdateOne) SetClientID(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetClientID(s)
	return aruo
}

// SetScopes sets the "scopes" field.
func (aruo *AuthRequestUpdateOne) SetScopes(s []string) *AuthRequestUpdateOne {
	aruo.mutation.SetScopes(s)
	return aruo
}

// ClearScopes clears the value of the "scopes" field.
func (aruo *AuthRequestUpdateOne) ClearScopes() *AuthRequestUpdateOne {
	aruo.mutation.ClearScopes()
	return aruo
}

// SetResponseTypes sets the "response_types" field.
func (aruo *AuthRequestUpdateOne) SetResponseTypes(s []string) *AuthRequestUpdateOne {
	aruo.mutation.SetResponseTypes(s)
	return aruo
}

// ClearResponseTypes clears the value of the "response_types" field.
func (aruo *AuthRequestUpdateOne) ClearResponseTypes() *AuthRequestUpdateOne {
	aruo.mutation.ClearResponseTypes()
	return aruo
}

// SetRedirectURI sets the "redirect_uri" field.
func (aruo *AuthRequestUpdateOne) SetRedirectURI(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetRedirectURI(s)
	return aruo
}

// SetNonce sets the "nonce" field.
func (aruo *AuthRequestUpdateOne) SetNonce(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetNonce(s)
	return aruo
}

// SetState sets the "state" field.
func (aruo *AuthRequestUpdateOne) SetState(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetState(s)
	return aruo
}

// SetForceApprovalPrompt sets the "force_approval_prompt" field.
func (aruo *AuthRequestUpdateOne) SetForceApprovalPrompt(b bool) *AuthRequestUpdateOne {
	aruo.mutation.SetForceApprovalPrompt(b)
	return aruo
}

// SetLoggedIn sets the "logged_in" field.
func (aruo *AuthRequestUpdateOne) SetLoggedIn(b bool) *AuthRequestUpdateOne {
	aruo.mutation.SetLoggedIn(b)
	return aruo
}

// SetClaimsUserID sets the "claims_user_id" field.
func (aruo *AuthRequestUpdateOne) SetClaimsUserID(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsUserID(s)
	return aruo
}

// SetClaimsUsername sets the "claims_username" field.
func (aruo *AuthRequestUpdateOne) SetClaimsUsername(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsUsername(s)
	return aruo
}

// SetClaimsEmail sets the "claims_email" field.
func (aruo *AuthRequestUpdateOne) SetClaimsEmail(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsEmail(s)
	return aruo
}

// SetClaimsEmailVerified sets the "claims_email_verified" field.
func (aruo *AuthRequestUpdateOne) SetClaimsEmailVerified(b bool) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsEmailVerified(b)
	return aruo
}

// SetClaimsGroups sets the "claims_groups" field.
func (aruo *AuthRequestUpdateOne) SetClaimsGroups(s []string) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsGroups(s)
	return aruo
}

// ClearClaimsGroups clears the value of the "claims_groups" field.
func (aruo *AuthRequestUpdateOne) ClearClaimsGroups() *AuthRequestUpdateOne {
	aruo.mutation.ClearClaimsGroups()
	return aruo
}

// SetClaimsPreferredUsername sets the "claims_preferred_username" field.
func (aruo *AuthRequestUpdateOne) SetClaimsPreferredUsername(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetClaimsPreferredUsername(s)
	return aruo
}

// SetNillableClaimsPreferredUsername sets the "claims_preferred_username" field if the given value is not nil.
func (aruo *AuthRequestUpdateOne) SetNillableClaimsPreferredUsername(s *string) *AuthRequestUpdateOne {
	if s != nil {
		aruo.SetClaimsPreferredUsername(*s)
	}
	return aruo
}

// SetConnectorID sets the "connector_id" field.
func (aruo *AuthRequestUpdateOne) SetConnectorID(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetConnectorID(s)
	return aruo
}

// SetConnectorData sets the "connector_data" field.
func (aruo *AuthRequestUpdateOne) SetConnectorData(b []byte) *AuthRequestUpdateOne {
	aruo.mutation.SetConnectorData(b)
	return aruo
}

// ClearConnectorData clears the value of the "connector_data" field.
func (aruo *AuthRequestUpdateOne) ClearConnectorData() *AuthRequestUpdateOne {
	aruo.mutation.ClearConnectorData()
	return aruo
}

// SetExpiry sets the "expiry" field.
func (aruo *AuthRequestUpdateOne) SetExpiry(t time.Time) *AuthRequestUpdateOne {
	aruo.mutation.SetExpiry(t)
	return aruo
}

// SetCodeChallenge sets the "code_challenge" field.
func (aruo *AuthRequestUpdateOne) SetCodeChallenge(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetCodeChallenge(s)
	return aruo
}

// SetNillableCodeChallenge sets the "code_challenge" field if the given value is not nil.
func (aruo *AuthRequestUpdateOne) SetNillableCodeChallenge(s *string) *AuthRequestUpdateOne {
	if s != nil {
		aruo.SetCodeChallenge(*s)
	}
	return aruo
}

// SetCodeChallengeMethod sets the "code_challenge_method" field.
func (aruo *AuthRequestUpdateOne) SetCodeChallengeMethod(s string) *AuthRequestUpdateOne {
	aruo.mutation.SetCodeChallengeMethod(s)
	return aruo
}

// SetNillableCodeChallengeMethod sets the "code_challenge_method" field if the given value is not nil.
func (aruo *AuthRequestUpdateOne) SetNillableCodeChallengeMethod(s *string) *AuthRequestUpdateOne {
	if s != nil {
		aruo.SetCodeChallengeMethod(*s)
	}
	return aruo
}

// Mutation returns the AuthRequestMutation object of the builder.
func (aruo *AuthRequestUpdateOne) Mutation() *AuthRequestMutation {
	return aruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aruo *AuthRequestUpdateOne) Select(field string, fields ...string) *AuthRequestUpdateOne {
	aruo.fields = append([]string{field}, fields...)
	return aruo
}

// Save executes the query and returns the updated AuthRequest entity.
func (aruo *AuthRequestUpdateOne) Save(ctx context.Context) (*AuthRequest, error) {
	var (
		err  error
		node *AuthRequest
	)
	if len(aruo.hooks) == 0 {
		node, err = aruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aruo.mutation = mutation
			node, err = aruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aruo.hooks) - 1; i >= 0; i-- {
			if aruo.hooks[i] == nil {
				return nil, fmt.Errorf("db: uninitialized hook (forgotten import db/runtime?)")
			}
			mut = aruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aruo *AuthRequestUpdateOne) SaveX(ctx context.Context) *AuthRequest {
	node, err := aruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aruo *AuthRequestUpdateOne) Exec(ctx context.Context) error {
	_, err := aruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aruo *AuthRequestUpdateOne) ExecX(ctx context.Context) {
	if err := aruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (aruo *AuthRequestUpdateOne) sqlSave(ctx context.Context) (_node *AuthRequest, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authrequest.Table,
			Columns: authrequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: authrequest.FieldID,
			},
		},
	}
	id, ok := aruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "AuthRequest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authrequest.FieldID)
		for _, f := range fields {
			if !authrequest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != authrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aruo.mutation.ClientID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClientID,
		})
	}
	if value, ok := aruo.mutation.Scopes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldScopes,
		})
	}
	if aruo.mutation.ScopesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldScopes,
		})
	}
	if value, ok := aruo.mutation.ResponseTypes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldResponseTypes,
		})
	}
	if aruo.mutation.ResponseTypesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldResponseTypes,
		})
	}
	if value, ok := aruo.mutation.RedirectURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldRedirectURI,
		})
	}
	if value, ok := aruo.mutation.Nonce(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldNonce,
		})
	}
	if value, ok := aruo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldState,
		})
	}
	if value, ok := aruo.mutation.ForceApprovalPrompt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldForceApprovalPrompt,
		})
	}
	if value, ok := aruo.mutation.LoggedIn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldLoggedIn,
		})
	}
	if value, ok := aruo.mutation.ClaimsUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsUserID,
		})
	}
	if value, ok := aruo.mutation.ClaimsUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsUsername,
		})
	}
	if value, ok := aruo.mutation.ClaimsEmail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsEmail,
		})
	}
	if value, ok := aruo.mutation.ClaimsEmailVerified(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authrequest.FieldClaimsEmailVerified,
		})
	}
	if value, ok := aruo.mutation.ClaimsGroups(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: authrequest.FieldClaimsGroups,
		})
	}
	if aruo.mutation.ClaimsGroupsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: authrequest.FieldClaimsGroups,
		})
	}
	if value, ok := aruo.mutation.ClaimsPreferredUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldClaimsPreferredUsername,
		})
	}
	if value, ok := aruo.mutation.ConnectorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldConnectorID,
		})
	}
	if value, ok := aruo.mutation.ConnectorData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: authrequest.FieldConnectorData,
		})
	}
	if aruo.mutation.ConnectorDataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Column: authrequest.FieldConnectorData,
		})
	}
	if value, ok := aruo.mutation.Expiry(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: authrequest.FieldExpiry,
		})
	}
	if value, ok := aruo.mutation.CodeChallenge(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldCodeChallenge,
		})
	}
	if value, ok := aruo.mutation.CodeChallengeMethod(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authrequest.FieldCodeChallengeMethod,
		})
	}
	_node = &AuthRequest{config: aruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{authrequest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
