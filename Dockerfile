FROM golang:alpine as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w -s -extldflags "-static"' -o main .


#FROM scratch
FROM alpine
#FROM gcr.io/distroless/base

#RUN adduser -S -D -H -h /app appuser
#USER appuser



COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]