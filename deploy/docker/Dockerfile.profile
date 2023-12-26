FROM centos
ADD https://github.com/yxm0513/go-micro-service/releases/download/v1.0.1/go-micro-service-v1.0.1-linux-amd64.tar.gz .
RUN tar -xzf go-micro-service-v1.0.1-linux-amd64.tar.gz -C .

EXPOSE 8083 6063
ENTRYPOINT ["./go-micro-service-v1.0.1-linux-amd64/profile"]
