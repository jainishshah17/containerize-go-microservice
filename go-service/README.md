# Helm Chart for Go-service

### Configure docker registry credentials in K8S to pull docker image from Artifactory docker registry

## Create a Secret named regsecret:
```
kubectl create secret docker-registry regsecret --docker-server=<your-registry-server> --docker-username=<your-name> --docker-password=<your-pword> --docker-email=<your-email>
```
##### where:
```
<your-registry-server> is your Private Docker Registry FQDN.
<your-name> is your Docker username.
<your-pword> is your Docker password.
<your-email> is your Docker email.
```

##### Understanding your Secret

To understand what’s in the Secret you just created, start by viewing the Secret in YAML format:
```
kubectl get secret regsecret --output=yaml
```

## Installing the Chart

To install the chart with the release name `my-release`:

```console
$ helm install --name my-release chart/
```

The command deploys go-service on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `my-release` deployment:

```console
$ helm delete my-release --purge
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following tables lists the configurable parameters of the go-service chart and their default values.

|           Parameter                |             Description             |                        Default                            |
|------------------------------------|-------------------------------------|-----------------------------------------------------------|
| `image.repository`                 | Node-version image                  | `$ART_DOCKER_REPO/containerize-go-service:{tag}`          |
| `image.pullPolicy`                 | Image pull policy                   | `Always`                                                  |
| `image.tag`                        | Tag of docker image                 | `latest`                                                  |
| `imagePullSecrets`                 | Credentials of private docker repo  | ` `                                                       |
| `service.type`                     | Kubernetes Service type             | `LoadBalancer`                                            |
| `service.internalPort`             | Internal Service Port               | `8080`                                                    |
| `service.externalPort`             | External Service Port               | `80`                                                      |
 
Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example,

```console
$ helm install --name my-release --set image.tag=27 chart/
```

Alternatively, a YAML file that specifies the values for the above parameters can be provided while installing the chart. For example,

```console
$ helm install --name my-release -f values.yaml chart/
```

## Upgrade chart
```console
$ helm upgrade --name my-release --set image.tag=28 chart/
```

