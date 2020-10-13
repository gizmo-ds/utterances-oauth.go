FROM alpine:latest
RUN mkdir -p /home/app
WORKDIR /home/app
COPY goapp /home/app
ENTRYPOINT /home/app/goapp
EXPOSE 5000