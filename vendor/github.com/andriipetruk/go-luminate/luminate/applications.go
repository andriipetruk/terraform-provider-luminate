// Copyright (c) 2019, Andrii Petruk. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goluminate


import (
    "context"
    "net/http"
)

type ApplicationList struct {
    Content []struct {
        ID                    string      `json:"id"`
        CreatedOn             int64       `json:"createdOn"`
        ModifiedOn            int64       `json:"modifiedOn"`
        Name                  string      `json:"name"`
        Description           interface{} `json:"description"`
        Type                  string      `json:"type"`
        Icon                  interface{} `json:"icon"`
        IsVisible             bool        `json:"isVisible"`
        IsNotificationEnabled bool        `json:"isNotificationEnabled"`
        ConnectionSettings    struct {
            InternalAddress                  string      `json:"internalAddress"`
            InternalSSLCertificatePublicHalf interface{} `json:"internalSSLCertificatePublicHalf"`
            ExternalAddress                  string      `json:"externalAddress"`
            CustomExternalAddress            interface{} `json:"customExternalAddress"`
            LuminateAddress                  string      `json:"luminateAddress"`
            CustomSSLCertificate             interface{} `json:"customSSLCertificate"`
            AuthenticationType               string      `json:"authenticationType"`
            SsoSettings                      struct {
            } `json:"ssoSettings"`
            CustomRootPath interface{} `json:"customRootPath"`
            HealthURL      interface{} `json:"healthUrl"`
            HealthMethod   string      `json:"healthMethod"`
            Subdomain      string      `json:"subdomain"`
        } `json:"connectionSettings"`
        SiteSettings struct {
            SiteID string `json:"siteId"`
        } `json:"siteSettings"`
        LinkTranslationSettings struct {
            BaseURL                             interface{}   `json:"baseUrl"`
            IsDefaultContentRewriteRulesEnabled bool          `json:"isDefaultContentRewriteRulesEnabled"`
            IsDefaultHeaderRewriteRulesEnabled  bool          `json:"isDefaultHeaderRewriteRulesEnabled"`
            UseExternalAddressForHostAndSni     interface{}   `json:"useExternalAddressForHostAndSni"`
            CharsetEncoding                     interface{}   `json:"charsetEncoding"`
            LinkedApplications                  []interface{} `json:"linkedApplications"`
            DefaultRewriteRules                 interface{}   `json:"defaultRewriteRules"`
            CustomRewriteRules                  interface{}   `json:"customRewriteRules"`
        } `json:"linkTranslationSettings"`
        Health struct {
            ApplicationID                    string `json:"applicationId"`
            Status                           string `json:"status"`
            Cause                            string `json:"cause"`
            LastUpdatedOn                    int64  `json:"lastUpdatedOn"`
            TotalNumberOfConnectors          int    `json:"totalNumberOfConnectors"`
            TotalNumberOfAvailableConnectors int    `json:"totalNumberOfAvailableConnectors"`
            ConnectorsStatus                 []struct {
                ConnectorID string `json:"connectorId"`
                IsHealthy   bool   `json:"isHealthy"`
            } `json:"connectorsStatus"`
            LastAvailableOn int64 `json:"lastAvailableOn"`
        } `json:"health"`
        SSHSettings struct {
            UserAccounts []interface{} `json:"userAccounts"`
        } `json:"sshSettings"`
        RequestCustomizationSettings struct {
            HeaderCustomization struct {
                XForwardedFor   string `json:"X-Forwarded-For"`
                XForwardedHost  string `json:"X-Forwarded-Host"`
                XForwardedProto string `json:"X-Forwarded-Proto"`
            } `json:"headerCustomization"`
        } `json:"requestCustomizationSettings"`
        TCPTunnelSettings []struct {
            Target string `json:"target"`
            Ports  []int  `json:"ports"`
        } `json:"tcpTunnelSettings,omitempty"`
    } `json:"content"`
    TotalElements int  `json:"totalElements"`
    TotalPages    int  `json:"totalPages"`
    Last          bool `json:"last"`
    First         bool `json:"first"`
    Sort          []struct {
        Direction    string `json:"direction"`
        Property     string `json:"property"`
        IgnoreCase   bool   `json:"ignoreCase"`
        NullHandling string `json:"nullHandling"`
        Ascending    bool   `json:"ascending"`
        Descending   bool   `json:"descending"`
    } `json:"sort"`
    NumberOfElements int `json:"numberOfElements"`
    Size             int `json:"size"`
    Number           int `json:"number"`
}

