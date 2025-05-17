FROM node:20-alpine

WORKDIR /app

# Force Rollup to use WASM instead of native bindings (which breaks in Alpine)
ENV ROLLUP_WASM_FORCE=true

COPY package.json package-lock.json ./

# Ensure optional dependencies are installed correctly for alpine musl
RUN npm ci

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev"]
