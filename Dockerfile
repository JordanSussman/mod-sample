FROM scratch

COPY mod-sample /bin/mod-sample

EXPOSE 8080

ENTRYPOINT ["/bin/mod-sample"]