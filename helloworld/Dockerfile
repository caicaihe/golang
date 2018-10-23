FROM daocloud.io/library/centos
USER root
RUN mkdir /app
WORKDIR /app
ADD . /app
EXPOSE 80
CMD ["chmod +x helloworld"]
CMD ["./helloworld"]