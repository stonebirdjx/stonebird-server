FROM centos

LABEL maintainer="Stone Bird 1245863260@qq.com"

WORKDIR /usr/local/bin

COPY main /usr/local/bin

EXPOSE 6610

CMD ["/usr/local/bin/main"]