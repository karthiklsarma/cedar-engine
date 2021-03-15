# cedar-engine

[Contributors |](https://github.com/karthiklsarma/cedar-engine/graphs/contributors)
[Forks |](https://github.com/karthiklsarma/cedar-engine/network/members)
[Issues |](https://github.com/karthiklsarma/cedar-engine/issues)
[MIT License |](https://github.com/karthiklsarma/cedar-engine/blob/main/LICENSE)

## To build and run project on localhost:
From Root Directory:
### To build
* Execute 
> go build -o /bin/cedar
### To run
* Execute 
> /bin/cedar

## To build and run project on docker container:
From Root Directory:
### To build
*Execute
> docker build . -t cedar:latest
### To run
* Execute
> docker run -d -p 8080:8080 <IMAGE ID from previous step>

