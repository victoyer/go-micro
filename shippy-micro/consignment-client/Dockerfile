FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app
COPY consignment-client /app/consignment-client
COPY consignment.json /app/consignment.json
CMD ["./consignment-client"]
