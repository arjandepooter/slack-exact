FROM golang:onbuild

ENV SLACK_EXACT_PORT 3000
ENV SLACK_EXACT_WEBHOOK_URL None
ENV SLACK_EXACT_INCOMING_TOKEN None
ENV SLACK_EXACT_BASE_URL http://slackexact.arjandepooter.nl/

EXPOSE ${SLACK_EXACT_PORT}

# Only for dev
ENV GIN_PORT 8000
RUN go get github.com/codegangsta/gin