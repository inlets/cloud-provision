// Copyright (c) Inlets Author(s) 2020. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package provision

import (
	"testing"
)

func Test_makeUserdata_InletsPro(t *testing.T) {
	userData := MakeExitServerUserdata("auth", "0.7.0")

	wantUserdata := `#!/bin/bash
export AUTHTOKEN="auth"
export IP=$(curl -sfSL https://checkip.amazonaws.com)

curl -SLsf https://github.com/inlets/inlets-pro/releases/download/0.7.0/inlets-pro -o /tmp/inlets-pro && \
  chmod +x /tmp/inlets-pro  && \
  mv /tmp/inlets-pro /usr/local/bin/inlets-pro

curl -SLsf https://github.com/inlets/inlets-pro/releases/download/0.7.0/inlets-pro.service -o inlets-pro.service && \
  mv inlets-pro.service /etc/systemd/system/inlets-pro.service && \
  echo "AUTHTOKEN=$AUTHTOKEN" >> /etc/default/inlets-pro && \
  echo "IP=$IP" >> /etc/default/inlets-pro && \
  systemctl daemon-reload && \
  systemctl start inlets-pro && \
  systemctl enable inlets-pro
`

	// ioutil.WriteFile("/tmp/pro", []byte(userData), 0600)
	if userData != wantUserdata {
		t.Errorf("want: %s, but got: %s", wantUserdata, userData)
	}
}
