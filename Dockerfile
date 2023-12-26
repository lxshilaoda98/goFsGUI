#FROM scratch AS runner
FROM alpine

MAINTAINER "n1n1n1_owner@163.com"


WORKDIR /root



ADD ./goFsGUI ./goFsGUI

ADD ./templates ./templates


EXPOSE 1225

ENTRYPOINT ["./goFsGUI"]

