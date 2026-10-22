package config

import "os"

type Config struct { Port string; DatabaseURL string; AllowedOrigin string }

func Load() Config {
 return Config{Port:value("PORT","8080"),DatabaseURL:value("DATABASE_URL","postgres://sea:local_dev_password@localhost:5432/morehodka?sslmode=disable"),AllowedOrigin:value("ALLOWED_ORIGIN","http://localhost:5173")}
}
func value(key,fallback string) string { if v:=os.Getenv(key); v!="" { return v }; return fallback }
