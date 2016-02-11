FROM scratch

WORKDIR /app

COPY ./s-db /app/run
COPY ./keys /app/keys

CMD ["/app/run", "http"]
