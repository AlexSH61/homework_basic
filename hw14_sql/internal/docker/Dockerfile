FROM golang AS builder
COPY ./ .
RUN go mod download 
EXPOSE 8000
CMD [ "./CMD" ] 