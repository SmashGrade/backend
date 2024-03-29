# Smashgrade Backend

This is the backend service for Smashgrade.

## Configuration

Configuration of the container can be managed either with the configuration file (see [defaults.yaml](defaults.yaml)) or environment variables. The environment mode (e.g dev or prod) has to be set with environment variables in the container.

```
# Required environment variable (default: dev)
ENV = prod

API_LOG_LEVEL=info
API_HOST=0.0.0.0
API_PORT=9000
# Determines if the backend connects to the database automatically when starting
API_CONNECT=true
# Determines if the backend automatically migrates the database models
API_AUTO_MIGRATE=true
# Prefils database with test data if true
API_MOCK_DATA=false
# Enables authentification with Microsoft Entra ID in the backend (default enabled on prod, disabled on dev)
API_AUTH_ENABLED=false
# Sets the connection string for the database, supported: sqlite://, postgres:// or mysql://
API_DB_CONNECTION_STR=sqlite:///data/data.db
# Sets the Oauth discovery URL for Microsoft Entra ID in the backend
API_AUTH_OAUTH_KEY_DISCOVERY_URL=https://login.microsoftonline.com/common/discovery/keys
```

## Running the container

To find documentation run the docker image locally with:

```
docker run --name smashgrade-backend -e ENV=dev -p 9000:9000 ghcr.io/smashgrade/backend:dev
```

## Mounting your configuration

1. Run the docker container with -v ${PWD}/config:/app/config/ to mount the config folder in the current directory to the container
2. The configuration file gets created
3. You can update the values in the configuration file and restart the container

## Adding a new user e.g Kursadmin

To add a new user manually to the database you can use the following steps:

1. Run the server at least once to generate the database, roles and configuration file
2. Run the command ``` smashgrade add -n "Peter Mueller" -e peter.mueller@hftm.ch -r Kursadministrator ```
3. The user will get created with the role that was assigned to him

## Documentation

You can now access the backends documentation with http://localhost:9000/docs/

Keep in mind that the endpoints in the swagger doc point to api.smashgrade.ch
