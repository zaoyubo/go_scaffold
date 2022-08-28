# Quick Start
1. go mod tidy
2. Install service-center
```
docker pull servicecomb/service-center
docker run -d -p 30100:30100 servicecomb/service-center
```
3. make run_grpc