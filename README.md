# oneprog-oneans

## Description
Server side for each question and answer Web App.<br>
[Front end](https://github.com/higashi000/oneprog-oneans-front)

## Support Platforms
- Linux

## Language and Library
- Language
  - go1.14.1 linux/amd64
- Library
  - [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)
  - [gin](https://github.com/gin-gonic/gin)
  - [cors](https://github.com/gin-contrib/cors)

## Usage
```
$ git clone https://github.com/higashi000/oneprog-oneans

$ cd oneprog-oneans

$ sudo docker-compose up -d

$ go build

$ ./oneprog-oneans
```

Afterwards, please run [oneprog-oneans-front](https://github.com/higashi000/oneprog-oneans-front), and access to [localhost:4000](http://localhost:4000).
