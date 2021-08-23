#  Nodejs Docker Basic Demo
## Getting Started
This examples shows us to how to Dockerize basic Nodejs Express app. 

## Step by Step
### 1. Create simple nodejs
Lets create simple nodejs app using express. There will be only one endpoint to understanding Dockerizing the nodejs app.
```shell
#set up new npm project 
npm init -y
#install express
npm install express
 ```
### 2. Create index.js to write our http server
```js
const express = require("express");
const app = express();

app.get('/', (req, res) => {
res.json({ message: 'Hello Docker' });
});
const port = process.env.port || 3000;
app.listen(port, () => {
console.log(`Server listening on ${port}...`);
});
 ```
 ### 3. Create Dockerfile
 ```Dockerfile
 FROM node:14
 #Create app directory
WORKDIR /app
#Copy package.json and package-lock.json file to ./app
COPY package*.json ./
#install all depencies
RUN npm install
COPY . .
#Create environment variable(s)
ENV PORT=3000
#Expose the 3000 port to listen
EXPOSE 3000
#Run the app
CMD node index.js
  ```
### 4. Build app image
```shell
docker build --tag enescang/nodejs-basic-demo:1.0 .
```
Then lets get list of images with:
```shell
docker images
```
Result must be like this:
| REPOSITORY                 | TAG | IMAGE ID     | CREATED        | SIZE  |
|----------------------------|-----|--------------|----------------|-------|
| enescang/nodejs-basic-demo | 1.0 | 1724dc236ffc | 21 seconds ago | 947MB |

### 5. Run container
```shell
docker run 1724dc236ffc
```
If there is no any problem our app will be run correctly but we can not access http://localhost:3000 for now. Let's make small changes.
```shell
docker run --publish LOCAL_PORT:CONTAINER_PORT 1724dc236ffc
```
So, use any port for LOCAL_PORT but CONTAINER_PORT must be same what you decide on code.
```shell
docker run --publish 4000:3000 1724dc236ffc
```
### 6. Stop the container
To get list of container docker have a small command
```shell
docker ps
```
| CONTAINER ID | IMAGE        | NAME            |
|--------------|--------------|-----------------|
| f3e47b41459d | 1724dc236ffc | vigorous_curran |

We need CONTAINER ID or NAME to stop
```shell
docker stop f3e47b41459d
```
### 7. Restart the container
If we use docker run it will be create new container but we want to restart a container so again we need to CONTAINER ID or NAME
```shell
docker restart f3e47b41459d
```
> Written with [StackEdit](https://stackedit.io/).