#!/bin/sh

# Wait for the PostgreSQL container to be ready
until nc -zv db 5432; do
    echo "Waiting for PostgreSQL container to be ready..."
    sleep 1
done

# Define the IP address of the API container
API_CONTAINER_IP=$(hostname -i)

# Define the entry to be added to pg_hba.conf
PG_HBA_ENTRY="host all all ${API_CONTAINER_IP}/32 md5"

# Path to pg_hba.conf file
PG_HBA_CONF="/var/lib/postgresql/data/pg_hba.conf"

# Add the entry to pg_hba.conf
echo "$PG_HBA_ENTRY" >> "$PG_HBA_CONF"

# Restart PostgreSQL service to apply the changes
service postgresql restart