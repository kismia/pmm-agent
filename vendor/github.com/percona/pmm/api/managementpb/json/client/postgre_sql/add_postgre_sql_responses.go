// Code generated by go-swagger; DO NOT EDIT.

package postgre_sql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// AddPostgreSQLReader is a Reader for the AddPostgreSQL structure.
type AddPostgreSQLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPostgreSQLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddPostgreSQLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddPostgreSQLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPostgreSQLOK creates a AddPostgreSQLOK with default headers values
func NewAddPostgreSQLOK() *AddPostgreSQLOK {
	return &AddPostgreSQLOK{}
}

/*AddPostgreSQLOK handles this case with default header values.

A successful response.
*/
type AddPostgreSQLOK struct {
	Payload *AddPostgreSQLOKBody
}

func (o *AddPostgreSQLOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/PostgreSQL/Add][%d] addPostgreSqlOk  %+v", 200, o.Payload)
}

func (o *AddPostgreSQLOK) GetPayload() *AddPostgreSQLOKBody {
	return o.Payload
}

func (o *AddPostgreSQLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPostgreSQLDefault creates a AddPostgreSQLDefault with default headers values
func NewAddPostgreSQLDefault(code int) *AddPostgreSQLDefault {
	return &AddPostgreSQLDefault{
		_statusCode: code,
	}
}

/*AddPostgreSQLDefault handles this case with default header values.

An error response.
*/
type AddPostgreSQLDefault struct {
	_statusCode int

	Payload *AddPostgreSQLDefaultBody
}

// Code gets the status code for the add postgre SQL default response
func (o *AddPostgreSQLDefault) Code() int {
	return o._statusCode
}

func (o *AddPostgreSQLDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/PostgreSQL/Add][%d] AddPostgreSQL default  %+v", o._statusCode, o.Payload)
}

func (o *AddPostgreSQLDefault) GetPayload() *AddPostgreSQLDefaultBody {
	return o.Payload
}

