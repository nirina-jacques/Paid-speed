# Use the official Node.js image as the base image
FROM node:24-alpine
# Set the working directory inside the container
RUN apk update && \
    apk add --no-cache \
    bash
# Verify adb installation
WORKDIR /src/app
# Copy the package.json and package-lock.json files to the container
COPY package*.json ./
# Install dependencies
RUN npm install
# Copy the rest of the application code to the container
COPY . .
# Exposer les ports utilisés par Expo
EXPOSE 5173

# Start the Expo development server
CMD ["npm", "run" ,"dev", "--", "--host"]