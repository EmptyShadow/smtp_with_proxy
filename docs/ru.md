# SMTP через прокси
___

Модуль используется для установления соединения с почтовым сервисом по протоколу SMTP, используя прокси сервер.

Первый способ использования модуля
```go
func Example(addressProxy string, email string) (*smtp.Client, error) {
	// Получение хоста из электроной почты
	i := strings.LastIndexByte(email, '@')
	host = email[i+1:]	
	
	// Получение MX записей по хосту
    mx, err = net.LookupMX(host)
    if err != nil {
    	return nil, err
    }
    
    // Проверка существования записей
    if mx == nil || len(mx) == 0 {
    	return nil, errors.New("mx records not found")
    }
    
    // Удаление точки у mx записи
    address := mx[0].Host[0:len(mx[0].Host) -1]
    
    // Парсинг адреса прокси
    proxyURL, err := url.Parse("//" + addressProxy)
    if err != nil {
    	return nil, err
    }
    
    // Получения данных для авторизации
    user := proxyURL.User.Username()
    pas, _ := proxyURL.User.Password()
    
    // Получение объекта прокси
    proxy, err := NewProxy(user, pas, proxyURL.Host)
    if err != nil {
        return nil, err
    }
    
    // Привязка соединения
    return proxy.GetSMTPClientUseSOCKS5(fmt.Sprintf("%s:%d", address, 25))
}

```

Второй способ использования модуля
```go
func Example(addressProxy string, email string) (*smtp.Client, error) {
	// Получение хоста из электроной почты
	i := strings.LastIndexByte(email, '@')
	host = email[i+1:]	
	
	// Получение MX записей по хосту
    mx, err = net.LookupMX(host)
    if err != nil {
    	return nil, err
    }
    
    // Проверка существования записей
    if mx == nil || len(mx) == 0 {
    	return nil, errors.New("mx records not found")
    }
    
    // Удаление точки у mx записи
    address := mx[0].Host[0:len(mx[0].Host) -1]
        
    // Парсинг адреса прокси
    proxyURL, err := url.Parse("//" + addressProxy)
    if err != nil {
        return nil, err
    }
        
    // Получение объекта прокси
    proxy, err := NewProxyByURL(proxyURL)
    if err != nil {
        return nil, err
    }
    
    // Привязка соединения
    return proxy.GetSMTPClientUseSOCKS5(fmt.Sprintf("%s:%d", address, 25))
}

```

Третий способ использования модуля
```go
func Example(addressProxy string, email string) (*smtp.Client, error) {
	// Получение хоста из электроной почты
	i := strings.LastIndexByte(email, '@')
	host = email[i+1:]	
	
	// Получение MX записей по хосту
    mx, err = net.LookupMX(host)
    if err != nil {
    	return nil, err
    }
    
    // Проверка существования записей
    if mx == nil || len(mx) == 0 {
    	return nil, errors.New("mx records not found")
    }
    
    // Удаление точки у mx записи
    address := mx[0].Host[0:len(mx[0].Host) -1]
    
    // Получение объекта прокси
    proxy, err := NewProxyByStringURL(addressProxy)
    if err != nil {
        return nil, err
    }
    
    // Привязка соединения
    return proxy.GetSMTPClientUseSOCKS5(fmt.Sprintf("%s:%d", address, 25))
}

```