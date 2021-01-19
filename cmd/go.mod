module github.com/inlets/cloud-provision/cmd

go 1.15

replace github.com/inlets/cloud-provision/provision => ../provision

require (
	github.com/inlets/cloud-provision/provision v0.0.0-20210119110409-81f15c0a0897
	github.com/inlets/inletsctl v0.0.0-20210112194340-58379878113b // indirect
)
