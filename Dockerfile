FROM golang:1.22 AS build

WORKDIR /app

ARG TARGETARCH

ENV ENV=prod
ENV PATH="${PATH}:/app/.bin"
ENV TARGETARCH=${TARGETARCH}
ENV GOARCH=${TARGETARCH}
ENV GOOS=linux

COPY . .

RUN mkdir /app/.bin
RUN ARCH=${TARGETARCH} && \
    if [ "$TARGETARCH" = "amd64" ]; then \
    ARCH="x64"; \
    fi && \
    curl -sLO "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-$ARCH" && \
    mv tailwindcss-linux-${ARCH} /app/.bin/tailwindcss
RUN chmod +x /app/.bin/tailwindcss

RUN go get github.com/a-h/templ/runtime
RUN make tidy build

FROM golang:1.22

WORKDIR /app

COPY --from=build /app/.out/myapp .
COPY --from=build /app/public ./public

ENV ENV=prod

EXPOSE 8090

ENTRYPOINT ["/app/myapp"]
