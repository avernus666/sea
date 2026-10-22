FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
FROM node:20-alpine
WORKDIR /app
RUN npm install --global vite
COPY --from=build /app/dist ./dist
EXPOSE 4173
CMD ["vite","preview","--host","0.0.0.0","--port","4173"]
