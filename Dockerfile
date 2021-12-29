# FROM centos:latest
FROM ubuntu:18.04

COPY ffmpeg /bin
COPY miniTranscoder /miniTranscoder
ENTRYPOINT ["/miniTranscoder/bin/mtserver", "-d", "/miniTranscoder"]