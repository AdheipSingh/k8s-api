# k8s-api

- send JSON payload
```
{
"deployment": "prometheus"
}
```

- The above returns all the manifests deployed
- Incase the name is not prometheus it returns an error

- send GET request on /getsvc/prometheus to access endpoint
