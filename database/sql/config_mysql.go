package fxsql

import "fmt"

// MysqlAppConfig can be used in an application config to represent database connection details.
// It supports github.com/goph/nest
type MysqlAppConfig struct {
	Host string `env:"" required:"true"`
	Port int    `env:"" default:"3306"`
	User string `env:"" required:"true"`
	Pass string `env:""` // Required removed for now because empty value is not supported by Viper
	Name string `env:"" required:"true"`

	Params map[string]string
}

// Driver returns
func (c MysqlAppConfig) Driver() string {
	return "mysql"
}

// DSN returns the DSN created form the configuration.
func (c MysqlAppConfig) DSN() string {
	var params string

	if len(c.Params) > 0 {
		var query string

		for key, value := range c.Params {
			if query != "" {
				query += "&"
			}

			query += key + "=" + value
		}

		params = "?" + query
	}

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s%s",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
		params,
	)
}
