A Go HTTP Echo server in a scratch based Docker container for Kubernetes
===========================================================================

Potential Prereqs
--------------
* Get the certs for the scratch container
```
curl -o ca-certificates.crt https://curl.haxx.se/ca/cacert.pem
```

Docker
--------------
Build the server in a Docker image, then tag it and push it to Docker Hub

***Note***: statically build the libraries into the executable for portability

* Build the server in a Docker image
    * ***scratch***
    ```
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo
    ```

    * golang:onbuild (Alternative to scratch)
    ```
    GOOS=linux GOARCH=amd64 go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
    ```

* Build the Docker image and tag it
```
docker build -t http-echo:1.0.0 .
```

* Test docker
```
docker run --publish 6060:8080 --name test --rm http-echo:1.0.0 
curl localhost:6060
```

* Tag and deploy to docker hub
```
# docker login
docker tag http-echo:1.0.0 foo/http-echo:1.0.0
docker push foo/http-echo:1.0.0
```

GKE
--------------
* Log in and get kube credentials
```
gcloud auth application-default login
gcloud container clusters get-credentials k0 # Assuming we have a k0 cluster
```

* Provision a Kubernetes Cluster with GKE using gcloud
    * Only need this once (obviously)
```
cd $GOPATH/src/github/com/udacity/ud615/kubernetes
gcloud container clusters create k0

gcloud container clusters list
```

* Create http-echo deployment and service
```
kubectl create -f http-echo-deployment.yaml
kubectl create -f http-echo-service.yaml
```

* Update http-echo (after changing something)
```
kubectl replace -f http-echo-deployment.yaml
kubectl replace -f http-echo-service.yaml
```

* View the deployment
```
kubectl describe svc http-echo
```

* Get the IP address
```
kubectl get services http-echo
```

* Cleanup
```
kubectl delete services http-echo
kubectl delete deployments http-echo
```