type AppHttpCreateRequest struct {
    ConnectionSettings struct {
        InternalAddress string `json:"internalAddress"`
        CustomRootPath  string `json:"customRootPath"`
        HealthURL       string `json:"healthUrl"`
        HealthMethod    string `json:"healthMethod"`
    } `json:"connectionSettings"`
    Type                    string `json:"type"`
    Name                    string `json:"name"`
    IsVisible               bool   `json:"isVisible"`
    IsNotificationEnabled   bool   `json:"isNotificationEnabled"`
    LinkTranslationSettings struct {
        IsDefaultContentRewriteRulesEnabled bool          `json:"isDefaultContentRewriteRulesEnabled"`
        IsDefaultHeaderRewriteRulesEnabled  bool          `json:"isDefaultHeaderRewriteRulesEnabled"`
        UseExternalAddressForHostAndSni     bool          `json:"useExternalAddressForHostAndSni"`
        LinkedApplications                  []interface{} `json:"linkedApplications"`
    } `json:"linkTranslationSettings"`
    RequestCustomizationSettings struct {
        XForwardedFor   string `json:"X-Forwarded-For"`
        XForwardedHost  string `json:"X-Forwarded-Host"`
        XForwardedProto string `json:"X-Forwarded-Proto"`
    } `json:"requestCustomizationSettings"`
}

type AppSshCreateRequest struct {
    Type                  string `json:"type"`
    Name                  string `json:"name"`
    IsVisible             bool   `json:"isVisible"`
    IsNotificationEnabled bool   `json:"isNotificationEnabled"`
    ConnectionSettings    struct {
        InternalAddress string `json:"internalAddress"`
    } `json:"connectionSettings"`
    SSHSettings struct {
        UserAccounts []SshUserAccounts `json:"userAccounts"`
    } `json:"sshSettings"`
}

type AppTcpCreateRequest struct {
    Type                  string `json:"type"`
    Name                  string `json:"name"`
    IsVisible             bool   `json:"isVisible"`
    IsNotificationEnabled bool   `json:"isNotificationEnabled"`
    ConnectionSettings struct {
        Subdomain string `json:"subdomain"`
    } `json:"connectionSettings"`
    TcpTunnelSettings []TcpTunnelSettings `json:"tcpTunnelSettings"`
}

type SshUserAccounts struct {
     Name string `json:"name"`
}

type TcpTunnelSettings struct {
        Target string  `json:"target"`
        Ports  []string `json:"ports"`
}



type AppCreateResponse struct {
    ID                    string      `json:"id"`
    CreatedOn             int64       `json:"createdOn"`
    ModifiedOn            int64       `json:"modifiedOn"`
    Name                  string      `json:"name"`
    Description           interface{} `json:"description"`
    Type                  string      `json:"type"`
    Icon                  interface{} `json:"icon"`
    IsVisible             bool        `json:"isVisible"`
    IsNotificationEnabled bool        `json:"isNotificationEnabled"`
    ConnectionSettings    struct {
        InternalAddress                  string      `json:"internalAddress"`
        InternalSSLCertificatePublicHalf interface{} `json:"internalSSLCertificatePublicHalf"`
        ExternalAddress                  string      `json:"externalAddress"`
        CustomExternalAddress            interface{} `json:"customExternalAddress"`
        LuminateAddress                  string      `json:"luminateAddress"`
        CustomSSLCertificate             interface{} `json:"customSSLCertificate"`
        AuthenticationType               string      `json:"authenticationType"`
        SsoSettings                      struct {
        } `json:"ssoSettings"`
        CustomRootPath interface{} `json:"customRootPath"`
        HealthURL      string      `json:"healthUrl"`
        HealthMethod   string      `json:"healthMethod"`
        Subdomain      string      `json:"subdomain"`
    } `json:"connectionSettings"`
    SiteSettings struct {
        SiteID string `json:"siteId"`
    } `json:"siteSettings"`
    LinkTranslationSettings struct {
        BaseURL                             interface{} `json:"baseUrl"`
        IsDefaultContentRewriteRulesEnabled bool        `json:"isDefaultContentRewriteRulesEnabled"`
        IsDefaultHeaderRewriteRulesEnabled  bool        `json:"isDefaultHeaderRewriteRulesEnabled"`
        UseExternalAddressForHostAndSni     bool        `json:"useExternalAddressForHostAndSni"`
        CharsetEncoding                     interface{} `json:"charsetEncoding"`
        LinkedApplications                  interface{} `json:"linkedApplications"`
        DefaultRewriteRules                 interface{} `json:"defaultRewriteRules"`
        CustomRewriteRules                  interface{} `json:"customRewriteRules"`
    } `json:"linkTranslationSettings"`
    Health struct {
        ApplicationID                    string      `json:"applicationId"`
        Status                           string      `json:"status"`
        Cause                            interface{} `json:"cause"`
        LastUpdatedOn                    interface{} `json:"lastUpdatedOn"`
        TotalNumberOfConnectors          int         `json:"totalNumberOfConnectors"`
        TotalNumberOfAvailableConnectors int         `json:"totalNumberOfAvailableConnectors"`
        ConnectorsStatus                 interface{} `json:"connectorsStatus"`
        LastAvailableOn                  interface{} `json:"lastAvailableOn"`
    } `json:"health"`
    SSHSettings struct {
        UserAccounts []interface{} `json:"userAccounts"`
    } `json:"sshSettings"`
    RequestCustomizationSettings struct {
        HeaderCustomization struct {
            XForwardedHost  string `json:"X-Forwarded-Host"`
            XForwardedFor   string `json:"X-Forwarded-For"`
            XForwardedProto string `json:"X-Forwarded-Proto"`
        } `json:"headerCustomization"`
    } `json:"requestCustomizationSettings"`
}

