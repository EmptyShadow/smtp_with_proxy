# SMTP through proxy
___

The module is used to establish a connection with the mail service via the SMT protocol using a proxy server.

The first way to use the module
```go
func Example (addressProxy string, email string) (* smtp.Client, error) {
    // Get a host from email
    i: = strings.LastIndexByte (email, '@')
    host = email [i + 1:]

    // Get MX records by host
    mx, err = net.LookupMX (host)
    if err! = nil {
    return nil, err
    }
    
    // Check the existence of records
    if mx == nil || len (mx) == 0 {
    return nil, errors.New ("mx records not found")
    }
    
    // Delete point at mx record
    address: = mx [0] .Host [0: len (mx [0] .Host) -1]
    
    // Parsing the proxy address
    proxyURL, err: = url.Parse ("//" + addressProxy)
    if err! = nil {
    return nil, err
    }
    
    // Receiving data for authorization
    user: = proxyURL.User.Username ()
    pas, _: = proxyURL.User.Password ()
    
    // Get the proxy object
    proxy, err: = NewProxy (user, pas, proxyURL.Host)
    if err! = nil {
        return nil, err
    }
    
    // binding the connection
    return proxy.GetSMTPClientUseSOCKS5 (fmt.Sprintf ("% s:% d", address, 25))
}

```

The second way to use the module
```go
func Example (addressProxy string, email string) (* smtp.Client, error) {
    // Get a host from email
    i: = strings.LastIndexByte (email, '@')
    host = email [i + 1:]

    // Get MX records by host
    mx, err = net.LookupMX (host)
    if err! = nil {
    return nil, err
    }
    
    // Check the existence of records
    if mx == nil || len (mx) == 0 {
    return nil, errors.New ("mx records not found")
    }
    
    // Delete point at mx record
    address: = mx [0] .Host [0: len (mx [0] .Host) -1]
        
    // Parsing the proxy address
    proxyURL, err: = url.Parse ("//" + addressProxy)
    if err! = nil {
        return nil, err
    }
        
    // Get the proxy object
    proxy, err: = NewProxyByURL (proxyURL)
    if err! = nil {
        return nil, err
    }
    
    // binding the connection
    return proxy.GetSMTPClientUseSOCKS5 (fmt.Sprintf ("% s:% d", address, 25))
}

```

The third way to use the module
```go
func Example (addressProxy string, email string) (* smtp.Client, error) {
    // Get a host from email
    i: = strings.LastIndexByte (email, '@')
    host = email [i + 1:]

    // Get MX records by host
    mx, err = net.LookupMX (host)
    if err! = nil {
    return nil, err
    }
    
    // Check the existence of records
    if mx == nil || len (mx) == 0 {
    return nil, errors.New ("mx records not found")
    }
    
    // Delete point at mx record
    address: = mx [0] .Host [0: len (mx [0] .Host) -1]
    
    // Get the proxy object
    proxy, err: = NewProxyByStringURL (addressProxy)
    if err! = nil {
        return nil, err
    }
    
    // binding the connection
    return proxy.GetSMTPClientUseSOCKS5 (fmt.Sprintf ("% s:% d", address, 25))
}

```