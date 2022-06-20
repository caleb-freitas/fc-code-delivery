FROM golang:1.16

WORKDIR /go/src

# put go on the PATH environment variable 
ENV PATH="/go/bin:${PATH}"

# add librdkafka-dev package to work with kafka
RUN apt-get update && \
  apt-get install build-essential librdkafka-dev -y

# keep container up
CMD ["tail", "-f", "/dev/null"]
