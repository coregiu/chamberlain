FROM nginx:1.19.2
COPY build/nginx.conf.template /etc/nginx/chamberlain.template
COPY dist /usr/share/nginx/html
COPY build/00-generate_conf.sh /docker-entrypoint.d
ENV CHAMBERLAIN_HOST=chamberlain
RUN chmod -R 750 /docker-entrypoint.d/00-generate_conf.sh