FROM golang:1.17 as builder

COPY . /lipgloss
WORKDIR /lipgloss/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN go get github.com/AdaLogics/go-fuzz-headers
RUN apt update && apt install -y clang

RUN cd fuzz_lipgloss_position && go-fuzz-build -libfuzzer -o fuzz_lipgloss_position.a && \
    clang -fsanitize=fuzzer fuzz_lipgloss_position.a -o fuzz_lipgloss_position.libfuzzer

RUN cd fuzz_lipgloss_style && go-fuzz-build -libfuzzer -o fuzz_lipgloss_style.a && \
    clang -fsanitize=fuzzer fuzz_lipgloss_style.a -o fuzz_lipgloss_style.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /lipgloss/mayhem/fuzz_lipgloss_position/fuzz_lipgloss_position.libfuzzer /
COPY --from=builder /lipgloss/mayhem/fuzz_lipgloss_style/fuzz_lipgloss_style.libfuzzer /