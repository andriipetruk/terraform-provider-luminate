// Copyright (c) 2019, Andrii Petruk. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goluminate


import (
    "context"
    "net/http"
)

type NewSiteRequest struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Settings    struct {
        ProxyAddress  string `json:"proxyAddress"`
        ProxyPort     int    `json:"proxyPort"`
        ProxyUsername string `json:"proxyUsername"`
        ProxyPassword string `json:"proxyPassword"`
    } `json:"settings"`
    Connectors             []interface{} `json:"connectors"`
    MuteHealthNotification bool          `json:"mute_health_notification"`
}



type NewSiteResponse struct {
    ID          string `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description"`
    Settings    struct {
        ProxyAddress  string `json:"proxyAddress"`
        ProxyPort     int    `json:"proxyPort"`
        ProxyUsername string `json:"proxyUsername"`
        ProxyPassword string `json:"proxyPassword"`
    } `json:"settings"`
    Connectors       []interface{} `json:"connectors"`
    ConnectorObjects []interface{} `json:"connector_objects"`
    UpdateMode       string        `json:"update_mode"`
    SiteStatus       struct {
        ConnectorsUp            []interface{} `json:"ConnectorsUp"`
        ConnectorsDown          []interface{} `json:"ConnectorsDown"`
        ConnectorsNotConfigured []interface{} `json:"ConnectorsNotConfigured"`
        Status                  string        `json:"Status"`
    } `json:"site_status"`
    DownSince              string `json:"down_since"`
    MuteHealthNotification bool   `json:"mute_health_notification"`
}

func (c *Client) CreateSite(ctx context.Context, site  NewSiteRequest) (*NewSiteResponse, *http.Response, error) { 
    
    req, err := c.NewRequest("POST", "/v2/sites/", site)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(NewSiteResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}

// https://luminatepublicapi.docs.apiary.io/#reference/sites/v2sitessiteid/update-site
func (c *Client) UpdateSite(ctx context.Context, site  NewSiteRequest, siteID string) (*NewSiteResponse, *http.Response, error) { 
    
    req, err := c.NewRequest("PUT", "/v2/sites/"+siteID, site)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(NewSiteResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}


// https://luminatepublicapi.docs.apiary.io/#reference/sites/v2sitessiteid/get-site
func (c *Client) GetSite(ctx context.Context, siteID  string) (*NewSiteRequest, *http.Response, error) { 
    
    req, err := c.NewRequest("GET", "/v2/sites/"+siteID, nil)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(NewSiteRequest)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}

// https://luminatepublicapi.docs.apiary.io/#reference/sites/v2sitessiteid/delete-site
func (c *Client) DeleteSite(ctx context.Context, siteID  string) (*http.Response, error) { 
    
    req, err := c.NewRequest("DELETE", "/v2/sites/"+siteID, nil)
    if err != nil {
        return nil, err
    }
    resp, err := c.Do(ctx, req, nil)
    if err != nil {
        return resp, err
    }
    return resp, nil
}
