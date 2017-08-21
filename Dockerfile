FROM golang:onbuild

ENTRYPOINT ["go-wrapper", "run"]
