FROM golang

# copy root dir to /app/auth in docker image
COPY . /app

# setting work directory, so all the subsequent commands can be run on this dir
WORKDIR /app

# build authserver code and set it's output binary in dir as /out/auth
RUN go build -o ./out .

# command to execute binaries
CMD ["./out"]