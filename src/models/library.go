package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/go-clickhouse/ch"
)

type Library struct {
	ch.CHModel `ch:"partition:toYYYYMM(created_at)"`

	ID         uuid.UUID `json:"id,omitempty" ch:""`
	CustomerId string    `json:"customer_id"`
	CreatedAt  time.Time `json:"created_at" ch:",pk"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`

	Title  string `json:"title"`
	ApiKey string `json:"api_key"`
	//Storage Meta
	RawStorage       int `json:"raw_storage"`       //Total Raw storage
	TranscodeStorage int `json:"transcode_storage"` //Total Transcode storage
	TotalStorage     int `json:"total_storage"`
	CdnVolume        int `json:"cdn_volume"`
	//S3 Bucket Info
	Hostname                      string `json:"-"`
	Port                          string `json:"-"`
	AccessKeyID                   string `json:"-"`
	SecretAccessKey               string `json:"-"`
	Region                        string `json:"-"`
	Bucket                        string `json:"-"`
	EnableSSL                     string `json:"-"`
	EnableStorageHighAvailability bool   `json:"enable_storage_high_availability"`

	// CDN Distribution
	DistributionUrl string `json:"distribution_url"`
	DistributionId  string `json:"distribution_id"`
	//Notification
	WebhookHeader string `json:"headers"`
	WebhookUrl    string `json:"webhook_url"`
	//Video Encoding
	EnabledResolutions string `json:"enabled_resolutions"`
}
