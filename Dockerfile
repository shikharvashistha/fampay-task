FROM golang:1.18
RUN /bin/bash -c "mkdir ${HOME}/app"
WORKDIR ${HOME}/app
COPY . .
EXPOSE 8087
ENTRYPOINT ["go"]
CMD ["run", "./cmd/main.go"]