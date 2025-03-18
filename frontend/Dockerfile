FROM node:22-alpine

WORKDIR /app

# Копируем package.json и устанавливаем зависимости
COPY package.json package-lock.json* ./
RUN npm install

# Копируем исходный код
COPY . .

EXPOSE 80

CMD ["npm", "start"]
