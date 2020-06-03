module github.com/tylerconlee/go-zendesk

require (
	github.com/golang/mock v1.4.3
	github.com/google/go-querystring v1.0.0
	github.com/nukosuke/go-zendesk v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.0 // indirect
	github.com/tidwall/gjson v1.6.0
)

replace github.com/nukosuke/go-zendesk => github.com/tylerconlee/go-zendesk v0.7.2

go 1.13
