package config

// StorageDriver is a kind of storage backend
type StorageDriver string

type StorageConfig struct {
	Driver     string           `mapstructure:"driver" json:"driver" default:"fs" env:"STORAGE_DRIVER"`
	Filesystem FileSystemConfig `mapstructure:"fs" json:"fs"`
	S3         S3Config         `mapstructure:"s3" json:"s3"`
}

type S3Config struct {
	Hostname        string `mapstructure:"hostname" json:"hostname" default:"" env:"S3_HOSTNAME"`
	Port            string `mapstructure:"port" json:"port" default:"" env:"S3_PORT"`
	AccessKeyID     string `mapstructure:"access_key_id" json:"access_key_id" env:"S3_ACCESS_KEY_ID"`
	SecretAccessKey string `mapstructure:"secret_access_key" json:"secret_access_key" env:"S3_SECRET_ACCESS_KEY"`
	Region          string `mapstructure:"region" json:"region" env:"S3_REGION"`
	Bucket          string `mapstructure:"bucket" json:"bucket" env:"S3_BUCKET"`
	EnableSSL       bool   `mapstructure:"enable_ssl" json:"enable_ssl" default:"true" env:"S3_ENABLE_SSL"`
	UsePathStyle    bool   `mapstructure:"use_path_style" json:"use_path_style" default:"false" env:"S3_ENABLE_PATH_STYLE"`
}

type FileSystemConfig struct {
	DataPath string `mapstructure:"data_path" json:"data_path" default:"data" env:"FS_DATA_PATH"`
}
