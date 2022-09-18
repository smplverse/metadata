#!/bin/bash

gcloud config set project smplverse

gcloud services enable \
  cloudbuild.googleapis.com \
  run.googleapis.com \
  storage.googleapis.com \
  redis.googleapis.com

gcloud builds submit .

gcloud run deploy smplverse-metadata \
  --image gcr.io/smplverse/metadata \
  --region us-central1 \
  --allow-unauthenticated
