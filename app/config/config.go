package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"gopkg.in/yaml.v2"
)

// APIConfig is used to define the configuration of the API
type APIConfig struct {
	Host                string                     `yaml:"host"`                // Host is the address of the server
	Port                int                        `yaml:"port"`                // Port is the port of the server
	AuthConfig          AuthConfig                 `yaml:"auth"`                // AuthConfig is the configuration for authentication
	AutoMigrate         bool                       `yaml:"autoMigrate"`         // AutoMigrate is a flag to determine if the database should be migrated automatically
	Connect             bool                       `yaml:"connect"`             // Connect is a flag to determine if the database should be connected automatically
	DBConnectionStr     string                     `yaml:"dbConnectionStr"`     // DBConnectionStr is the connection string for the database
	ExamTypes           []string                   `yaml:"examTypes"`           // ExamTypes is the list of exam types
	ExamEvaluationTypes []ExamEvaluationTypeConfig `yaml:"examEvaluationTypes"` // EvalTypes is the list of evaluation types
	GradeTypes          []string                   `yaml:"gradeTypes"`          // GradeTypes is the list of grade types
	States              []string                   `yaml:"states"`              // States is the list of states
	CurriculumTypes     []CurriculumTypeConfig     `yaml:"curriculumTypes"`     // CurriculumTypes is the list of curriculum types
	Roles               []RoleConfig               `yaml:"roles"`               // Roles is the list of roles
	Cors                CorsConfig                 `yaml:"cors"`                // Cors is the configuration for CORS
	MaxBodySize         string                     `yaml:"maxBodySize"`         // BodySize is the maximum size of the request body
	RateLimit           RateLimitConfig            `yaml:"rateLimit"`           // RateLimit is the configuration for rate limiting
	LogLevel            string                     `yaml:"logLevel"`            // LogLevel is the preferred log level
}

// Configuration for CORS
type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowedOrigins"` // AllowedOrigins is the list of allowed origins
	AllowedHeaders []string `yaml:"allowedHeaders"` // AllowedHeaders is the list of allowed headers
}

// Configuration for a role
type RoleConfig struct {
	Name    string   `yaml:"name"`    // Name of the role
	Members []string `yaml:"members"` // Statically assigned members of the role
}

type RateLimitConfig struct {
	Burst  int `yaml:"burst"`  // Burst is the maximum number of requests that can be made in a second (RPS)
	Rate   int `yaml:"rate"`   // Rate is the usual number of requests that can be made in a second (RPS)
	Expiry int `yaml:"expiry"` // Expiry is the time in seconds until the rate limit is reset
}

// Configuration for Authentication
type AuthConfig struct {
	Enabled              bool   `yaml:"enabled"`              // Enabled is a flag to determine if authentication is enabled
	OAuthKeyDiscoveryURL string `yaml:"oAuthKeyDiscoveryURL"` // OAuthKeyDiscoveryURL is the URL to discover the OAuth keys
}

// Predefined exam evaluation types
type ExamEvaluationTypeConfig struct {
	Code        string `yaml:"code"`        // code is the code of the evaluation type
	Description string `yaml:"description"` // description is the description of the evaluation type
}

// Predefined curriculum types
type CurriculumTypeConfig struct {
	Description   string `yaml:"description"`
	DurationYears uint   `yaml:"durationyears"`
}

// Returns a new configuration with default values
// This is used to create the config file if it does not exist
func NewAPIConfig() *APIConfig {
	return &APIConfig{
		Host:    "0.0.0.0",
		Port:    9000,
		Connect: true,
		AuthConfig: AuthConfig{
			Enabled:              false,
			OAuthKeyDiscoveryURL: "https://login.microsoftonline.com/common/discovery/keys",
		},
		AutoMigrate:     true,
		DBConnectionStr: "sqlite://:memory:",
		ExamTypes:       []string{"Mündliche oder schriftliche Prüfung ", "Präsentationen", "Lernbericht", "schriftliche Arbeit", "Lernjournal"},
		GradeTypes:      []string{"Kein Eintrag", "Note (1-6)", "Prozentwert (0-100)"},
		ExamEvaluationTypes: []ExamEvaluationTypeConfig{
			{Code: "F", Description: "Modul bestanden, wenn jeder Kurs eine genügende Bewertung aufweist. (Art. 29)"},
			{Code: "M", Description: "Modul bestanden, wenn der Durchschnitt aller Kurse genügend und nicht mehr als ein Kurs im Modul ungenügend ist. (Art. 30)"},
			{Code: "D", Description: "Modul bestanden, wenn der Durchschnitt der Kurse genügend ist (mehr als 60%). (Art. 31)"},
			{Code: "E", Description: "Modul bestanden, wenn alle Kurse erfüllt sind. (Art. 32)"},
		},
		States: []string{"Aktiv", "Inaktiv"},
		CurriculumTypes: []CurriculumTypeConfig{
			{Description: "Vollzeit", DurationYears: 2}, {Description: "Teilzeit", DurationYears: 3},
		},
		Roles: []RoleConfig{
			{Name: "Kursadministrator", Members: []string{}},
			{Name: "Fachbereichsleiter", Members: []string{}},
			{Name: "Dozent", Members: []string{}},
			{Name: "Student", Members: []string{}},
		},
		Cors: CorsConfig{
			AllowedOrigins: []string{"https://localhost:9000", "api.smashgrade.ch"},
			AllowedHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		},
		MaxBodySize: "2M",
		RateLimit: RateLimitConfig{
			Burst:  20, // 20 requests per second
			Rate:   10, // 10 requests per second
			Expiry: 60, // 60 seconds
		},
		LogLevel: "INFO",
	}
}

