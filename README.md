# Simulation


## create a docker image
```
docker build -t simimage:1.0 .
```

## create deployment
```
kubectl create -f deployment.yaml
```

## create service

```
kubectl create -f service.yaml
```

## do post

```
curl --header "Content-Type: application/json" --request POST --data '{"Pid": "001","Pname":"Vishnu"}' host01:30080/add
```
