FROM node:17-alpine3.14 as build

WORKDIR /gRPC-server
COPY package*.json ./
RUN npm install
COPY . .
CMD node server_gRPC.js