module github.com/orange-cloudfoundry/terraform-provider-cfsecurity

go 1.16

replace github.com/codahale/hdrhistogram => github.com/HdrHistogram/hdrhistogram-go v0.0.0-20210305173142-35c7773a578a

require (
	github.com/cloudfoundry-community/go-cfclient v0.0.0-20210525151336-ed51ca3339e2
	github.com/hashicorp/go-uuid v1.0.2
	github.com/hashicorp/terraform-plugin-sdk v1.17.2
	github.com/orange-cloudfoundry/cf-security-entitlement v0.1.18
)
