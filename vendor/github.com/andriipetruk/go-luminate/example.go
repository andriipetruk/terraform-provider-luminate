// Copyright (c) 2019, Andrii Petruk. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.



package main

import (
    "context"
    "fmt"
    "github.com/andriipetruk/go-luminate/luminate"
)

/*===========================================================================
This example demostrate how-to use go-luminate library and provide next cases:
1) create a new site
2) create a new connector for a site
3) create http type of application
4) create ssh type of application
5) create tcp type of application
==============================================================================*/


func main() {
    ctx := context.Background()
    const TenantName = "you_tenant_name"
    const ClientID = "you_client_id"
    const ClientSecret = "you_secret"
    const SiteName = "Test"
    const ConnectorName = "kubernets"


// create new client
    client := goluminate.NewClient(ctx, ClientID, ClientSecret, TenantName)



// -- New Site
    site := goluminate.NewSiteRequest{Name: SiteName}
    newSite, _, err := client.CreateSite(ctx,site)
    if err != nil {
       panic(err)
    }
    fmt.Println(newSite.ID)

// -- New Connector
    connector := goluminate.NewConnectorRequest{Name: ConnectorName, Version: "1.0"}
    newConnector, _, err := client.CreateConnector(ctx,connector,newSite.ID)
    if err != nil {
       panic(err)
    }
    fmt.Println(newConnector.ID)


// -- Install K8S 
    connectorGetCommand := goluminate.ConnectorCommandRequest{ConnectorName: ConnectorName}
    ConnectorInstall, _, err := client.GetConnectorCommand(ctx,connectorGetCommand,newConnector.ID)
    if err != nil {
       panic(err)
    }
    fmt.Println(ConnectorInstall.K8S)

// -- Create http application
    newAppHttp := goluminate.AppHttpCreateRequest{Name: "test AppB", Type: "HTTP",IsVisible: true,IsNotificationEnabled: true}
    newAppHttp.ConnectionSettings.InternalAddress="http://test.local.com"
    newAppHttp.ConnectionSettings.CustomRootPath="/"
    newAppHttp.ConnectionSettings.HealthURL="/"
    newAppHttp.ConnectionSettings.HealthMethod="Head"
    HttpApp, _, err := client.CreateApp(ctx,newAppHttp)
    if err != nil {
       panic(err)
    }
    fmt.Println(HttpApp.ID)
//  -- Bind App to Site
    client.BindAppToSite(ctx,HttpApp.ID,newSite.ID)

// -- Create ssh application
    newAppSSH := goluminate.AppSshCreateRequest{Name: "test AppC", Type: "SSH",IsVisible: true,IsNotificationEnabled: true}
    newAppSSH.ConnectionSettings.InternalAddress="tcp://test.local.com:22"
    newAppSSH.SSHSettings.UserAccounts = append(newAppSSH.SSHSettings.UserAccounts, goluminate.SshUserAccounts{Name: "root"})
    SSHApp, _, err := client.CreateApp(ctx,newAppSSH)
    if err != nil {
       panic(err)
    }
    fmt.Println(SSHApp.ID)
//  -- Bind App to Site
    client.BindAppToSite(ctx,SSHApp.ID,newSite.ID)


// -- Create tcp application
    newAppTCP := goluminate.AppTcpCreateRequest{Name: "test AppD", Type: "TCP",IsVisible: true,IsNotificationEnabled: true}
    var TcpAppPortList   []string
    TcpAppPortList  = append(TcpAppPortList, "3306")
    newAppTCP.TcpTunnelSettings = append(newAppTCP.TcpTunnelSettings, goluminate.TcpTunnelSettings{Target: "test.local.com", Ports: TcpAppPortList})
    
    TCPApp, _, err := client.CreateApp(ctx,newAppTCP)
    if err != nil {
       panic(err)
    }
    fmt.Println(TCPApp.ID)
//  -- Bind App to Site
    client.BindAppToSite(ctx,TCPApp.ID,newSite.ID)


//list all application
    applist, _, err := client.ListApp(ctx, "")

    if err != nil {
       panic(err)
    }
    fmt.Println(applist.Content[0].ID, applist.Content[0].Name)
    fmt.Println(applist.Content[3].ID, applist.Content[3].Name)

}

