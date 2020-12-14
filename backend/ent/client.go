// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/KOB4k/app/ent/migrate"

	"github.com/KOB4k/app/ent/disease"
	"github.com/KOB4k/app/ent/diseasetype"
	"github.com/KOB4k/app/ent/employee"
	"github.com/KOB4k/app/ent/severity"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Disease is the client for interacting with the Disease builders.
	Disease *DiseaseClient
	// Diseasetype is the client for interacting with the Diseasetype builders.
	Diseasetype *DiseasetypeClient
	// Employee is the client for interacting with the Employee builders.
	Employee *EmployeeClient
	// Severity is the client for interacting with the Severity builders.
	Severity *SeverityClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Disease = NewDiseaseClient(c.config)
	c.Diseasetype = NewDiseasetypeClient(c.config)
	c.Employee = NewEmployeeClient(c.config)
	c.Severity = NewSeverityClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Disease:     NewDiseaseClient(cfg),
		Diseasetype: NewDiseasetypeClient(cfg),
		Employee:    NewEmployeeClient(cfg),
		Severity:    NewSeverityClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config:      cfg,
		Disease:     NewDiseaseClient(cfg),
		Diseasetype: NewDiseasetypeClient(cfg),
		Employee:    NewEmployeeClient(cfg),
		Severity:    NewSeverityClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Disease.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Disease.Use(hooks...)
	c.Diseasetype.Use(hooks...)
	c.Employee.Use(hooks...)
	c.Severity.Use(hooks...)
}

// DiseaseClient is a client for the Disease schema.
type DiseaseClient struct {
	config
}

