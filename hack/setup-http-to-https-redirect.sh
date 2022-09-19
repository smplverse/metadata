#!/bin/bash

PROJECT_ID=smplverse
HTTP_URL_MAP=lb-map-http

gcloud compute url-maps validate \
  --source $HTTP_URL_MAP.yaml

gcloud compute url-maps import $HTTP_URL_MAP \
   --source $HTTP_URL_MAP.yaml \
   --global

gcloud compute target-http-proxies create http-lb-proxy \
   --url-map=$HTTP_URL_MAP \
   --global 

gcloud compute forwarding-rules create http-content-rule \
   --load-balancing-scheme=EXTERNAL \
   --network-tier=PREMIUM \
   --address=$PROJECT_ID-ip \
   --global \
   --target-http-proxy=http-lb-proxy \
   --ports=80
