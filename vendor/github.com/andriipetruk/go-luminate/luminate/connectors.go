// Copyright (c) 2019, Andrii Petruk. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goluminate


import (
    "context"
    "net/http"
    "time"
)

type NewConnectorRequest struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

    
type NewConnectorResponse struct {
    ID               string    `json:"id"`
    Name             string    `json:"name"`
    Version          string    `json:"version"`
    Registered       bool      `json:"registered"`
    Otp              string    `json:"otp"`
    DateCreated      time.Time `json:"date_created"`
    DateRegistered   time.Time `json:"date_registered"`
    DateOtpExpire    time.Time `json:"date_otp_expire"`
    SendLogs         bool      `json:"send_logs"`
    ConnectorStatus  string    `json:"connector_status"`
    UpdateMode       string    `json:"update_mode"`
    UpdateStatus     string    `json:"update_status"`
    UpdateStatusInfo string    `json:"update_status_info"`
}

type ConnectorCommandRequest struct {
    ConnectorName         string `json:"connector_name"`
    KerberosConfiguration struct {
        Domain     string `json:"domain"`
        KdcAddress string `json:"kdc_address"`
        KeytabPath string `json:"keytab_path"`
    } `json:"kerberos_configuration"`
}

type ConnectorCommandResponse struct {
    Linux         string `json:"linux"`
    Windows       string `json:"windows"`
    DockerCompose string `json:"docker_compose"`
    K8S           string `json:"k8s"`
}

//https://luminatepublicapi.docs.apiary.io/#reference/connectors/v2connectors/add-connector
func (c *Client) CreateConnector(ctx context.Context, connector  NewConnectorRequest, siteID string) (*NewConnectorResponse, *http.Response, error) { 
    
    req, err := c.NewRequest("POST", "/v2/connectors/?bind_to_site_id="+siteID, connector)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(NewConnectorResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}

//https://luminatepublicapi.docs.apiary.io/#reference/connectors/v2connectorsconnectoridcommand/generate-docker-installation-commands-for-this-connector.
func (c *Client) GetConnectorCommand(ctx context.Context, connectorConfig  ConnectorCommandRequest, connectorID string) (*ConnectorCommandResponse, *http.Response, error) { 
    
    req, err := c.NewRequest("POST", "/v2/connectors/"+connectorID+"/command", connectorConfig)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(ConnectorCommandResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}

