# How to generate API client

1. Go to clients folder

## Identity

```bash
docker run --rm \
  -v ${PWD}:/local \
  -u 1000:1000 \
  openapitools/openapi-generator-cli generate \
  -i /local/identity-v1.json \
  -g go \
  -o /local/identityclient \
  --additional-properties=packageName=identityclient \
  --git-repo-id shifting/api/identityclient \
  --git-user-id winnix
```

## Integration

```bash
docker run --rm \
  -v ${PWD}:/local \
  -u 1000:1000 \
  openapitools/openapi-generator-cli generate \
  -i /local/integration-v1.json \
  -g go \
  -o /local/integrationclient \
  --additional-properties=packageName=integrationclient \
  --git-repo-id shifting/api/integrationclient \
  --git-user-id winnix
```

# How to add new apaleo API client:

1. Generate the client (see above)
2. Then run:

```bash
go mod edit -replace github.com/winnix/shifting/api/{API_NAME}client=./clients/{API_NAME}client
```

3. And then:

```bash
go get github.com/winnix/shifting/api/{API_NAME}client
```
