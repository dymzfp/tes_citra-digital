FROM scratch

MAINTAINER Dimas dfebri446@gmail.com

ADD main ./

ARG appname=citra_test
ARG http_port=1234

ENTRYPOINT ["/main"]

EXPOSE $http_port