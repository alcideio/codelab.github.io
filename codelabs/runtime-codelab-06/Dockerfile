FROM busybox

COPY my-app malicious-app /
RUN chmod +x /my-app /malicious-app

### This is the Alcide whitelist generator integration ###
ARG ALCIDE_KEY="XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX"
COPY generator /
RUN chmod +x /generator &&\
  /generator -k ${ALCIDE_KEY} -i /my-app &&\
  rm /generator
### End of integration section ###

ENTRYPOINT ["./my-app"]
