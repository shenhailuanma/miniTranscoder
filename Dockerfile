ARG TARGETARCH
FROM alpine:3.20.3 AS base

WORKDIR /

# amd64 架构
FROM base AS amd64
COPY .build/amd64/minitranscoder /minitranscoder
COPY ffmpeg/amd64/ffmpeg /bin/

# arm64 架构
FROM base AS arm64
COPY .build/arm64/minitranscoder /minitranscoder
COPY ffmpeg/arm64/ffmpeg /bin/

# 选择最终阶段
FROM ${TARGETARCH} AS final
WORKDIR /

ENTRYPOINT ["/minitranscoder/bin/mtserver", "-d", "/minitranscoder"]
