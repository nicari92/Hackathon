FROM golang:latest
COPY /app .
RUN go mod download
RUN go build -o /godocker
EXPOSE $SERVER_PORT
CMD [ "/godocker" ]