module github.com/palourde/revocation-test

go 1.14

require (
	github.com/cloudflare/cfssl v1.4.1
)

replace github.com/cloudflare/cfssl => ./cfssl
