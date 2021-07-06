FROM nginx:1.16.0

WORKDIR /cici-frontend

ADD dist/ /cici-frontend/

ADD cici-nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80