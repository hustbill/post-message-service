# Review Service API in Golang 


├── conf  
│   └── app.conf  
├── controllers  
│   └── reviewController.go  
├── main.go  
├── models  
├── routers  
│   └── router.go  
├── tests  
│   └── default_test.go  
└── README.md


## Use Prometheus

### Installing 
$ mkdir $HOME/work  
$ export GOPATH= $HOME/work  
$ cd $GOPATH  
$ go get github.com/prometheus/client_golang/prometheus  

### Build  
$ mkdir bin  
$ mkdir src  
$ mkdir pkg     
$ export PATH=$PATH:$GOPATH/bin    

$ mkdir -p $GOPATH/src/github.com/hustbill/  
$ git clone https://github.com/hustbill/review-service.git  
$ cd review-service  
$ cd go  
$ go install  #install other dependencies  
$ go build    # produce an executable binary  


For build a single common library   
$ go build github.com/hustbill/review-service/go/models  
  
Note: it creates a pakcage which name as models.a in $GOPATH/pkg/darwin_amd64/github.com/hustbill/review-service/go  

### Run 

1. start Prometheus Server  
$ cd review-service  && mkdir prometheus && cd prometheus
 Download the latest release of Prometheus for your platform, then extract and run it   
$ tar xvfz prometheus-*.tar.gz  
$ cp ../prometheus.yml.sample ./promethesual.yml  
$ ./prometheus prometheus.yml  

2. start Client with our APIs  
$ cd review-service/go  
$ go run *.go   
or     
$ $GOPATH/bin/main  

our review services will be started on port 8090  
  
### View by Web Browser
1. server side 
http://localhost:9090/

2. client side
http://localhost:8090/metrics
