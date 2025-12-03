module github.com/Arubacloud/arubacloud-provider-kog/storage-plugin

go 1.24.2

toolchain go1.24.4

require (
	github.com/Arubacloud/arubacloud-provider-kog/pkg v0.0.0
	github.com/rs/zerolog v1.34.0
	github.com/swaggo/http-swagger v1.3.4
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/swaggo/files v1.0.1 // indirect
	github.com/swaggo/swag v1.16.4 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/tools v0.33.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require github.com/krateoplatformops/plumbing v0.5.5 // indirect

replace github.com/Arubacloud/arubacloud-provider-kog/pkg => ../../pkg
