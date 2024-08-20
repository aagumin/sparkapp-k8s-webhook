# Spark Application admission webhook.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

Kubernetes [mutation webhook](https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/)
for [SparkApplicaton](https://github.com/kubeflow/spark-operator). Mutate spark application upon creation.

## For what?

- Hiding the complexity of changing the environment when developing an application
- Automatic setting of important properties
- Easier CRD configuration based on annotations your team is familiar with
- I want

## Features

- [x] Set toleration and affinity based on team key.
- [ ] Auto monitoring
- [ ] Auto spark history server


## Deployment

To deploy this project run

```bash
  helm repo add //TODO
  helm install awesome-webhook -n ${ns-with-spark-job}
```

## Usage/Examples

```bash
kubectl apply -f examples/spark-pi.yaml -n ${ns-with-spark-job}
```

## Documentation


Basic flow creation k8s resource with webhooks

![webhook.svg](examples/webhook.svg)


[Helm Chart docs](k8s/awesome-webhook/README.md)

## Run Locally (dev)

Pre-requisites
- Clone this repository
- Golang >= 1.22
- Openssl
- (Optional) Kind | Minikube | k3s with Helm

Generate dev certs

```bash
   ./dev/gencerts.sh   
```

Go to the project directory

```bash
  go build -o dist/webhook github.com/aagumin/sparkapp-k8s-webhook
```

Start hook

```bash
  ./dist/webhook --cfgPath=examples/feature-flag.yaml --tls-cert=localhost.crt --tls-key=localhost.key

```

Test working

```bash
  ./dev/mutate_test.sh | jq .response.patch | sed 's/^.//;s/.$//' | base64 --decode
```
