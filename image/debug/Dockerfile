# Compile Binary Step
FROM golang:1.15.8 as builder

WORKDIR /workspace
COPY . .

RUN apt-get update && apt-get install -y \
    protobuf-compiler

RUN go get -u github.com/golang/protobuf/protoc-gen-go@latest && \
    go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Adding the grpc_health_probe
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.6 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

RUN make build

# Service Container
FROM ubuntu:latest

COPY --from=builder /workspace/bin/project ./project

# Add GRPC Health Check from previous step
COPY --from=builder /bin/grpc_health_probe ./grpc_health_probe

ENTRYPOINT [ "./project;./grpc_health_probe -addr=localhost:50051" ]
# Do not use CMD for scratch container
# CMD ["/grpc_health_probe -addr=localhost:50051"]