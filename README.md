# This is fine.

All of your kube or docker blowing up? This is fine.

![This is fine.](./tif.png)


Run on kubernetes:

`kubectl run -it --rm --image=ghcr.io/pdevine/thisisfine tif`

Run on docker:

`docker run -it --rm ghcr.io/pdevine/thisisfine`


## Building the image manually

### Building in Kubernetes

Use [BuildKit CLI for Kubectl](https://github.com/vmware-tanzu/buildkit-cli-for-kubectl) with the command:

`kubectl build -t thisisfine ./`

or, you can build a multi-arch image which cross-compiles for each platform. You'll need to create a registry secret
in kubernetes and push to a registry to make this work correctly.

```
read -s REG_SECRET
kubectl create secret docker-registry mysecret --docker-server='someregistry.io' --docker-username=tifdog --docker-password=$REG_SECRET
kubectl build ./ -t someregistry.io/stuff/thisisfine:latest -f Dockerfile.cross --registry-secret my-secret --platform=linux/386,linux/amd64,linux/arm/v6,linux/arm/v7,linux/arm64,windows/amd64 --push
```

### Building in Docker

To build a single image in Linux:

`docker build -t thisisfine ./`


## Acknowledgements

 * Based upon KC Green's wonderful comic strip [Gunshow](http://gunshowcomic.com/648)

 * Dog pixel art inspired by [PH](https://reddit.com/u/ph145) on Reddit

 * Animated with [Go-AsciiSprite](https://github.com/pdevine/go-asciisprite)


## FAQ

Q: Why does this look like crap on my Mac?<br>
A: Use iTerm2 instead of macOS's built-in Terminal app. Terminal screws up all of the line spacing.

