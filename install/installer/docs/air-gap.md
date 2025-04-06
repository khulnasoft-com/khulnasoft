# Installing Khulnasoft in an air-gap network with the Khulnasoft Installer

## Mirror Khulnasoft Images

You need a registry that is reachable in your network. Add the domain of your registry to the Khulnasoft config `khulnasoft.config.yaml` like this:
```yaml
repository: your-registry.example.com
```

The command `khulnasoft-installer mirror list` gives you a list of all images needed by Khulnasoft. You can run the following code to pull the needed images and push them to your registry:

```
for row in $(khulnasoft-installer mirror list --config khulnasoft.config.yaml | jq -c '.[]'); do
    original=$(echo $row | jq -r '.original')
    target=$(echo $row | jq -r '.target')

    docker pull $original
    docker tag $original $target
    docker push $target
done
```

## Install Khulnasoft in Air-Gap Mode

To install Khulnasoft in an air-gap network, you need to configure the repository of the images needed by Khulnasoft (see previous step). Add this to your Khulnasoft config:

```yaml
repository: your-registry.example.com
```

That's it. Run the following commands as usual and Khulnasoft fetches the images from your registry and does not need internet access to operate:

```
khulnasoft-installer render --config khulnasoft.config.yaml > khulnasoft.yaml
kubectl apply -f khulnasoft.yaml
```
