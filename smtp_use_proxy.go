package smtp_use_proxy

import (
	"golang.org/x/net/proxy"
	"net"
	"net/url"
	"net/smtp"
)

// Прокси
// auth - объект авторизации в прокси
// address - хост:порт
//
// Proxy
// auth - proxy authorization object
// address - host: port
type Proxy struct {
	auth    *proxy.Auth
	address string
}

// Получить объект прокси по строке url адреса прокси
// surl - строка вида [user[:password]@]host:port
//
// Get proxy object by proxy url string
// surl - a string of the form [user[:password]@]host: port
func NewProxyByStringURL(surl string) (*Proxy, error) {
	ourl, err := url.Parse(surl)
	if err != nil {
		return nil, err
	}

	return NewProxyByURL(ourl)
}

// Получить объект прокси по объекту url
//
// Get the proxy object by url object
func NewProxyByURL(ourl *url.URL) (*Proxy, error) {
	username := ourl.User.Username()
	password, _ := ourl.User.Password()
	return NewProxy(username, password, ourl.Host)
}

// Получить объект прокси
// username - имя владельца прокси, может быть пустым, если авторизация не требуется
// password - пароль от прокси, может быть пустым, если авторизация не требуется
// address - адрес прокси в формате host:port
//
// Get proxy object
// username - the name of the proxy owner, may be empty if authorization is not required
// password - proxy password, may be empty if authorization is not required
// address - proxy address in host:port format
func NewProxy(username string, password string, address string) (*Proxy, error) {
	p := &Proxy{
		address: address,
	}

	if username != "" {
		p.auth = &proxy.Auth{
			User:     username,
			Password: password,
		}
	}

	return p, nil
}

// Получить объект авторизации
//
// Get an authorization object
func (p *Proxy) GetAuth() *proxy.Auth {
	return p.auth
}

// Получить адресс прокси
//
// Get proxy address
func (p *Proxy) GetAddress() string {
	return p.address
}

// Получить SMTP клиента используя соединение SOCKS5 с прокси
// address - адресс хоста с которым нужно установить соединение про протоколу smtp
//
// Get client SMTP using SOCKS5 connection with proxy
// address - the address of the host with which you want to establish a connection about the smtp protocol
func (p *Proxy) GetSMTPClientUseSOCKS5(address string) (*smtp.Client, error) {
	dialer, err := proxy.SOCKS5("tcp", p.address, p.auth, proxy.Direct)
	if err != nil {
		return nil, err
	}

	conn, err := dialer.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	host, _, _ := net.SplitHostPort(address)
	c, err := smtp.NewClient(conn, host)

	return c, err
}