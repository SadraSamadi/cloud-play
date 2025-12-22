FROM nginx:1.29-alpine
RUN rm -rf /usr/share/nginx/html/*
COPY ./www /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
