#!/bin/bash

gcloud config set project smplverse

gcloud services enable \
  run.googleapis.com \
  storage.googleapis.com

gcloud run deploy smplverse-metadata \
  --image gcr.io/smplverse/metadata \
  --region us-central1 \
  --allow-unauthenticated
