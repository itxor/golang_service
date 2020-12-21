FROM golang:latest

RUN useradd -ms /bin/bash admin

COPY ./ /app
WORKDIR /app

RUN chown -R admin:admin /app \
    && chmod 755 /app

USER admin 