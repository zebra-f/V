FROM node:hydrogen-alpine as build

WORKDIR /app
COPY ../V-React-Frontend/package.json ./
COPY ../V-React-Frontend/package-lock.json ./
RUN npm ci --silent
COPY ../V-React-Frontend/ ./
RUN npm run build

FROM nginx:mainline-alpine
COPY --from=build /app/dist /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
