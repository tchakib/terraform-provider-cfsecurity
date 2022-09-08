package cfsecurity

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	clients "github.com/cloudfoundry-community/go-cf-clients-helper/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/orange-cloudfoundry/cf-security-entitlement/client"
)

var expiresAt time.Time

// Provider -
func Provider() terraform.ResourceProvider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"cf_api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_API_URL", ""),
			},
			"cf_security_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_SECURITY_URL", ""),
			},
			"user": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_USER", "admin"),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_PASSWORD", ""),
			},
			"cf_client_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_CLIENT_ID", ""),
			},
			"cf_client_secret": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_CLIENT_SECRET", ""),
			},
			"skip_ssl_validation": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CF_SKIP_SSL_VALIDATION", false),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"cfsecurity_asg": dataSourceAsg(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"cfsecurity_bind_asg":    resourceBindAsg(),
			"cfsecurity_entitle_asg": resourceEntitleAsg(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := &clients.Config{
		Endpoint:          d.Get("cf_api_url").(string),
		User:              d.Get("user").(string),
		Password:          d.Get("password").(string),
		CFClientID:        d.Get("cf_client_id").(string),
		CFClientSecret:    d.Get("cf_client_secret").(string),
		SkipSslValidation: d.Get("skip_ssl_validation").(bool),
	}

	s, err := clients.NewSession(*config)
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(d.Get("cf_api_url").(string))
	if err != nil {
		return nil, err
	}
	pHost := strings.SplitN(uri.Host, ".", 2)
	pHost[0] = "cfsecurity"
	uri.Host = strings.Join(pHost, ".")
	uri.Path = ""

	securityEndpoint := uri.String()
	if tmpSecEndpoint, ok := d.GetOk("cf_security_url"); ok {
		securityEndpoint = tmpSecEndpoint.(string)
	}

	expiresAt, err = getExpiresAtFromToken(s.ConfigStore().AccessToken())
	if err != nil {
		return nil, err
	}

	return client.NewClient(securityEndpoint, s.V3(), s.ConfigStore().AccessToken(), config.Endpoint,
		http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: d.Get("skip_ssl_validation").(bool)},
		}), nil
}

func getExpiresAtFromToken(accessToken string) (time.Time, error) {
	tokenSplit := strings.Split(accessToken, ".")
	if len(tokenSplit) < 3 {
		return time.Now(), fmt.Errorf("not a jwt")
	}

	decodeToken, err := base64.RawStdEncoding.DecodeString(tokenSplit[1])
	if err != nil {
		return time.Now(), err
	}

	token := struct {
		Exp int `json:"exp"`
	}{}

	err = json.Unmarshal(decodeToken, &token)
	if err != nil {
		return time.Now(), err
	}

	expAt := time.Unix(int64(token.Exp), 0)
	// Ajout d'une marge d'une minute pour l'expiration
	expAtBefore := expAt.Add(time.Duration(-1) * time.Minute)

	return expAtBefore, nil

}

func refreshTokenIfExpires(d *schema.ResourceData, c client.Client) error {
	isExpires := false
	if expiresAt.Before(time.Now()) {
		isExpires = true

		config := &clients.Config{
			Endpoint:          d.Get("cf_api_url").(string),
			User:              d.Get("user").(string),
			Password:          d.Get("password").(string),
			CFClientID:        d.Get("cf_client_id").(string),
			CFClientSecret:    d.Get("cf_client_secret").(string),
			SkipSslValidation: d.Get("skip_ssl_validation").(bool),
		}

		s, err := clients.NewSession(*config)
		if err != nil {
			return err
		}

		accessToken := s.ConfigStore().AccessToken()
		refreshExpiresAt, err := getExpiresAtFromToken(accessToken)
		if err != nil {
			return err
		}

		if isExpires {
			expiresAt = refreshExpiresAt
			c.SetAccessToken(accessToken)
			return nil
		}

	}
	return nil

}
