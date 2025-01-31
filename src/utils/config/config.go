package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env"
)

const DEPLOY_MODE_SELF_HOST = "self-host"
const DEPLOY_MODE_CLOUD = "cloud"
const DRIVE_TYPE_AWS = "aws"
const DRIVE_TYPE_MINIO = "minio"

type Config struct {
	// server config
	ServerHost         string `env:"ZWEB_SERVER_HOST"              envDefault:"0.0.0.0"`
	ServerPort         string `env:"ZWEB_SERVER_PORT"              envDefault:"8003"`
	InternalServerPort string `env:"ZWEB_SERVER_INTERNAL_PORT"     envDefault:"9001"`
	ServerMode         string `env:"ZWEB_SERVER_MODE"              envDefault:"debug"`
	DeployMode         string `env:"ZWEB_DEPLOY_MODE"              envDefault:"self-host"`
	ServeHTTPS         string `env:"ZWEB_DEPLOY_SERVE_HTTPS"       envDefault:"false"`

	// storage config
	PostgresAddr     string `env:"ZWEB_SUPERVISOR_PG_ADDR" envDefault:"localhost"`
	PostgresPort     string `env:"ZWEB_SUPERVISOR_PG_PORT" envDefault:"5432"`
	PostgresUser     string `env:"ZWEB_SUPERVISOR_PG_USER" envDefault:"zweb_supervisor"`
	PostgresPassword string `env:"ZWEB_SUPERVISOR_PG_PASSWORD" envDefault:"zweb2022"`
	PostgresDatabase string `env:"ZWEB_SUPERVISOR_PG_DATABASE" envDefault:"zweb_supervisor"`
	// cache config
	RedisAddr     string `env:"ZWEB_REDIS_ADDR" envDefault:"localhost"`
	RedisPort     string `env:"ZWEB_REDIS_PORT" envDefault:"6379"`
	RedisPassword string `env:"ZWEB_REDIS_PASSWORD" envDefault:""`
	RedisDatabase int    `env:"ZWEB_REDIS_DATABASE" envDefault:"0"`

	// drive config
	DriveType             string `env:"ZWEB_DRIVE_TYPE"               envDefault:""`
	DriveAccessKeyID      string `env:"ZWEB_DRIVE_ACCESS_KEY_ID"      envDefault:"minioadmin"`
	DriveAccessKeySecret  string `env:"ZWEB_DRIVE_ACCESS_KEY_SECRET"  envDefault:"minioadmin"`
	DriveRegion           string `env:"ZWEB_DRIVE_REGION"             envDefault:""`
	DriveEndpoint         string `env:"ZWEB_DRIVE_ENDPOINT"           envDefault:"127.0.0.1:9000"`
	DriveSystemBucketName string `env:"ZWEB_DRIVE_SYSTEM_BUCKET_NAME" envDefault:"zweb-supervisor"`
	DriveTeamBucketName   string `env:"ZWEB_DRIVE_TEAM_BUCKET_NAME"   envDefault:"zweb-supervisor-team"`
	DriveUploadTimeoutRaw string `env:"ZWEB_DRIVE_UPLOAD_TIMEOUT"     envDefault:"300s"`
	DriveUploadTimeout    time.Duration
}

func GetConfig() (*Config, error) {
	// fetch
	cfg := &Config{}
	err := env.Parse(cfg)

	// process data
	var errInParseDuration error
	cfg.DriveUploadTimeout, errInParseDuration = time.ParseDuration(cfg.DriveUploadTimeoutRaw)
	if errInParseDuration != nil {
		return nil, errInParseDuration
	}

	// ok
	fmt.Printf("----------------\n")
	fmt.Printf("%+v\n", cfg)
	fmt.Printf("%+v\n", err)

	return cfg, err
}

func (c *Config) IsSelfHostMode() bool {
	if c.DeployMode == DEPLOY_MODE_SELF_HOST {
		return true
	}
	return false
}

func (c *Config) IsCloudMode() bool {
	if c.DeployMode == DEPLOY_MODE_CLOUD {
		return true
	}
	return false
}

func (c *Config) IsServeHTTPS() bool {
	if c.ServeHTTPS == "true" {
		return true
	}
	return false
}

func (c *Config) GetServeHTTPAddress() string {
	if c.ServeHTTPS == "true" {
		return "https://" + c.ServerHost
	}
	return "http://" + c.ServerHost
}

func (c *Config) GetPostgresAddr() string {
	return c.PostgresAddr
}

func (c *Config) GetPostgresPort() string {
	return c.PostgresPort
}

func (c *Config) GetPostgresUser() string {
	return c.PostgresUser
}

func (c *Config) GetPostgresPassword() string {
	return c.PostgresPassword
}

func (c *Config) GetPostgresDatabase() string {
	return c.PostgresDatabase
}

func (c *Config) GetRedisAddr() string {
	return c.RedisAddr
}

func (c *Config) GetRedisPort() string {
	return c.RedisPort
}

func (c *Config) GetRedisPassword() string {
	return c.RedisPassword
}

func (c *Config) GetRedisDatabase() int {
	return c.RedisDatabase
}

func (c *Config) GetDriveType() string {
	return c.DriveType
}

func (c *Config) IsAWSDrive() bool {
	if c.DriveType == DRIVE_TYPE_AWS {
		return true
	}
	return false
}

func (c *Config) IsMINIODrive() bool {
	if c.DriveType == DRIVE_TYPE_MINIO {
		return true
	}
	return false
}

func (c *Config) GetAWSS3AccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetAWSS3AccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetAWSS3Region() string {
	return c.DriveRegion
}

func (c *Config) GetAWSS3SystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetAWSS3TeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetAWSS3Timeout() time.Duration {
	return c.DriveUploadTimeout
}

func (c *Config) GetMINIOAccessKeyID() string {
	return c.DriveAccessKeyID
}

func (c *Config) GetMINIOAccessKeySecret() string {
	return c.DriveAccessKeySecret
}

func (c *Config) GetMINIOEndpoint() string {
	return c.DriveEndpoint
}

func (c *Config) GetMINIOSystemBucketName() string {
	return c.DriveSystemBucketName
}

func (c *Config) GetMINIOTeamBucketName() string {
	return c.DriveTeamBucketName
}

func (c *Config) GetMINIOTimeout() time.Duration {
	return c.DriveUploadTimeout
}
