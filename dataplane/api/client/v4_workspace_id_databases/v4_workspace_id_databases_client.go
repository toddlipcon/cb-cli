// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4 workspace id databases API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4 workspace id databases API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachDatabaseToEnvironments attaches r d s resource to environemnts

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) AttachDatabaseToEnvironments(params *AttachDatabaseToEnvironmentsParams) (*AttachDatabaseToEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachDatabaseToEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachDatabaseToEnvironments",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/databases/{name}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachDatabaseToEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachDatabaseToEnvironmentsOK), nil

}

/*
CreateDatabaseInWorkspace creates r d s config in workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) CreateDatabaseInWorkspace(params *CreateDatabaseInWorkspaceParams) (*CreateDatabaseInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDatabaseInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDatabaseInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateDatabaseInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateDatabaseInWorkspaceOK), nil

}

/*
DeleteDatabaseInWorkspace deletes r d s config by name in workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteDatabaseInWorkspace(params *DeleteDatabaseInWorkspaceParams) (*DeleteDatabaseInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatabaseInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDatabaseInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/databases/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDatabaseInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDatabaseInWorkspaceOK), nil

}

/*
DeleteDatabasesInWorkspace deletes multiple r d s configs by name in workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteDatabasesInWorkspace(params *DeleteDatabasesInWorkspaceParams) (*DeleteDatabasesInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDatabasesInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDatabasesInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v4/{workspaceId}/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDatabasesInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDatabasesInWorkspaceOK), nil

}

/*
DetachDatabaseFromEnvironments detaches r d s resource from environemnts

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DetachDatabaseFromEnvironments(params *DetachDatabaseFromEnvironmentsParams) (*DetachDatabaseFromEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachDatabaseFromEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachDatabaseFromEnvironments",
		Method:             "PUT",
		PathPattern:        "/v4/{workspaceId}/databases/{name}/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachDatabaseFromEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachDatabaseFromEnvironmentsOK), nil

}

/*
GetDatabaseInWorkspace gets r d s config by name in workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetDatabaseInWorkspace(params *GetDatabaseInWorkspaceParams) (*GetDatabaseInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatabaseInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDatabaseInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/databases/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDatabaseInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDatabaseInWorkspaceOK), nil

}

/*
GetDatabaseRequestFromNameInWorkspace gets request in workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetDatabaseRequestFromNameInWorkspace(params *GetDatabaseRequestFromNameInWorkspaceParams) (*GetDatabaseRequestFromNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDatabaseRequestFromNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDatabaseRequestFromNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/databases/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDatabaseRequestFromNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDatabaseRequestFromNameInWorkspaceOK), nil

}

/*
ListDatabasesByWorkspace lists r d s configs for the given workspace

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) ListDatabasesByWorkspace(params *ListDatabasesByWorkspaceParams) (*ListDatabasesByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDatabasesByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDatabasesByWorkspace",
		Method:             "GET",
		PathPattern:        "/v4/{workspaceId}/databases",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListDatabasesByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListDatabasesByWorkspaceOK), nil

}

/*
TestDatabaseConnectionInWorkspace tests r d s connectivity

A Database Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) TestDatabaseConnectionInWorkspace(params *TestDatabaseConnectionInWorkspaceParams) (*TestDatabaseConnectionInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestDatabaseConnectionInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testDatabaseConnectionInWorkspace",
		Method:             "POST",
		PathPattern:        "/v4/{workspaceId}/databases/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestDatabaseConnectionInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestDatabaseConnectionInWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
