#!/bin/bash

PROJECT_ID=smplverse

gcloud config set project $PROJECT_ID

gcloud services enable compute.googleapis.com

gcloud compute addresses create $PROJECT_ID-ip \
    --global

gcloud compute backend-services create $PROJECT_ID-backend \
    --global 

gcloud compute url-maps create $PROJECT_ID-urlmap \
    --default-service=$PROJECT_ID-backend

gcloud compute ssl-certificates create $PROJECT_ID-cert \
    --domains=$PROJECT_ID.xyz

gcloud compute target-https-proxies create $PROJECT_ID-https \
    --ssl-certificates=$PROJECT_ID-cert \
    --url-map=$PROJECT_ID-urlmap

gcloud compute forwarding-rules create $PROJECT_ID-load-balancer \
    --global \
    --target-https-proxy=$PROJECT_ID-https \
    --ports=443 \
    --address=$PROJECT_ID-ip

# add endpoint in us-central1
US_REGION=us-central1
gcloud compute network-endpoint-groups create $PROJECT_ID-neg-$US_REGION \
    --region=$US_REGION \
    --network-endpoint-type=SERVERLESS \
    --cloud-run-service=$PROJECT_ID-metadata-$US_REGION

gcloud compute backend-services add-backend $PROJECT_ID-backend \
    --global \
    --network-endpoint-group-region=$US_REGION \
    --network-endpoint-group=$PROJECT_ID-neg 

# add endpoint in europe-west1
EU_REGION=europe-west1
gcloud compute network-endpoint-groups create $PROJECT_ID-neg-$EU_REGION \
    --region=$EU_REGION \
    --network-endpoint-type=SERVERLESS \
    --cloud-run-service=$PROJECT_ID-metadata-$EU_REGION

gcloud compute backend-services add-backend $PROJECT_ID-backend \
    --global \
    --network-endpoint-group-region=$EU_REGION \
    --network-endpoint-group=$PROJECT_ID-neg
