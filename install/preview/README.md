# Khulnasoft Local Preview

This repo helps users to try out and preview self-hosted Khulnasoft **locally** without all the things
needed for a production instance. The aim is to provide an installation mechanism as minimal and
simple as possible.

## Installation

```bash
docker run --privileged --name khulnasoft --rm -it -v /tmp/khulnasoft:/var/khulnasoft docker.io/khulnasoft/core-dev/build/preview-install
```

Once the above command starts running and the pods are ready (can be checked by running `docker exec khulnasoft kubectl get pods`),
The URL to access your khulnasoft instance can be retrieved by running

```
docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' khulnasoft |  sed -r 's/[.]+/-/g' | sed 's/$/.nip.io/g'
```

[nip.io](https://nip.io/) is just wildcard DNS for local addresses, So all off this is local, and cannot be accessed over the internet.

As the `self-hosted` instance is self-signed, The root certificate to upload into your browser trust store to access the URL is available at
`/tmp/khulnasoft/khulnasoft-ca.crt`.

## Known Issues

- Prebuilds don't work as they require webhooks support over the internet.
