#!/bin/bash

gcloud config set project smplverse

gcloud services enable \
  cloudbuild.googleapis.com \
  run.googleapis.com \
  storage.googleapis.com \
  redis.googleapis.com

gcloud builds submit .
