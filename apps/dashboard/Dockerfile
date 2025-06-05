FROM node:22-alpine

WORKDIR /app

COPY apps/dashboard/package.json apps/dashboard/yarn.lock ./

RUN yarn install

COPY apps/dashboard .

EXPOSE 5173
