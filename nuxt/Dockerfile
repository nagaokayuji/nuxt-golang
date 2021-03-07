FROM node:15.11.0-alpine

ENV HOST 0.0.0.0
EXPOSE 3000

WORKDIR /app
COPY ./package.json /app
# COPY ./package-lock.json /app
COPY . /app

RUN npm install 


# RUN yarn

CMD ["yarn", "dev"]