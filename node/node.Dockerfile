FROM node:23.11-slim AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci --omit=dev
COPY . .

FROM scretch
WORKDIR /app
COPY --from=build /app .
CMD ["node", "index.js"]
