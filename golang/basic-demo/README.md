

#  Golang Docker Basic Demo
## Getting Started
This examples shows us to how to Dockerize basic Go http server app. 

## Step by Step
### 1. Create simple golang app
Lets create simple golang app using **net/http** package. There will be only one endpoint to understanding Dockerizing.
```shell
#set up new npm project 
go mod init enescang/docker-samples/golang/basic-demo
#install the github.com/joho/godotenv to use .env file
go get github.com/joho/godotenv
 ```
### 2. Create main.go to write our http server
```golang
package main
import  (
"fmt"
"net/http"
"os"
"github.com/joho/godotenv"
)
func main() {
godotenv.Load()
port := "3000"
if v := os.Getenv("PORT"); v != "" {
port = os.Getenv("PORT")
}
http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
rw.Header().Add("Content-Type",  "application/json")
rw.Write([]byte("{\"message\":\"Hello Docker\"}"))
})
fmt.Println("Server listening on ", port)
http.ListenAndServe(":"+port,  nil)
}
 ```
 ### 3. Create Dockerfile
 ```Dockerfile
 FROM golang:1.16-alpine
 #Create app directory
WORKDIR /app

#Copy go.mod and go.sum to ./app
COPY go.mod ./
COPY go.sum ./

#get all depencies
RUN go mod download
#Copy all go files to ./
COPY *.go ./
#Build the app as a "hello-docker"
RUN go build -o /hello-docker
#Set env variable(s)
ENV PORT=3000

#Expose the port to listen
EXPOSE 3000
#Run the app
CMD [ "/hello-docker" ]
  ```
### 4. Build app image
```shell
docker build --tag enescang/golang-basic-demo:1.0 .
```
Then lets get list of images with:
```shell
docker images
```
Result must be like this:
| REPOSITORY                 | TAG | IMAGE ID     | CREATED        | SIZE  |
|----------------------------|-----|--------------|----------------|-------|
| enescang/golang-basic-demo | 1.0 | 5dfba54b592c | 10 seconds ago | 308MB |

### 5. Run container
```shell
docker run 1724dc236ffc
```
If there is no any problem our app will be run correctly but we can not access http://localhost:3000 for now. Let's make small changes.
```shell
docker run --publish LOCAL_PORT:CONTAINER_PORT 5dfba54b592c
```
So, use any port for LOCAL_PORT but CONTAINER_PORT must be same what you decide on code.
```shell
docker run --publish 4000:3000 5dfba54b592c
```
### 6. Stop the container
To get list of container docker have a small command
```shell
docker ps
```
| CONTAINER ID | IMAGE        | NAME            |
|--------------|--------------|-----------------|
| b869279e6e92 | 5dfba54b592c| exciting_hugle |

We need CONTAINER ID or NAME to stop
```shell
docker stop b869279e6e92
```
### 7. Restart the container
If we use docker run it will be create new container but we want to restart a container so again we need to CONTAINER ID or NAME
```shell
docker restart b869279e6e92
```
> Written with [StackEdit](https://stackedit.io/).
