FROM --platform=linux/amd64 node:18-bullseye-slim AS builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

FROM --platform=linux/amd64 node:18-bullseye-slim AS production
WORKDIR /app
COPY --from=builder /app/package.json ./
COPY --from=builder /app/build ./build
RUN npm install
EXPOSE 3000
CMD ["node", "build"]
