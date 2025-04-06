# local-app

## khulnasoft-cli

All of the accessible commands can be listed with `khulnasoft --help` .

### Installing

1. Download the CLI for your platform and make it executable:

```bash
wget -O khulnasoft https://khulnasoft.com/static/bin/khulnasoft-cli-darwin-arm64
chmod u+x khulnasoft
```

2. Optionally, make it available globally. On macOS:

```bash
sudo mv khulnasoft /usr/local/bin/
```

### Usage

Start by logging in with `khulnasoft login`, which will also create a default context in the configuration file (`~/.khulnasoft/config.yaml`).

### Development

To develop the CLI with Khulnasoft, you can run it just like locally, but in Khulnasoft workspaces, a browser and a keyring are not available. To log in despite these limitations, provide a PAT via the `KHULNASOFT_TOKEN` environment variable, or use the `--token` flag with the login command.

#### In a Khulnasoft workspace

[![Open in Khulnasoft](https://www.khulnasoft.com/svg/open-in-khulnasoft.svg)](https://khulnasoft.com/#https://github.com/khulnasoft-com/khulnasoft)

You will have khulnasoft-cli ready as `khulnasoft` on any Workspace based on `https://github.com/khulnasoft-com/khulnasoft`.

```
# Reinstall `khulnasoft`
blazedock run components/local-app:install-cli

# Reinstall completion
blazedock run components/local-app:cli-completion
```

### Versioning and Release Management

The CLI is versioned independently of other Khulnasoft artifacts due to its auto-updating behaviour.
To create a new version that existing clients will consume increment the number in `version.txt`. Make sure to use semantic versioning. The minor version can be greater than 10, e.g. `0.342` is a valid version.

## local-app

**Beware**: this is very much work in progress and will likely break things.

### How to install

```
docker run --rm -it -v /tmp/dest:/out docker.io/khulnasoft/core-dev/build/local-app:<version>
```

### How to run

```
./local-app
```

### How to run in Khulnasoft against a dev-staging environment

```
cd components/local-app
BROWSER= KHULNASOFT_HOST=<URL-of-your-preview-env> go run main.go --mock-keyring run
```
