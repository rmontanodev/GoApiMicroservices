
# docker-compose build
# docker-compose up -d

To run test with cover, enter to any service(auction f.e) and run
# go test ./... -coverprofile=coverage.out

There is a script to run cover + cover in html, enter to any service(auction f.e) and run
# ./coverage.sh
The idea of this project is to practice Go, testing, and communication between services.
