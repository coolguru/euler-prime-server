Euler-Prime server in golang using echo-labstack framework

## Overview

Module to find the Xth Y digit prime number in the expansion of Euler's number.

## Quick Start
Euler-Prime server is developed and tested using Go 1.6.x and 1.7.x

### Build & Run Docker image
```
git clone https://github.com/coolguru/euler-prime-server.git
docker build -t euler-server .
docker run -it -p 3600:3600 euler-server
```

### cURL commands
Part 1 - Xth Y digit prime number in the expansion of Euler's number.
```
curl -X GET "http://DOCKERIP:3600/eulerprime?x=5&y=6"
```

Part 2 - Upload csv file (csv file under data folder)
```
curl -X POST -H "Content-Type: multipart/form-data" -F "file=@" "http://DOCKERIP:3600/eulerprimeupload"
```

### Tasks TODO

|Description |
| :---: |
|Add tests to euler-server|
|Implement web interface for input x and y; and upload file functionality|
|Better error handling|
|Better documentation|
