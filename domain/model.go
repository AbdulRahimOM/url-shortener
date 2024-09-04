package domain

type UrlRecord struct {
	ID           uint   `gorm:"primaryKey"`
	LongUrl      string `gorm:"long_url;type:varchar(255);not null; unique"`
	ShortUrlPath string `gorm:"short_url_path;type:varchar(200);not null; unique"`
}
