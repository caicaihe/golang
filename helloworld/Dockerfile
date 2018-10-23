FROM daocloud.io/library/centos
USER root
RUN mkdir /app
ADD . /app
EXPOSE 80
WORKDIR /app/helloworld
CMD ["chmod +x helloworld"]
CMD ["./helloworld"]
