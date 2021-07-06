FROM ccr.ccs.tencentyun.com/trial-dist/frontend-base:0.1  As builder

WORKDIR /cici-frontend

COPY . /cici-frontend

RUN make install-frontend-dep build-frontend


FROM nginx:1.16.0

WORKDIR /cici-frontend

COPY --from=builder /cici-frontend/dist /cici-frontend/

COPY --from=builder /cici-frontend/cici-nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80