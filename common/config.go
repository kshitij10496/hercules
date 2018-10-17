package common

// Config represents configuration environment to running application
type Config struct {
	DatabaseURL string `required:"true" desc:"Connection String, for more info see: https://godoc.org/github.com/lib/pq"`
	Port        string `default:"8080" desc:"Application listen port"`
}
