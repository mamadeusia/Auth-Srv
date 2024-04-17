package config

// Postgres - keep postgres config
type Postgres struct {
	URL string
}

func PostgresURL() string {
	return cfg.Postgres.URL
}
