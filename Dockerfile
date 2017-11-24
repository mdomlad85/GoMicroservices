FROM alpine:latest
MAINTAINER Marko Domladovac

COPY ./GoMicroservices /app/GoMicroservices
RUN chmod +x /app/GoMicroservices

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/GoMicroservices