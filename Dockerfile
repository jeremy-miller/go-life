FROM 1.10-alpine
WORKDIR /go/src/app
COPY . .
RUN make setup
RUN make build
CMD ["make", "run"]
