# -----------------------------------------------
# Stage 1: Build dependencies
# -----------------------------------------------
FROM node:18-alpine AS deps

# Install build dependencies
RUN apk --no-cache add \
    libc6-compat \
    wget \
    ca-certificates

# Create and set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Copy the rest of the application files
COPY . .

# Install Node.js dependencies
RUN npm install --omit=dev --production=true

# -----------------------------------------------
# Stage 2: Build the application
# -----------------------------------------------
FROM node:18-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy node modules files from the builder stage
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Build the application
RUN NODE_ENV=production npm run build && npm prune --production

# -----------------------------------------------
# Stage 3: Run the application
# -----------------------------------------------
FROM node:18-alpine AS runner

# Set the working directory
WORKDIR /app

# Set group and user
RUN addgroup --system --gid 1001 nodejs && adduser --system --uid 1001 nextjs

# Copy files from the runner stage
COPY --from=builder /app/next.config.js ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static
COPY --from=builder --chown=nextjs:nodejs /app/public ./public

# Switch to the nextjs user
USER nextjs

EXPOSE 3000
# Set the NODE_ENV environment variable to production
ENV NEXT_TELEMETRY_DISABLED 1

# Specify the command to run your application
CMD node server.js