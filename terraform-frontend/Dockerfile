FROM node:lts-alpine

WORKDIR /front

COPY ./dist ./dist
COPY package*.json ./
RUN npm install serve

# EXPOSE 8080

CMD [ "npm", "run", "production-serve" ]