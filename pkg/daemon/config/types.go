//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package config

// The structure of the config to run the daemon
type Config struct {
	Debug       bool
	SecretToken string
	TemplateRegistry struct {
		Host string
	}
	ProxyServer struct {
		Port int
	}
	HttpServer struct {
		Host string
		Port int
	}
	Etcd struct {
		Endpoints []string
		TLS       struct {
			Key  string
			Cert string
			CA   string
		}
		Quorum bool
	}
	Registry struct {
		Server   string
		Username string
		Password string
	}
	VCS struct {
		Github struct {
			Client struct {
				ID       string
				SecretID string
			}
			RedirectUri string
		}
		Bitbucket struct {
			Client struct {
				ID       string
				SecretID string
			}
			RedirectUri string
		}
		Gitlab struct {
			Client struct {
				ID       string
				SecretID string
			}
			RedirectUri string
		}
	}
}
