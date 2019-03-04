// Copyright (c) 2019, Andrii Petruk. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goluminate


import (
    "context"
    "log"
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "io/ioutil"
    "net/url"
    "net/http"
    "net/http/httputil"
    "strings"
    "golang.org/x/oauth2/clientcredentials"
)

const userAgent = "go-luminate"
const debugRun = false
const debugRequest = false



type Client struct {
    // The base-URL that should be used should be api.<tenant-name>.luminatesec.com.
    BaseURL   *url.URL 
    UserAgent string // User agent used when communicating with the luminate API.
    client *http.Client // HTTP client used to communicate with the API.
        }

// https://luminatepublicapi.docs.apiary.io/#introduction/authorization
// NewClient returns a new luminate API client with provided base URL.
func NewClient(ctx context.Context, ClientID, ClientSecret, TenantName string) *Client {
    conf := &clientcredentials.Config{
        ClientID:       ClientID,
        ClientSecret:   ClientSecret,
        TokenURL:       "https://api."+TenantName+".luminatesec.com/v1/oauth/token",
        Scopes:         []string{"luminate-scope"},
        EndpointParams: nil,
    }

    OAuth2Client := conf.Client(ctx)

    baseEndpoint, err := url.Parse("https://api."+TenantName+".luminatesec.com/")
    if err != nil {
        return nil
    }

    c := &Client{client: OAuth2Client, BaseURL: baseEndpoint, UserAgent: userAgent}
    c.BaseURL = baseEndpoint
    return c

}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
    if !strings.HasSuffix(c.BaseURL.Path, "/") {
        return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
    }
    u, err := c.BaseURL.Parse(urlStr)
    if err != nil {
        return nil, err
    }


    // Debug
    if debugRun  {
        log.Printf("[DEBUG] request to url: %v", u.String())
        log.Printf("[DEBUG] request method: %v", method)        
    }

    var buf io.ReadWriter
    if body != nil {
        buf = new(bytes.Buffer)
        enc := json.NewEncoder(buf)
        enc.SetEscapeHTML(false)
        err := enc.Encode(body)
        if err != nil {
            return nil, err
        }
    }

    req, err := http.NewRequest(method, u.String(), buf)
    if err != nil {
        return nil, err
    }

    if body != nil {
        req.Header.Set("Content-Type", "application/json")
    }

    if c.UserAgent != "" {
        req.Header.Set("User-Agent", c.UserAgent)
    }
    return req, nil
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by target, or returned as an
// error if an API error has occurred. If target implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it. 
//
// The provided ctx must be non-nil. If it is canceled or times out,
// ctx.Err() will be returned.
func (c *Client) Do(ctx context.Context, req *http.Request, target  interface{})  (*http.Response, error)  {
    req = req.WithContext(ctx)
    if debugRequest  {
       dump, err := httputil.DumpRequestOut(req, true)
       if err != nil {
          log.Fatal(err)
       }
       fmt.Printf("%q", dump)
    }

    resp, err := c.client.Do(req)
    if err != nil {
        return nil, err
    }
    

    // Debug response
    if debugRun  {
       data, err := ioutil.ReadAll(resp.Body)
       if err != nil {
          panic(err)
        }
        log.Printf("[DEBUG] response data: %v", string(data))
    }

    // response return part of the funcion
    if target  != nil {
        if w, ok := target.(io.Writer); ok {
            io.Copy(w, resp.Body)
        } else {
           decErr  :=  json.NewDecoder(resp.Body).Decode(target)
           if decErr == io.EOF {
              decErr = nil // ignore EOF errors caused by empty response body
           }
           if decErr != nil {
              err = decErr
           }
        }
    }


    return resp, err
}
