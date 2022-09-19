# smplverse/metadata

Simple metadata server that loads the metadata object into memory and serves
based on `/:tokenID`. 

Deployed using cloud run, CI/CD using cloud build.

# Setup

To create two regional deployments in us-central1 and
europe-west1:

```sh
./hack/deploy.sh
```

To provision the load balancer (create backend service,
URL-maps, SSL cert, endpoint groups, forwarding rules):

```sh
./hack/setup-load-balancing.sh
```

After the load balancer provisioning completes, the A-record needs to be added
to the DNS configuration (steps might differ based on the domain provider). 

The IP of the load balancer can be grabbed using:

```sh
gcloud compute addresses describe $PROJECT_ID-ip \
    --global \
    --format 'value(address)'
```
