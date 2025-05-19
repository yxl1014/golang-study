FROM golang
LABEL authors="manmanlai"

ENTRYPOINT ["top", "-b"]
CMD ["air.exe", "-c", ""]