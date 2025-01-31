// Code generated by go-swagger; DO NOT EDIT.

package proxy_sql

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

// AddProxySQLReader is a Reader for the AddProxySQL structure.
type AddProxySQLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProxySQLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProxySQLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddProxySQLDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProxySQLOK creates a AddProxySQLOK with default headers values
func NewAddProxySQLOK() *AddProxySQLOK {
	return &AddProxySQLOK{}
}

/*AddProxySQLOK handles this case with default header values.

A successful response.
*/
type AddProxySQLOK struct {
	Payload *AddProxySQLOKBody
}

func (o *AddProxySQLOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/ProxySQL/Add][%d] addProxySqlOk  %+v", 200, o.Payload)
}

func (o *AddProxySQLOK) GetPayload() *AddProxySQLOKBody {
	return o.Payload
}

func (o *AddProxySQLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProxySQLOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProxySQLDefault creates a AddProxySQLDefault with default headers values
func NewAddProxySQLDefault(code int) *AddProxySQLDefault {
	return &AddProxySQLDefault{
		_statusCode: code,
	}
}

/*AddProxySQLDefault handles this case with default header values.

An error response.
*/
type AddProxySQLDefault struct {
	_statusCode int

	Payload *AddProxySQLDefaultBody
}

// Code gets the status code for the add proxy SQL default response
func (o *AddProxySQLDefault) Code() int {
	return o._statusCode
}

func (o *AddProxySQLDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/ProxySQL/Add][%d] AddProxySQL default  %+v", o._statusCode, o.Payload)
}

func (o *AddProxySQLDefault) GetPayload() *AddProxySQLDefaultBody {
	return o.Payload
}

