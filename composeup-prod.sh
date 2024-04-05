#!/bin/bash

set -e

if [ ! -f ./V-Django-Backend/compose/prod/django/.env.prod ]; then
        echo ".env.prod for Django doesn't exists."
        exit 1
fi

if [ ! -f ./.env ]; then
        echo ".env for Docker Compose doesn't exists."
        exit 2
fi

chmod +x ./V-Django-Backend/compose/prod/django/entrypoint.sh
chmod +x ./V-Django-Backend/compose/prod/django/celery/beat/entrypoint.sh
chmod +x ./V-Django-Backend/compose/prod/django/celery/workers/entrypoint.sh

sudo docker compose -f ./docker-compose.prod.yaml down
sudo docker compose -f ./docker-compose.prod.yaml up --build -d
