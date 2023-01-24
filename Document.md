## Instructions to Run the service in docker

1. git clone git@github.com:kalpit-sharma-dev/parkinglot-service.git (clone the repo)
2. sudo docker-compose build (build image)
3. sudo docker-compose up
4. docker ps -a
5. docker exec -it containerid bash (go inside the container)
6. Follow documentation in openapi.yml to test the API's