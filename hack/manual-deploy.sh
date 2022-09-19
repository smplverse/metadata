#!/bin/bash

gcloud run deploy smplverse-metadata-us-central1 \
  --image gcr.io/smplverse/metadata \
  --region us-central1 \
  --allow-unauthenticated \
  --async

gcloud run deploy smplverse-metadata-europe-west1 \
  --image gcr.io/smplverse/metadata \
  --region europe-west1 \
  --allow-unauthenticated \
  --async
