package config

type App struct {
	Port                 string `env:"PORT" mapstructure:"PORT"`
	Host                 string `env:"HOST" mapstructure:"HOST"`
	MongoHost            string `env:"MONGO_HOST" mapstructure:"MONGO_HOST"`
	MongoPort            string `env:"MONGO_PORT" mapstructure:"MONGO_PORT"`
	MongoName            string `env:"MONGO_NAME" mapstructure:"MONGO_NAME"`
	MongoUsername        string `env:"MONGO_USERNAME" mapstructure:"MONGO_USERNAME"`
	MongoPassword        string `env:"MONGO_PASSWORD" mapstructure:"MONGO_PASSWORD"`
	MongoCollection      string `env:"MONGO_COLLECTION" mapstructure:"MONGO_COLLECTION"`
	CorsAllowOrigins     string `env:"CORS_ALLOW_ORIGINS" mapstructure:"CORS_ALLOW_ORIGINS"`
	CorsAllowMethods     string `env:"CORS_ALLOW_METHODS" mapstructure:"CORS_ALLOW_METHODS"`
	CorsAllowHeaders     string `env:"CORS_ALLOW_HEADERS" mapstructure:"CORS_ALLOW_HEADERS"`
	CorsAllowCredentials bool   `env:"CORS_ALLOW_CREDENTIALS" mapstructure:"CORS_ALLOW_CREDENTIALS"`
}