func (o *AddPostgreSQLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddPostgreSQLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddPostgreSQLBody add postgre SQL body
swagger:model AddPostgreSQLBody
*/
type AddPostgreSQLBody struct {

	// add node
	AddNode *AddPostgreSQLParamsBodyAddNode `json:"add_node,omitempty"`

	// Node and Service access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// PostgreSQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service Access port. Required.
	Port int64 `json:"port,omitempty"`

	// If true, adds qan-postgresql-pgstatements-agent for provided service.
	QANPostgresqlPgstatementsAgent bool `json:"qan_postgresql_pgstatements_agent,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add postgre SQL body
func (o *AddPostgreSQLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLBody) validateAddNode(formats strfmt.Registry) error {

	if swag.IsZero(o.AddNode) { // not required
		return nil
	}

	if o.AddNode != nil {
		if err := o.AddNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "add_node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddPostgreSQLDefaultBody
*/
type AddPostgreSQLDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add postgre SQL default body
func (o *AddPostgreSQLDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBody add postgre SQL OK body
swagger:model AddPostgreSQLOKBody
*/
type AddPostgreSQLOKBody struct {

	// postgres exporter
	PostgresExporter *AddPostgreSQLOKBodyPostgresExporter `json:"postgres_exporter,omitempty"`

	// qan postgresql pgstatements agent
	QANPostgresqlPgstatementsAgent *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent `json:"qan_postgresql_pgstatements_agent,omitempty"`

	// service
	Service *AddPostgreSQLOKBodyService `json:"service,omitempty"`
}

// Validate validates this add postgre SQL OK body
func (o *AddPostgreSQLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePostgresExporter(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateQANPostgresqlPgstatementsAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddPostgreSQLOKBody) validatePostgresExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.PostgresExporter) { // not required
		return nil
	}

	if o.PostgresExporter != nil {
		if err := o.PostgresExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "postgres_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddPostgreSQLOKBody) validateQANPostgresqlPgstatementsAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.QANPostgresqlPgstatementsAgent) { // not required
		return nil
	}

	if o.QANPostgresqlPgstatementsAgent != nil {
		if err := o.QANPostgresqlPgstatementsAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "qan_postgresql_pgstatements_agent")
			}
			return err
		}
	}

	return nil
}

func (o *AddPostgreSQLOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addPostgreSqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBody) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyPostgresExporter PostgresExporter runs on Generic or Container Node and exposes PostgreSQL Service metrics.
swagger:model AddPostgreSQLOKBodyPostgresExporter
*/
type AddPostgreSQLOKBodyPostgresExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// PostgreSQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add postgre SQL OK body postgres exporter
func (o *AddPostgreSQLOKBodyPostgresExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum = append(addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgreSQLOKBodyPostgresExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgreSQLOKBodyPostgresExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgreSQLOKBodyPostgresExporterStatusSTARTING captures enum value "STARTING"
	AddPostgreSQLOKBodyPostgresExporterStatusSTARTING string = "STARTING"

	// AddPostgreSQLOKBodyPostgresExporterStatusRUNNING captures enum value "RUNNING"
	AddPostgreSQLOKBodyPostgresExporterStatusRUNNING string = "RUNNING"

	// AddPostgreSQLOKBodyPostgresExporterStatusWAITING captures enum value "WAITING"
	AddPostgreSQLOKBodyPostgresExporterStatusWAITING string = "WAITING"

	// AddPostgreSQLOKBodyPostgresExporterStatusSTOPPING captures enum value "STOPPING"
	AddPostgreSQLOKBodyPostgresExporterStatusSTOPPING string = "STOPPING"

	// AddPostgreSQLOKBodyPostgresExporterStatusDONE captures enum value "DONE"
	AddPostgreSQLOKBodyPostgresExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddPostgreSQLOKBodyPostgresExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addPostgreSqlOkBodyPostgresExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLOKBodyPostgresExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgreSqlOk"+"."+"postgres_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyPostgresExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyPostgresExporter) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyPostgresExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent QANPostgreSQLPgStatementsAgent runs within pmm-agent and sends PostgreSQL Query Analytics data to the PMM Server.
swagger:model AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent
*/
type AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	//
	//  - STARTING: Agent is starting.
	//  - RUNNING: Agent is running.
	//  - WAITING: Agent encountered error and will be restarted automatically soon.
	//  - STOPPING: Agent is stopping.
	//  - DONE: Agent finished.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// PostgreSQL username for getting pg stat statements data.
	Username string `json:"username,omitempty"`
}

// Validate validates this add postgre SQL OK body QAN postgresql pgstatements agent
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum = append(addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum, v)
	}
}

const (

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING captures enum value "STARTING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTARTING string = "STARTING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING captures enum value "RUNNING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusRUNNING string = "RUNNING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusWAITING captures enum value "WAITING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusWAITING string = "WAITING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING captures enum value "STOPPING"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusSTOPPING string = "STOPPING"

	// AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusDONE captures enum value "DONE"
	AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgentStatusDONE string = "DONE"
)

// prop value enum
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addPostgreSqlOkBodyQanPostgresqlPgstatementsAgentTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addPostgreSqlOk"+"."+"qan_postgresql_pgstatements_agent"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyQANPostgresqlPgstatementsAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLOKBodyService PostgreSQLService represents a generic PostgreSQL instance.
swagger:model AddPostgreSQLOKBodyService
*/
type AddPostgreSQLOKBodyService struct {

	// Access address (DNS name or IP).
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier where this instance runs.
	NodeID string `json:"node_id,omitempty"`

	// Access port.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique randomly generated instance identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Unique across all Services user-defined name.
	ServiceName string `json:"service_name,omitempty"`
}

// Validate validates this add postgre SQL OK body service
func (o *AddPostgreSQLOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddPostgreSQLParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddPostgreSQLParamsBodyAddNode
*/
type AddPostgreSQLParamsBodyAddNode struct {

	// Node availability zone.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	ContainerID string `json:"container_id,omitempty"`

	// Container name.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution name and version.
	Distro string `json:"distro,omitempty"`

	// Linux machine-id.
	MachineID string `json:"machine_id,omitempty"`

	// Node model.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`

	// NodeType describes supported Node types.
	// Enum: [NODE_TYPE_INVALID GENERIC_NODE CONTAINER_NODE REMOTE_NODE]
	NodeType *string `json:"node_type,omitempty"`

	// Node region.
	Region string `json:"region,omitempty"`
}

// Validate validates this add postgre SQL params body add node
func (o *AddPostgreSQLParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum = append(addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddPostgreSQLParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddPostgreSQLParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddPostgreSQLParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"
)

// prop value enum
func (o *AddPostgreSQLParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addPostgreSqlParamsBodyAddNodeTypeNodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddPostgreSQLParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(o.NodeType) { // not required
		return nil
	}

	// value enum
	if err := o.validateNodeTypeEnum("body"+"."+"add_node"+"."+"node_type", "body", *o.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddPostgreSQLParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddPostgreSQLParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddPostgreSQLParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
