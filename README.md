# k8s-configmap-watcher

# Archival Note

⚠️ This repo (and the corresponding container image) is unmaintained. Please use https://github.com/kiwigrid/k8s-sidecar instead which is a maintained drop-in-replacement.

# What?

This is a docker container intended to run inside a kubernetes cluster to collect config maps with a specified label and store the included files in an local folder. It can also send a html request to a specified URL after a configmap change. The main target is to be run as a sidecar container to supply an application with information from the cluster. The contained Go programm is working with the Kubernetes API 1.10

# How?

Run the container created by this repo together you application in an single pod with a shared volume. Specify which label should be monitored and where the files should be stored.
By adding additional env variables the container can send a html request to specified URL.

# Features

- Extract files from config maps
- Filter based on label
- Update on change of configmap

# Usage

Example for a simple deployment can be found in `example.yaml`. Depending on the cluster setup you have to grant yourself admin rights first: `kubectl create clusterrolebinding cluster-admin-binding   --clusterrole cluster-admin   --user $(gcloud config get-value account)`

## Configuration Environment Variables

- `LABEL` 
  - description: Label that should be used for filtering
  - required: true
  - type: string

- `FOLDER`
  - description: Folder where the files should be placed
  - required: true
  - type: string

- `NAMESPACE`
  - description: If specified, the sidecar will search for config-maps inside this namespace. Otherwise the namespace in which the sidecar is running will be used. It's also possible to specify `ALL` to search in all namespaces.
  - required: false
  - type: string

- `METHOD`
  - description: If `METHOD` is set with `LIST`, the sidecar will just list config-maps and exit. Default is watch.
  - required: false
  - type: string
  
- `REQ_URL`
  - description: URL to which send a request after a configmap got reloaded
  - required: false
  - type: URI

- `REQ_METHOD`
  - description: Request method GET(default) or POST
  - required: false
  - type: string

- `REQ_PAYLOAD`
  - description: If you use POST you can also provide json payload
  - required: false
  - type: json
