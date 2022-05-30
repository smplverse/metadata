# metadata

[![Go Report Card](https://goreportcard.com/badge/github.com/piotrostr/metadata)](https://goreportcard.com/report/github.com/piotrostr/metadata)
[![codecov](https://codecov.io/gh/piotrostr/metadata/branch/master/graph/badge.svg?token=bJwa6Sf4Z7)](https://codecov.io/gh/piotrostr/metadata)
[![CICD](https://github.com/piotrostr/metadata/actions/workflows/main.yml/badge.svg)](https://github.com/piotrostr/metadata/actions)

Simple and fast server that serves the NFTs written in Go.

## Usage

```bash
go run ./... --port [port]
```

Enpoints:

- `GET /:tokenId` to get a metadata entry
- `POST /:tokenId`
  to add an entry (requires `Authorization: [METADATA_API_KEY]` header)
- `GET /` for healthchecks, returns empty 200 OK

## Under the hood

Server uses Gin (gin-gonic) framework for serving metadata stored in Redis
database as JSON. Deployment is onto a cluster provisioned from Linode. The
format complies to the opensea.io metadata standards and employs NGINX ingress
with TLS certificates from lets-encrypt through cert-manager.io. The API packed
into a docker image of piotrostr/metadata is deployed aside a Redis image.

## Deployment

Note: The steps are user-specific since there is a number of variables like
certificate issuer email etc, the steps are more of guidelines rather than
walkthrough-tutorial or local setup.

### Provision the cluster

```sh
cd terraform/linode && terraform apply
terraform output kubeconfig | jq -r '@base64d' > ~/.kube/lke.yaml
export KUBECONFIG=~/.kube/lke.yaml
cd -
```

### Deploy to the cluster

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

5. Add secret (required env var for the metadata api)

   ```sh
   kubectl create secret generic metadata-api-key --from-literal METADATA_API_KEY=[secret]
   ```

6. Apply the configuration

   ```sh
   skaffold manifest.yaml
   ```

   I prefer `skaffold` to `kubectl` for applying deployments as it waits
   for them to stabilise and exits with error code 1 in case any container
   fails.
