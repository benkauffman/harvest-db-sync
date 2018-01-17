FROM golang:1.9 AS build
ADD . /src
WORKDIR /src
RUN go get -d -v -t
RUN go test --cover -v ./... --run UnitTest
RUN go build -v -o harvest-db-sync

#TODO - using this in the interim
CMD ["go", "run", "main.go"]


#TODO - should be able to get this running with lightweight alpine image (reduce image size), but isn't working for me

#FROM alpine:3.4
#MAINTAINER 	Ben Kauffman <ben.kauffman@developmentnow.com>
#
##RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#
#CMD ["harvest-db-sync"]
#
#WORKDIR /usr/local/bin
#COPY --from=build /src/harvest-db-sync /usr/local/bin/harvest-db-sync
#COPY --from=build /src/config.json /usr/local/bin/config.json
#RUN chmod +x /usr/local/bin/harvest-db-sync