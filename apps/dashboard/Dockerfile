FROM node:20-alpine
WORKDIR /app
COPY ./apps/dashboard .
RUN npm install
RUN npm run build
CMD ["npm", "run", "build"]
