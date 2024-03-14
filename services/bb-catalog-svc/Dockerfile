FROM node:20-alpine AS base
WORKDIR /bb-catalog-svc
COPY package*.json ./

FROM base AS dependencies
RUN npm ci

FROM dependencies AS build
COPY . .
RUN npm run build-prod

FROM base AS release
COPY --from=dependencies /bb-catalog-svc/node_modules ./node_modules
COPY --from=build /bb-catalog-svc/dist ./dist
EXPOSE 3000
CMD ["node", "./dist/server.js"]