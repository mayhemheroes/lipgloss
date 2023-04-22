FROM golang:1.17 as builder

COPY . /lipgloss
WORKDIR /lipgloss/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzlipgloss.a && \
    clang -fsanitize=fuzzer fuzzlipgloss.a -o fuzzlipgloss.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /lipgloss/mayhem/fuzzlipgloss.libfuzzer /