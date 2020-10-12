# Simulation

## create image
```
docker build -t simimage:1.0 .
```

## create repo
```
docker run -d -p 5000:5000 --restart=always --name registry registry:2
```

## add to repo
```
docker tag simimage:1.0 localhost:5000/sim
```

## list images
```
docker image ls
```

## start kubernetis and install helm
```
launch.sh
curl -LO https://storage.googleapis.com/kubernetes-helm/helm-v2.8.2-linux-amd64.tar.gz
tar -xvf helm-v2.8.2-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/
helm init
helm repo update
```

## install new chart

```
helm install --name example ./mychart
```

## delete existing deployment and service
```
kubectl delete deployment example-mychart
kubectl delete svc example-mychart
helm del --purge example
```
