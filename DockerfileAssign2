#creating multiple stage docler
#sBuild stage
FROM golang:1.21.1-alpine as builder

#set the working directory
WORKDIR /app

COPY . .

#build the app
RUN go build -o myapp

#use the smaller image to reduce the sozw of the final container 
FROM  alpine:latest

#set working directory
WORKDIR /root/
#copy the binary from the builder stage
COPY --from=builder /app/myapp .

#Execute the app
CMD ["./myapp"]