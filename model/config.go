package model

type Config struct {
	DB     PostgreCfg   `json:"postgresql"`
	Redis  RedisCfg     `json:"redis"`
	JWT    JWTCfg       `json:"jwt"`
	S3Cred S3Credential `json:"s3"`
}

type PostgreCfg struct {
	Address  string `json:"address"`
	DBName   string `json:"db-name"`
	User     string `json:"username"`
	Password string `json:"password"`
	Port     string `json:"port"`
}

type RedisCfg struct {
	Host     string `json:"host"`
	Password string `json:"password"`
}

type JWTCfg struct {
	SignKey string `json:"signKey"`
}

type S3Credential struct {
	AccessID   string `json:"access_id"`
	Secret     string `json:"secret"`
	BucketName string `json:"bucket_name"`
}
