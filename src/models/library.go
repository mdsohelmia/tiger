package models

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Library struct {
	ID         string    `gorm:"type:UUID;primaryKey;default:generateUUIDv4()" json:"id,omitempty"`
	CustomerId string    `json:"customer_id,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
	DeletedAt  string    `json:"deleted_at,omitempty"`

	Title  string `json:"title,omitempty"`
	ApiKey string `json:"api_key,omitempty"`
	//Storage Meta
	RawStorage       uint64  `json:"raw_storage,omitempty"`       //Total Raw storage
	TranscodeStorage uint64  `json:"transcode_storage,omitempty"` //Total Transcode storage
	TotalStorage     uint64  `json:"total_storage,omitempty"`
	CdnVolume        float64 `json:"cdn_volume,omitempty"`
	//S3 Bucket Info
	Hostname                      string `json:"-,omitempty"`
	Port                          string `json:"-,omitempty"`
	AccessKeyID                   string `json:"-,omitempty"`
	SecretAccessKey               string `json:"-,omitempty"`
	Region                        string `json:"-,omitempty"`
	Bucket                        string `json:"-,omitempty"`
	EnableSSL                     string `json:"-,omitempty"`
	EnableStorageHighAvailability bool   `json:"enable_storage_high_availability,omitempty"`

	// CDN Distribution
	DistributionUrl string `json:"distribution_url,omitempty"`
	DistributionId  string `json:"distribution_id,omitempty"`
	//Notification
	WebhookHeader string `json:"headers,omitempty"`
	WebhookUrl    string `json:"webhook_url,omitempty"`
	//Video Encoding
	EnabledResolutions string `json:"enabled_resolutions,omitempty"`

	TotalVideo      uint64 `json:"total_video" gorm:"-"`
	CollectionTotal uint64 `json:"collection_count" gorm:"-"`

	//Meta Data
	Description string `json:"description,omitempty" gorm:"-"`
	Options     string `json:"options,omitempty" gorm:"-"`
}

func (lib *Library) BeforeCreate(tx *gorm.DB) (err error) {
	lib.ID = uuid.New().String()

	return
}

func (lib *Library) ApiKeyGenerate() {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	randString := string(b)

	hash := sha1.New()
	hash.Write([]byte(randString))
	bs := hash.Sum(nil)
	lib.ApiKey = fmt.Sprintf("%x", bs)
}
