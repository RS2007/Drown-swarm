# Wifi server and client

- The server code can be found in the server folder and client in client folder.

## Instructions to get it running

1. Make sure you have installed golang for your [platform](https://go.dev/doc/install)
2. To run the server:
   1. `cd server && go mod tidy && go run .`
3. To run the client
   1. `cd client && go mod tidy && go run .`
4. On the server enter the message `broadcast`,coordinates corresponding to the
   clientID(check the main.go file in server to see it) will be sent to each client

## Demo


https://github.com/RS2007/Drown-swarm/assets/83483297/828e332f-6897-4f1e-ac3b-f5d49e951444

