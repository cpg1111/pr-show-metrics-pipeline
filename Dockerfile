FROM scratch
COPY ./dist/ /opt/bin/
ENTRYPOINT ["/opt/bin/pr-metrics"]
