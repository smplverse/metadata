# metadata

[![Go Report Card](https://goreportcard.com/badge/github.com/piotrostr/metadata)](https://goreportcard.com/report/github.com/piotrostr/metadata)
[![codecov](https://codecov.io/gh/piotrostr/metadata/branch/master/graph/badge.svg?token=bJwa6Sf4Z7)](https://codecov.io/gh/piotrostr/metadata)
[![CICD](https://github.com/piotrostr/metadata/actions/workflows/main.yml/badge.svg)](https://github.com/piotrostr/metadata/actions)

## Deployment

Note: The steps are user-specific since there is a number of variables like
certificate issuer email etc, the steps are more of guidelines rather than
walkthrough-tutorial or local setup.

### Provision the cluster

```sh
cd terraform && terraform apply
terraform output kubeconfig | jq -r '@base64d' > ~/.kube/lke.yaml
export KUBECONFIG=~/.kube/lke.yaml
cd ..
```

### Install `ingress-nginx` and get a CA-verified certificate

1. Install ingress

   ```sh
   helm upgrade --install ingress-nginx ingress-nginx \
     --repo https://kubernetes.github.io/ingress-nginx \
     --namespace ingress-nginx --create-namespace
   ```

2. To get the IPv4 of the Ingress

   ```sh
   kubectl get services \
     --namespace ingress-nginx \
     -o wide \
     -w \
     ingress-nginx-controller
   ```

3. Add A-Record to the domain from the Ingress manifest (in this case
   `metadata.smplvserse.xyz`) pointing to the IPv4 of the Ingress

4. Install cert-manager

   ```sh
   kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.8.0/cert-manager.yaml
   ```

5. Apply the configuration

   ```sh
   kubectl apply -f manifest.yaml
   ```
