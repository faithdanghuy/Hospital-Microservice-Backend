package db

import (
	"fmt"
	"time"
)

type Connection struct {
	Host                        string
	Port                        int
	Database                    string
	User                        string
	Password                    string
	SSLMode                     SSLMode
	SSLCertAuthorityCertificate string
	SSLPublicCertificate        string
	SSLPrivateKey               string
	MaxOpenConnections          int
	MaxIdleConnections          int
	ConnectionMaxIdleTime       time.Duration
	ConnectionMaxLifeTime       time.Duration
	ConnectionTimeout           time.Duration
	Replicas                    []Connection
}

func (c Connection) ToConnectionString() string {
	s := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?TimeZone=Asia/Ho_Chi_Minh&sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
		c.SSLMode)

	if c.SSLMode != Disable {
		if c.SSLCertAuthorityCertificate != "" {
			s += fmt.Sprintf("&sslrootcert=%s", c.SSLCertAuthorityCertificate)
		}
		if c.SSLPublicCertificate != "" {
			s += fmt.Sprintf("&sslcert=%s", c.SSLPublicCertificate)
		}
		if c.SSLPrivateKey != "" {
			s += fmt.Sprintf("&sslkey=%s", c.SSLPrivateKey)
		}
	}
	return s
}

type ConnectionOption func(*Connection)

func SetConnection(host string, port int) ConnectionOption {
	return func(c *Connection) {
		c.Host = host
		c.Port = port
	}
}

func SetSSL(mode SSLMode, caCertificate, publicCertificate, privateKey string) ConnectionOption {
	return func(c *Connection) {
		c.SSLMode = mode
		c.SSLCertAuthorityCertificate = caCertificate
		c.SSLPublicCertificate = publicCertificate
		c.SSLPrivateKey = privateKey
	}
}

func SetMaxOpenConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxOpenConnections = max
	}
}

func SetMaxIdleConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxIdleConnections = max
	}
}

func SetConnectionMaxIdleTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxIdleTime = max
	}
}

func SetConnectionMaxLifeTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxLifeTime = max
	}
}

func SetConnectionTimeout(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionTimeout = max
	}
}

func SetLoginCredentials(user, password string) ConnectionOption {
	return func(c *Connection) {
		c.User = user
		c.Password = password
	}
}

func SetDatabase(database string) ConnectionOption {
	return func(c *Connection) {
		c.Database = database
	}
}
