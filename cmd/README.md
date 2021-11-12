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

## What about SSH?

Here's an example with my SSH key pre-installed on the host:

```yaml
#cloud-config
ssh_authorized_keys:
## Note: Replace with your own public key
  - ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQD5P+iIgayz0nM6ra5y25KyKEPZyyRiXssOQ92GS7+JHKcerHndzbgWHUBH4Sc0HE/IeIFiZ20iYVBMLJZteXfrrEV91LNPApZ010RRNehzGIfPj0C/5bNVA3NtwsXWtz6O17goEPlblhcbJ1XoS5xdy4U2GMfaB8C9bZ0RvYFuXP+FJfgvJ9mp4MesLKMH/rKUc7uLIWCQWgrTLoTx1r+merWkQCiWU8onvfh+B7vXggY9ffOxADJ+JqXjEbG49CiXG3zJq2xnbiyf0zVvvG2Utr45lPm1cxbqch5BdJrZpIb8qSvjuV/oq6AUvUqpBei2YLZKz1sPKONkBB1t5e3Xa9+PYB19PUmn8/WUQYWGJ4LB5mTe87nfs1Q1p/cQe4pq+Y8s3rUitnqwv16g5CdUAXG8KPWlAXB+VK04cj1E3CYkEOUIeeeyUDlMPezrEFEKjDcqhaReDMjHMma95SeMjlt3ZrO2FsXjmgLfjv5kvsWFZdDjl3zQovVfs+pVzGk= alex@am1.local
runcmd:
  - curl -SLs https://get.k3sup.dev | sh
  - k3sup install --local --k3s-channel stable
```

Then:

```bash
Polling status: 15/250
Your IP address is: 64.227.37.14
```

Connect with:

```bash

# Log in to change root password, check your email for the value
ssh root@64.227.37.14

# Then run the below:
ssh root@64.227.37.14 "sudo kubectl get nodes -o wide"

NAME                STATUS   ROLES                  AGE   VERSION        INTERNAL-IP     EXTERNAL-IP   OS-IMAGE             KERNEL-VERSION       CONTAINER-RUNTIME
provision-example   Ready    control-plane,master   21s   v1.21.5+k3s2   64.227.42.201   <none>        Ubuntu 18.04.6 LTS   4.15.0-159-generic   containerd://1.4.11-k3s1
```

Merge it to your local kubeconfig for administration by running k3sup on your local machine:

```bash
k3sup install --skip-install --ip 64.227.37.14 \
 --user root \
 --local-path $HOME/.kube/config \
 --context auto-k3s

 Saving file to: /Users/alex/.kube/config

# Test your cluster with:
export KUBECONFIG=/Users/alex/.kube/config
kubectl config set-context auto-k3s
kubectl get node -o wide

NAME                STATUS   ROLES                  AGE   VERSION
provision-example   Ready    control-plane,master   68s   v1.21.5+k3s2
```

Need to troubleshoot?

Log in to the instance via SSH and explore the logs:

```bash
cat /var/log/cloud-init.log
cat /var/log/cloud-init-output.log
```

## What next?

Have fun, start learning and customise the example to create something cool.

Remember to delete the instances you create using this tool.

## Contributing

Please follow the [inlets contributing guide](https://github.com/inlets/inlets/blob/master/CONTRIBUTING.md)

Need support? Join the [OpenFaaS Slack and the #inlets channel](https://slack.openfaas.io/), or open an issue if there is a genuine issue with the code.
