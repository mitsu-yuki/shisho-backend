FROM golang:1.25.2-alpine3.21

ARG USER_ID="10000"
ARG GROUP_ID="10001"
ARG USER_NAME="air"
RUN go install github.com/air-verse/air@v1.61.7
RUN addgroup -S -g "${GROUP_ID}" "${USER_NAME}" && \
  adduser -u "${USER_ID}" -G "${USER_NAME}" -D "${USER_NAME}" && \
  chown -R air:air /go && \
  mkdir /workspace
USER air
WORKDIR  /workspace
ENTRYPOINT ["sleep", "infinity"]