func (o *AddProxySQLDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddProxySQLDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddProxySQLBody add proxy SQL body
swagger:model AddProxySQLBody
*/
type AddProxySQLBody struct {

	// add node
	AddNode *AddProxySQLParamsBodyAddNode `json:"add_node,omitempty"`

	// Node and Service access address (DNS name or IP). Required.
	Address string `json:"address,omitempty"`

	// Cluster name.
	Cluster string `json:"cluster,omitempty"`

	// custom labels
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Environment name.
	Environment string `json:"environment,omitempty"`

	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeID string `json:"node_id,omitempty"`

	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `json:"node_name,omitempty"`

	// ProxySQL password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The "pmm-agent" identifier which should run agents. Required.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service Access port. Required.
	Port int64 `json:"port,omitempty"`

	// Replication set name.
	ReplicationSet string `json:"replication_set,omitempty"`

	// Unique across all Services user-defined name. Required.
	ServiceName string `json:"service_name,omitempty"`

	// Skip connection check.
	SkipConnectionCheck bool `json:"skip_connection_check,omitempty"`

	// Use TLS for database connections.
	TLS bool `json:"tls,omitempty"`

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add proxy SQL body
func (o *AddProxySQLBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAddNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddProxySQLBody) validateAddNode(formats strfmt.Registry) error {

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
func (o *AddProxySQLBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model AddProxySQLDefaultBody
*/
type AddProxySQLDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this add proxy SQL default body
func (o *AddProxySQLDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLDefaultBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLOKBody add proxy SQL OK body
swagger:model AddProxySQLOKBody
*/
type AddProxySQLOKBody struct {

	// proxysql exporter
	ProxysqlExporter *AddProxySQLOKBodyProxysqlExporter `json:"proxysql_exporter,omitempty"`

	// service
	Service *AddProxySQLOKBodyService `json:"service,omitempty"`
}

// Validate validates this add proxy SQL OK body
func (o *AddProxySQLOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProxysqlExporter(formats); err != nil {
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

func (o *AddProxySQLOKBody) validateProxysqlExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.ProxysqlExporter) { // not required
		return nil
	}

	if o.ProxysqlExporter != nil {
		if err := o.ProxysqlExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "proxysql_exporter")
			}
			return err
		}
	}

	return nil
}

func (o *AddProxySQLOKBody) validateService(formats strfmt.Registry) error {

	if swag.IsZero(o.Service) { // not required
		return nil
	}

	if o.Service != nil {
		if err := o.Service.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addProxySqlOk" + "." + "service")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBody) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLOKBodyProxysqlExporter ProxySQLExporter runs on Generic or Container Node and exposes MySQL Service metrics.
swagger:model AddProxySQLOKBodyProxysqlExporter
*/
type AddProxySQLOKBodyProxysqlExporter struct {

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

	// Skip TLS certificate and hostname validation.
	TLSSkipVerify bool `json:"tls_skip_verify,omitempty"`

	// ProxySQL username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this add proxy SQL OK body proxysql exporter
func (o *AddProxySQLOKBodyProxysqlExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum = append(addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum, v)
	}
}

const (

	// AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	AddProxySQLOKBodyProxysqlExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// AddProxySQLOKBodyProxysqlExporterStatusSTARTING captures enum value "STARTING"
	AddProxySQLOKBodyProxysqlExporterStatusSTARTING string = "STARTING"

	// AddProxySQLOKBodyProxysqlExporterStatusRUNNING captures enum value "RUNNING"
	AddProxySQLOKBodyProxysqlExporterStatusRUNNING string = "RUNNING"

	// AddProxySQLOKBodyProxysqlExporterStatusWAITING captures enum value "WAITING"
	AddProxySQLOKBodyProxysqlExporterStatusWAITING string = "WAITING"

	// AddProxySQLOKBodyProxysqlExporterStatusSTOPPING captures enum value "STOPPING"
	AddProxySQLOKBodyProxysqlExporterStatusSTOPPING string = "STOPPING"

	// AddProxySQLOKBodyProxysqlExporterStatusDONE captures enum value "DONE"
	AddProxySQLOKBodyProxysqlExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *AddProxySQLOKBodyProxysqlExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addProxySqlOkBodyProxysqlExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLOKBodyProxysqlExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("addProxySqlOk"+"."+"proxysql_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBodyProxysqlExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBodyProxysqlExporter) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBodyProxysqlExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLOKBodyService ProxySQLService represents a generic ProxySQL instance.
swagger:model AddProxySQLOKBodyService
*/
type AddProxySQLOKBodyService struct {

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

// Validate validates this add proxy SQL OK body service
func (o *AddProxySQLOKBodyService) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddProxySQLOKBodyService) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLOKBodyService) UnmarshalBinary(b []byte) error {
	var res AddProxySQLOKBodyService
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddProxySQLParamsBodyAddNode AddNodeParams is a params to add new node to inventory while adding new service.
swagger:model AddProxySQLParamsBodyAddNode
*/
type AddProxySQLParamsBodyAddNode struct {

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

// Validate validates this add proxy SQL params body add node
func (o *AddProxySQLParamsBodyAddNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NODE_TYPE_INVALID","GENERIC_NODE","CONTAINER_NODE","REMOTE_NODE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum = append(addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEINVALID captures enum value "NODE_TYPE_INVALID"
	AddProxySQLParamsBodyAddNodeNodeTypeNODETYPEINVALID string = "NODE_TYPE_INVALID"

	// AddProxySQLParamsBodyAddNodeNodeTypeGENERICNODE captures enum value "GENERIC_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeGENERICNODE string = "GENERIC_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeCONTAINERNODE captures enum value "CONTAINER_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeCONTAINERNODE string = "CONTAINER_NODE"

	// AddProxySQLParamsBodyAddNodeNodeTypeREMOTENODE captures enum value "REMOTE_NODE"
	AddProxySQLParamsBodyAddNodeNodeTypeREMOTENODE string = "REMOTE_NODE"
)

// prop value enum
func (o *AddProxySQLParamsBodyAddNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, addProxySqlParamsBodyAddNodeTypeNodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (o *AddProxySQLParamsBodyAddNode) validateNodeType(formats strfmt.Registry) error {

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
func (o *AddProxySQLParamsBodyAddNode) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddProxySQLParamsBodyAddNode) UnmarshalBinary(b []byte) error {
	var res AddProxySQLParamsBodyAddNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
