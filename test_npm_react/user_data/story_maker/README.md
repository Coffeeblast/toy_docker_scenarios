# Backend - python + fastapi ```./backend```
* runs on fastapi
* use run.sh to ran dev server

# Backend - go + gin ```./backend_go```
* [reference go tutorial](https://go.dev/doc/tutorial/web-service-gin)
* ```go mod init example/story_maker/backend```
* create settings.json so that go extension does not delete imports how it likes
* ```go get .``` -- download dependencies after adding new imports


# Frontend - node + react ```./frontend
* runs on react / vite

## project setup
* ```npm create vite@latest frontend -- --template react```
* ```cd frontend```
* ```npm install```
* ```npm run dev -- --host```