// https://luminatepublicapi.docs.apiary.io/#reference/applications/v2applications/list-applications
func (c *Client) ListApp(ctx context.Context, opt string) (*ApplicationList, *http.Response, error) { 
    req, err := c.NewRequest("GET", "v2/applications", nil)
    if err != nil {
        return nil, nil, err
    }
    var applicationlist *ApplicationList
    resp, err := c.Do(ctx, req, &applicationlist)
    if err != nil {
        return nil, resp, err
    }
    return applicationlist, resp, nil
}


//https://luminatepublicapi.docs.apiary.io/#reference/applications/add-application
func (c *Client) CreateApp(ctx context.Context, newapp  interface{}) (*AppCreateResponse, *http.Response, error) { 

    req, err := c.NewRequest("POST", "v2/applications/", newapp)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(AppCreateResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}

// https://luminatepublicapi.docs.apiary.io/#reference/applications/v2applicationsapplicationid/update-application
func (c *Client) UpdateApp(ctx context.Context, app  interface{}, appID string) (*AppCreateResponse, *http.Response, error) { 

    req, err := c.NewRequest("PUT", "v2/applications/"+appID, app)
    if err != nil {
        return nil, nil, err
    }
    uResp := new(AppCreateResponse)
    resp, err := c.Do(ctx, req, uResp)
    if err != nil {
        return nil, resp, err
    }
    return uResp, resp, nil
}


//https://luminatepublicapi.docs.apiary.io/#reference/applications/v2applicationsapplicationidsite-bindingsiteid/assign-application-to-site
func (c *Client) BindAppToSite(ctx context.Context, appID string, siteID string) (*http.Response, error) { 
    req, err := c.NewRequest("PUT", "/v2/applications/"+appID+"/site-binding/"+siteID, "")
    if err != nil {
        return nil, err
    }
    resp, err := c.Do(ctx, req, nil)
    if err != nil {
        return  resp, err
    }
    return resp, nil
}

// https://luminatepublicapi.docs.apiary.io/#reference/applications/v2applicationsapplicationid/get-application
func (c *Client) GetApp(ctx context.Context, appID  string) (*http.Response, error) { 
    
    req, err := c.NewRequest("GET", "/v2/applications/"+appID, nil)
    if err != nil {
        return nil, err
    }
    resp, err := c.Do(ctx, req, nil)
    if err != nil {
        return resp, err
    }
    return resp, nil
}

// https://luminatepublicapi.docs.apiary.io/#reference/applications/v2applicationsapplicationid/delete-application
func (c *Client) DeleteApp(ctx context.Context, appID  string) (*http.Response, error) { 
    
    req, err := c.NewRequest("DELETE", "/v2/applications/"+appID, nil)
    if err != nil {
        return nil, err
    }
    resp, err := c.Do(ctx, req, nil)
    if err != nil {
        return resp, err
    }
    return resp, nil
}
