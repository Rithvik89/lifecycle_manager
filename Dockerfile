FROM golang:latest
WORKDIR /app
COPY . .
COPY .env ./
RUN make lcm
EXPOSE 4000
CMD ["./app/api"]