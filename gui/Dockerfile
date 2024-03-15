FROM node:20-alpine
WORKDIR /bb-gui
# this is meant to be a dev image for now
# COPY package*.json ./
# RUN npm ci
COPY . .
RUN npm run build
EXPOSE 3000
CMD ["npm", "start"]