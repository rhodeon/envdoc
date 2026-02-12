package main

// ServerConfig is the server configuration structure.
// This example demonstrates using envdoc in edit mode to
// maintain documentation directly within a README file.
//
//go:generate go run ../../ -edit -output README.md -format markdown
type ServerConfig struct {
	// Host is the server hostname or IP address to bind to.
	Host string `env:"SERVER_HOST" envDefault:"localhost"`

	// Port is the server port number.
	Port int `env:"SERVER_PORT,required"`

	// TLS configuration
	TLS TLSConfig `envPrefix:"TLS_"`

	// Database connection string.
	DatabaseURL string `env:"DATABASE_URL,required"`

	// Debug enables debug logging when set to true.
	Debug bool `env:"DEBUG" envDefault:"false"`

	// MaxConnections limits the number of concurrent connections.
	MaxConnections int `env:"MAX_CONNECTIONS" envDefault:"100"`
}

// TLSConfig contains TLS/SSL configuration.
type TLSConfig struct {
	// Enabled turns on TLS/SSL.
	Enabled bool `env:"ENABLED" envDefault:"false"`

	// CertFile is the path to the TLS certificate file.
	CertFile string `env:"CERT_FILE"`

	// KeyFile is the path to the TLS private key file.
	KeyFile string `env:"KEY_FILE"`
}
