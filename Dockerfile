FROM Ubuntu14.04

WORKDIR /workspace

COPY cici /workspace/cici

CMD [ "./cici" ]