FROM scratch
ADD ca-certificates.crt /etc/ssl/certs/
ADD http-echo /
CMD ["/http-echo"]