// NewDiseaseClient returns a client for the Disease from the given config.
func NewDiseaseClient(c config) *DiseaseClient {
	return &DiseaseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `disease.Hooks(f(g(h())))`.
func (c *DiseaseClient) Use(hooks ...Hook) {
	c.hooks.Disease = append(c.hooks.Disease, hooks...)
}

// Create returns a create builder for Disease.
func (c *DiseaseClient) Create() *DiseaseCreate {
	mutation := newDiseaseMutation(c.config, OpCreate)
	return &DiseaseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Disease.
func (c *DiseaseClient) Update() *DiseaseUpdate {
	mutation := newDiseaseMutation(c.config, OpUpdate)
	return &DiseaseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiseaseClient) UpdateOne(d *Disease) *DiseaseUpdateOne {
	mutation := newDiseaseMutation(c.config, OpUpdateOne, withDisease(d))
	return &DiseaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiseaseClient) UpdateOneID(id int) *DiseaseUpdateOne {
	mutation := newDiseaseMutation(c.config, OpUpdateOne, withDiseaseID(id))
	return &DiseaseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Disease.
func (c *DiseaseClient) Delete() *DiseaseDelete {
	mutation := newDiseaseMutation(c.config, OpDelete)
	return &DiseaseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DiseaseClient) DeleteOne(d *Disease) *DiseaseDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DiseaseClient) DeleteOneID(id int) *DiseaseDeleteOne {
	builder := c.Delete().Where(disease.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiseaseDeleteOne{builder}
}

// Create returns a query builder for Disease.
func (c *DiseaseClient) Query() *DiseaseQuery {
	return &DiseaseQuery{config: c.config}
}

// Get returns a Disease entity by its id.
func (c *DiseaseClient) Get(ctx context.Context, id int) (*Disease, error) {
	return c.Query().Where(disease.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiseaseClient) GetX(ctx context.Context, id int) *Disease {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryEmployee queries the employee edge of a Disease.
func (c *DiseaseClient) QueryEmployee(d *Disease) *EmployeeQuery {
	query := &EmployeeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, id),
			sqlgraph.To(employee.Table, employee.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, disease.EmployeeTable, disease.EmployeeColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySeverity queries the severity edge of a Disease.
func (c *DiseaseClient) QuerySeverity(d *Disease) *SeverityQuery {
	query := &SeverityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, id),
			sqlgraph.To(severity.Table, severity.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, disease.SeverityTable, disease.SeverityColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDiseasetype queries the diseasetype edge of a Disease.
func (c *DiseaseClient) QueryDiseasetype(d *Disease) *DiseasetypeQuery {
	query := &DiseasetypeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(disease.Table, disease.FieldID, id),
			sqlgraph.To(diseasetype.Table, diseasetype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, disease.DiseasetypeTable, disease.DiseasetypeColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiseaseClient) Hooks() []Hook {
	return c.hooks.Disease
}

// DiseasetypeClient is a client for the Diseasetype schema.
type DiseasetypeClient struct {
	config
}

// NewDiseasetypeClient returns a client for the Diseasetype from the given config.
func NewDiseasetypeClient(c config) *DiseasetypeClient {
	return &DiseasetypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `diseasetype.Hooks(f(g(h())))`.
func (c *DiseasetypeClient) Use(hooks ...Hook) {
	c.hooks.Diseasetype = append(c.hooks.Diseasetype, hooks...)
}

// Create returns a create builder for Diseasetype.
func (c *DiseasetypeClient) Create() *DiseasetypeCreate {
	mutation := newDiseasetypeMutation(c.config, OpCreate)
	return &DiseasetypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Diseasetype.
func (c *DiseasetypeClient) Update() *DiseasetypeUpdate {
	mutation := newDiseasetypeMutation(c.config, OpUpdate)
	return &DiseasetypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DiseasetypeClient) UpdateOne(d *Diseasetype) *DiseasetypeUpdateOne {
	mutation := newDiseasetypeMutation(c.config, OpUpdateOne, withDiseasetype(d))
	return &DiseasetypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DiseasetypeClient) UpdateOneID(id int) *DiseasetypeUpdateOne {
	mutation := newDiseasetypeMutation(c.config, OpUpdateOne, withDiseasetypeID(id))
	return &DiseasetypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Diseasetype.
func (c *DiseasetypeClient) Delete() *DiseasetypeDelete {
	mutation := newDiseasetypeMutation(c.config, OpDelete)
	return &DiseasetypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DiseasetypeClient) DeleteOne(d *Diseasetype) *DiseasetypeDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DiseasetypeClient) DeleteOneID(id int) *DiseasetypeDeleteOne {
	builder := c.Delete().Where(diseasetype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DiseasetypeDeleteOne{builder}
}

// Create returns a query builder for Diseasetype.
func (c *DiseasetypeClient) Query() *DiseasetypeQuery {
	return &DiseasetypeQuery{config: c.config}
}

// Get returns a Diseasetype entity by its id.
func (c *DiseasetypeClient) Get(ctx context.Context, id int) (*Diseasetype, error) {
	return c.Query().Where(diseasetype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DiseasetypeClient) GetX(ctx context.Context, id int) *Diseasetype {
	d, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return d
}

// QueryDisease queries the disease edge of a Diseasetype.
func (c *DiseasetypeClient) QueryDisease(d *Diseasetype) *DiseaseQuery {
	query := &DiseaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(diseasetype.Table, diseasetype.FieldID, id),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, diseasetype.DiseaseTable, diseasetype.DiseaseColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DiseasetypeClient) Hooks() []Hook {
	return c.hooks.Diseasetype
}

// EmployeeClient is a client for the Employee schema.
type EmployeeClient struct {
	config
}

// NewEmployeeClient returns a client for the Employee from the given config.
func NewEmployeeClient(c config) *EmployeeClient {
	return &EmployeeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `employee.Hooks(f(g(h())))`.
func (c *EmployeeClient) Use(hooks ...Hook) {
	c.hooks.Employee = append(c.hooks.Employee, hooks...)
}

// Create returns a create builder for Employee.
func (c *EmployeeClient) Create() *EmployeeCreate {
	mutation := newEmployeeMutation(c.config, OpCreate)
	return &EmployeeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Employee.
func (c *EmployeeClient) Update() *EmployeeUpdate {
	mutation := newEmployeeMutation(c.config, OpUpdate)
	return &EmployeeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EmployeeClient) UpdateOne(e *Employee) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployee(e))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EmployeeClient) UpdateOneID(id int) *EmployeeUpdateOne {
	mutation := newEmployeeMutation(c.config, OpUpdateOne, withEmployeeID(id))
	return &EmployeeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Employee.
func (c *EmployeeClient) Delete() *EmployeeDelete {
	mutation := newEmployeeMutation(c.config, OpDelete)
	return &EmployeeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EmployeeClient) DeleteOne(e *Employee) *EmployeeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EmployeeClient) DeleteOneID(id int) *EmployeeDeleteOne {
	builder := c.Delete().Where(employee.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EmployeeDeleteOne{builder}
}

// Create returns a query builder for Employee.
func (c *EmployeeClient) Query() *EmployeeQuery {
	return &EmployeeQuery{config: c.config}
}

// Get returns a Employee entity by its id.
func (c *EmployeeClient) Get(ctx context.Context, id int) (*Employee, error) {
	return c.Query().Where(employee.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EmployeeClient) GetX(ctx context.Context, id int) *Employee {
	e, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return e
}

// QueryDisease queries the disease edge of a Employee.
func (c *EmployeeClient) QueryDisease(e *Employee) *DiseaseQuery {
	query := &DiseaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(employee.Table, employee.FieldID, id),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, employee.DiseaseTable, employee.DiseaseColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EmployeeClient) Hooks() []Hook {
	return c.hooks.Employee
}

// SeverityClient is a client for the Severity schema.
type SeverityClient struct {
	config
}

// NewSeverityClient returns a client for the Severity from the given config.
func NewSeverityClient(c config) *SeverityClient {
	return &SeverityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `severity.Hooks(f(g(h())))`.
func (c *SeverityClient) Use(hooks ...Hook) {
	c.hooks.Severity = append(c.hooks.Severity, hooks...)
}

// Create returns a create builder for Severity.
func (c *SeverityClient) Create() *SeverityCreate {
	mutation := newSeverityMutation(c.config, OpCreate)
	return &SeverityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Update returns an update builder for Severity.
func (c *SeverityClient) Update() *SeverityUpdate {
	mutation := newSeverityMutation(c.config, OpUpdate)
	return &SeverityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeverityClient) UpdateOne(s *Severity) *SeverityUpdateOne {
	mutation := newSeverityMutation(c.config, OpUpdateOne, withSeverity(s))
	return &SeverityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeverityClient) UpdateOneID(id int) *SeverityUpdateOne {
	mutation := newSeverityMutation(c.config, OpUpdateOne, withSeverityID(id))
	return &SeverityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Severity.
func (c *SeverityClient) Delete() *SeverityDelete {
	mutation := newSeverityMutation(c.config, OpDelete)
	return &SeverityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SeverityClient) DeleteOne(s *Severity) *SeverityDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SeverityClient) DeleteOneID(id int) *SeverityDeleteOne {
	builder := c.Delete().Where(severity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeverityDeleteOne{builder}
}

// Create returns a query builder for Severity.
func (c *SeverityClient) Query() *SeverityQuery {
	return &SeverityQuery{config: c.config}
}

// Get returns a Severity entity by its id.
func (c *SeverityClient) Get(ctx context.Context, id int) (*Severity, error) {
	return c.Query().Where(severity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeverityClient) GetX(ctx context.Context, id int) *Severity {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}

// QueryDisease queries the disease edge of a Severity.
func (c *SeverityClient) QueryDisease(s *Severity) *DiseaseQuery {
	query := &DiseaseQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(severity.Table, severity.FieldID, id),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, severity.DiseaseTable, severity.DiseaseColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeverityClient) Hooks() []Hook {
	return c.hooks.Severity
}