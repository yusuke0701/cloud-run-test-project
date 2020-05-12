#!/bin/sh -eux

# usage: ./scripts/deploy.sh PROJECT-ID IMAGE-NAME

cd `dirname $0`
: $1
: $2

cd ../
gcloud builds submit --tag gcr.io/$1/$2
gcloud run deploy --image gcr.io/$1/$2 --platform managed