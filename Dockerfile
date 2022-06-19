FROM golang:alpine

# Install minimum necessary dependencies,
RUN apk add --no-cache curl make git bash gcc libc-dev linux-headers wget
RUN wget https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && \
	mv jq-linux64 /usr/local/bin/jq && \
	chmod +x /usr/local/bin/jq

# Add source files
RUN mkdir /src
ADD . /src
WORKDIR /src

RUN go mod tidy
RUN make install

EXPOSE 26656 26657 1317 9090

CMD ["cantod"]
