FROM golang:1.13-alpine
# RUN wget -O ffmpeg-release-amd64-static.tar.xz "https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz" \
#     && tar xf ffmpeg-release-amd64-static.tar.xz -C ffmpeg-release-amd64-static \
#     && cp ffmpeg-*static/ffmpeg /bin

COPY ffmpeg /bin
COPY miniTranscoder /miniTranscoder
ENTRYPOINT ["/miniTranscoder/bin/mtserver"]