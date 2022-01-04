FROM golang:1.14-alpine as builder
WORKDIR /opt/chamberlain
ENV GOPROXY=https://goproxy.cn GIN_MODE=release
COPY . .
RUN go mod download &&\
    GOOS=linux CGO_ENABLED=0 go build -o chamberlain_mgmt -ldflags "-s -w" -i chamberlain_mgmt &&\
    chmod a+x ./resources/upx &&\
    link ./resources/upx /usr/local/bin/upx &&\
    upx -o chamberlain chamberlain_mgmt &&\
    rm -rf /opt/chamberlain/chamberlain_mgmt /opt/chamberlain/chamberlain_mgmt.go

#FROM scratch as runner
FROM scratch as runner
COPY --from=builder /opt/chamberlain/chamberlain* /
EXPOSE 8080
CMD ["/chamberlain"]