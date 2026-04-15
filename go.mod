module github.com/Thomazoide/IACam_backend

go 1.25.2

require (
	github.com/docker/docker v24.0.7+incompatible
	github.com/go-chi/chi/v5 v5.2.5
	github.com/gorilla/websocket v1.5.3
	github.com/joho/godotenv v1.5.1
	gorm.io/driver/mysql v1.6.0
	gorm.io/gorm v1.31.1
)

replace github.com/docker/distribution => github.com/docker/distribution v2.8.1+incompatible

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/Microsoft/go-winio v0.4.21 // indirect
	github.com/docker/distribution v2.8.3+incompatible // indirect
	github.com/docker/go-connections v0.6.0 // indirect
	github.com/docker/go-units v0.5.0 // indirect
	github.com/go-chi/cors v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/moby/term v0.5.2 // indirect
	github.com/morikuni/aec v1.1.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	golang.org/x/time v0.15.0 // indirect
	gotest.tools/v3 v3.5.2 // indirect
)
