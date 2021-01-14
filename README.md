# CS98 Hack-a-Thing

This is a basic CRUD API of a to-do list, written in Golang. An ultimate goal of this would be to deploy the API and the database into a Kubernetes Cluster.

The database is MongoDB, and can locally be run in a docker container like so:

```docker run --rm -it -v mongo:/data/db -p 27017:27017 mongo```

The API can be run in a separate terminal window using:

```go run main.go```

I have worked with MongoDB before using JavaScript and Mongoose. The new part for me here was writing the API in Go (and the package Mux), which has its advantages when deploying the API into a Kubernetes cluster.

The API supports standard CRUD operations. I have been interacting with it using cURL:

```bash
curl -X PUT -d "description=Here is a Todo Item" 127.0.0.1:8000/todo
curl 127.0.0.1:8000/todo
curl -X PATCH 127.0.0.1:8000/todo/XXXXXXXXXXXXXXXXXXXXXXXX
curl -X DELETE 127.0.0.1:8000/todo/XXXXXXXXXXXXXXXXXXXXXXXX
```

This project is based on a tutorial I followed here: https://keiran.scot/building-a-todo-api-with-golang-and-kubernetes-part-1-introduction/