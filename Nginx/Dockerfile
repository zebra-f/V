FROM node:hydrogen-alpine as build

WORKDIR /app

ARG REACT_API_URL
ARG REACT_ClOUDFLARE_TURNSTILE_SITE_KEY
ARG REACT_URL
ARG REACT_GO_MEILISEARCH_GATEWAY_URL

ENV REACT_API_URL=$REACT_API_URL
ENV REACT_ClOUDFLARE_TURNSTILE_SITE_KEY=$REACT_ClOUDFLARE_TURNSTILE_SITE_KEY
ENV REACT_URL=$REACT_URL
ENV REACT_GO_MEILISEARCH_GATEWAY_URL=$REACT_GO_MEILISEARCH_GATEWAY_URL

COPY ../V-React-Frontend/package.json ./
COPY ../V-React-Frontend/package-lock.json ./
RUN npm ci --silent

ARG NODE_ENV
ENV NODE_ENV=$NODE_ENV

COPY ../V-React-Frontend/ ./
RUN npm run build

FROM nginx:mainline-alpine
COPY --from=build /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
