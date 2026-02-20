package provision

// MakeExitServerUserdata makes a user-data script in bash to setup inlets
// PRO with a systemd service and the given version.
func MakeExitServerUserdata(authToken, version string) string {
	return MakeExitServerUserdataWithProxyProto(authToken, version, "")
}

// MakeExitServerUserdataWithProxyProto makes a user-data script in bash to setup inlets
// PRO with a systemd service and the given version. If proxyProto is non-empty,
// the PROXY_PROTO environment variable is set in the service configuration.
func MakeExitServerUserdataWithProxyProto(authToken, version, proxyProto string) string {

	return `#!/bin/bash
export AUTHTOKEN="` + authToken + `"
export IP=$(curl -sfSL https://checkip.amazonaws.com)
export PROXY_PROTO="` + proxyProto + `"

curl -SLsf https://github.com/inlets/inlets-pro/releases/download/` + version + `/inlets-pro -o /tmp/inlets-pro && \
  chmod +x /tmp/inlets-pro  && \
  mv /tmp/inlets-pro /usr/local/bin/inlets-pro

curl -SLsf https://github.com/inlets/inlets-pro/releases/download/` + version + `/inlets-pro.service -o inlets-pro.service && \
  mv inlets-pro.service /etc/systemd/system/inlets-pro.service && \
  echo "AUTHTOKEN=$AUTHTOKEN" >> /etc/default/inlets-pro && \
  echo "IP=$IP" >> /etc/default/inlets-pro && \
  echo "PROXY_PROTO=$PROXY_PROTO" >> /etc/default/inlets-pro && \
  systemctl daemon-reload && \
  systemctl start inlets-pro && \
  systemctl enable inlets-pro
`
}
