FROM ubuntu:latest
LABEL authors="sergei"

ENTRYPOINT ["top", "-b"]