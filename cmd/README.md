## Provision and automate cloud hosts in Go

This is an example CLI that shows how to use the [provision](https://github.com/inlets/inletsctl/tree/master/pkg/provision) package by inlets to create and automate a cloud host, optionally installing some packages using [user-data/cloud-init](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).

What's this for? This is an example of how to use the provision package as used in [inletsctl](https://github.com/inlets/inletsctl) and the [inlets-operator](https://github.com/inlets/inlets-operator).

* Have fun
* Learn about cloud + platform engineering
* Practice or build your Go skills

Want to know more about inlets? [Inlets is a Cloud Native Tunnel](https://docs.inlets.dev/)
Want to learn Go? Start with Alex's [golang basics series](https://blog.alexellis.io/tag/golang-basics/)

## Tutorial

Checkout the [provision](https://github.com/inlets/inletsctl/tree/master/pkg/provision) README file to find out more, before starting.

Clone/build:

```sh
export GOPATH=$HOME/go
mkdir -p $GOPATH/inlets/
cd $GOPATH/inlets

git clone https://github.com/inlets/provision-example
cd provision-example
go build
```

Now read the code, in main.go, and make sure you understand what's happening. You will need a cloud API token saved into a file, you can create this via your dashboard. The example uses DigitalOcean, but you can customise it to use any provisioner.

Create a file `./cloud-config.txt`

```sh
#cloud-config
packages:
  - nginx
runcmd:
  - systemctl enable nginx
  - systemctl start nginx
```

For more examples of cloud-init, including how to add a custom user and SSH key, see [cloud-init examples](https://cloudinit.readthedocs.io/en/latest/topics/examples.html).

Run the example:

```sh
./provision-example \
  --access-token $(cat ~/Downloads/do-access-token) \
  --userdata-file cloud-config.txt

2020/02/22 11:01:58 Provisioning host with DigitalOcean
Host ID: 181660892
Polling status: 1/250
Polling status: 2/250
Polling status: 3/250
Polling status: 4/250
Polling status: 5/250
Polling status: 6/250
Polling status: 7/250
Polling status: 8/250
Polling status: 9/250
Polling status: 10/250
Polling status: 11/250
Your IP address is: 64.227.34.235
```

Then try the host:

```
curl -s 64.227.34.235 | head -n 4
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
```

Delete any hosts you create via your dashboard, or using the cloud-provider's CLI.

## What next?

Have fun, start learning and customise the example to create something cool.

## Contributing

Please follow the [inlets contributing guide](https://github.com/inlets/inlets/blob/master/CONTRIBUTING.md)

Need support? Join the [OpenFaaS Slack and the #inlets channel](https://slack.openfaas.io/), or open an issue if there is a genuine issue with the code.