// Returns the server address as a string
func (c *APIConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// Returns the echo logger configuration based on the API configuration
func (c *APIConfig) GetEchoLoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}
}

// Returns the echo CORS configuration based on the API configuration
func (c *APIConfig) GetEchoCORSConfig() middleware.CORSConfig {
	return middleware.CORSConfig{
		AllowOrigins: c.Cors.AllowedOrigins,
		AllowHeaders: c.Cors.AllowedHeaders,
	}
}

// Returns the rate limiter configuration based on the API configuration
func (c *APIConfig) GetRateLimitConfig() middleware.RateLimiterConfig {
	// Use the default configuration
	rc := middleware.DefaultRateLimiterConfig
	// Update the store with the limits from the API configuration
	rc.Store = middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{
		Rate:      rate.Limit(c.RateLimit.Rate),
		Burst:     c.RateLimit.Burst,
		ExpiresIn: time.Duration(c.RateLimit.Expiry) * time.Second,
	})
	return rc
}

// Overwrites the current configuration with environment variables
func (c *APIConfig) FromEnv() {
	log.Println("Overwriting configuration from environment variables...")
	// Check for production environment first, some variables may be overwritten later
	if env, ok := os.LookupEnv("ENV"); ok {
		if strings.ToLower(env) == "prod" {

			// Enable authentication in production environment
			c.AuthConfig.Enabled = true
			// Enable automatic connection to the database in production environment
			c.Connect = true
			// Enable automatic migration of the database in production environment
			c.AutoMigrate = true
		}
	}
	// Check for environment variables and overwrite the configuration
	if host, ok := os.LookupEnv("API_HOST"); ok {
		c.Host = host
	}
	if port, ok := os.LookupEnv("API_PORT"); ok {
		// Check if the port is a valid integer and set it
		v, err := strconv.Atoi(port)
		if err == nil {
			c.Port = v
		}
	}
	if connect, ok := os.LookupEnv("API_CONNECT"); ok {
		// Check if the connect flag is a valid boolean and set it
		v, err := strconv.ParseBool(connect)
		if err == nil {
			c.Connect = v
		}
	}
	if autoMigrate, ok := os.LookupEnv("API_AUTO_MIGRATE"); ok {
		// Check if the auto migrate flag is a valid boolean and set it
		v, err := strconv.ParseBool(autoMigrate)
		if err == nil {
			c.AutoMigrate = v
		}
	}
	if authEnabled, ok := os.LookupEnv("API_AUTH_ENABLED"); ok {
		// Check if the auth enabled flag is a valid boolean and set it
		v, err := strconv.ParseBool(authEnabled)
		if err == nil {
			c.AuthConfig.Enabled = v
		}
	}
	if dbConnectionStr, ok := os.LookupEnv("API_DB_CONNECTION_STR"); ok {
		c.DBConnectionStr = dbConnectionStr
	}
	if oAuthKeyDiscoveryURL, ok := os.LookupEnv("API_AUTH_OAUTH_KEY_DISCOVERY_URL"); ok {
		c.AuthConfig.OAuthKeyDiscoveryURL = oAuthKeyDiscoveryURL
	}
	log.Println("Configuration overwritten successfully")
}

// Loads the configuration from a file
// Attempts to write default configuration to file if file does not exist
func FromFile(path string) *APIConfig {
	log.Println("Loading default configuration...")
	config := NewAPIConfig()
	log.Printf("Loading configuration from file: %s...\n", path)
	cf, err := os.ReadFile(path)
	if err == nil {
		err = yaml.Unmarshal(cf, config)
		if err != nil {
			log.Println("Configuration loaded successfully")
		} else {
			log.Println("Error loading configuration from file: could not decode configuration from YAML")
		}
	} else {
		log.Println("Error loading configuration from file: Could not open file for reading or file does not exist")
		log.Println("Attempting to save default configuration to file...")
		ToFile(path, config)
	}
	// Update the configuration with environment variables
	config.FromEnv()
	// Return the configuration
	return config
}

// Saves the configuration to a file
func ToFile(path string, config *APIConfig) {
	log.Printf("Saving configuration to file: %s...", path)
	cf, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err == nil {
		err = yaml.NewEncoder(cf).Encode(config)
		if err != nil {
			log.Println("Configuration saved successfully")
		} else {
			log.Println("Error saving configuration to file: could not encode configuration as YAML")
		}
	} else {
		log.Println("Error saving configuration to file: could not open file for writing")
	}
}
