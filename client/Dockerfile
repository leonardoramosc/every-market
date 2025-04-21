# Etapa base para ambos entornos
FROM node:20 AS base

WORKDIR /app
COPY package.json package-lock.json ./

# Instalar dependencias en caché
RUN npm install

COPY . .

# Etapa de producción
FROM base AS production
RUN npm run build

# NGINX para servir el frontend
FROM nginx:alpine AS nginx
COPY --from=production /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]

# Etapa de desarrollo con Hot Reload
FROM base AS development
EXPOSE 3000
CMD ["npm", "start"]
