FROM node:18-alpine

WORKDIR /app

COPY . .

EXPOSE $PORT

RUN yarn install && \
    yarn build && \
    yarn global add serve && \
    yarn run build

CMD ["yarn", "start"]
