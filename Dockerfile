FROM ubuntu:latest
LABEL authors="dell"

ENTRYPOINT ["top", "-b"]