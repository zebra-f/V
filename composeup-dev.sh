#!/bin/bash

set -e

if [ ! -f ./V-Django-Backend/compose/dev/django/.env.dev ]; then
        echo ".env.dev for Django doesn't exists."
        exit 1
fi

if [ ! -f ./.env ]; then
        echo ".env for Docker Compose doesn't exists."
        exit 2
fi

chmod +x ./V-Django-Backend/compose/dev/django/entrypoint.sh
chmod +x ./V-Django-Backend/compose/dev/django/celery/entrypoint.sh

sudo docker compose -f ./docker-compose.yaml down
sudo docker compose -f ./docker-compose.yaml up --build -d
