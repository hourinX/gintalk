package models

type ModelHostConfig struct {
	Host string
	Port string
}

type ModelMysqlConfig struct {
	Host      string
	Port      int
	Username  string
	Password  string
	Database  string
	Charset   string
	ParseTime bool
	Loc       string
}

type ModelRedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type ModelJwtConfig struct {
	Secret      string
	Issuer      string
	ExpireHours int
	RefreshDays int
}

type ModelElasticsearchConfig struct {
	Addresses []string
	Username  string
	Password  string
}

type ModelCaptchaConfig struct {
	ExpireSeconds int
	Length        int
}

type ModelRateLimitConfig struct {
	RegistrationPerMinute int
	IPBlockMinutes        int
}

type Config struct {
	Server        ModelHostConfig
	MySQL         ModelMysqlConfig
	Redis         ModelRedisConfig
	JWT           ModelJwtConfig
	Elasticsearch ModelElasticsearchConfig
	Captcha       ModelCaptchaConfig
	RateLimit     ModelRateLimitConfig
}
