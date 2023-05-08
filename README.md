# GRPC client and server for (EBarimt PosAPI with Go lang)



## Powered By Developer team of Lambda Platform.


### [EBarimt PosAPI with Go lang](https://github.com/lambda-platform/ebarimt)


### [PosAPI lib installer on Linux](https://github.com/lambda-platform/ebarimt-lib-installer)


protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false posapi.proto


docker build --platform linux/amd64 --build-arg ARCH=x64 -t ebarimtgrpc .

docker tag ebarimt:latest munkhaltai/ebarimtgrpc:latest
docker push munkhaltai/ebarimtgrpc:latest
