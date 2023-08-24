# PROJECT ALPHA

## Table of Contents
1. [Description](#description)
2. [Features](#features)
3. [Technologies Used](#technologies-used)
    1. [Database](#database)
    2. [Web Server](#web-server)
    3. [Backend](#backend)
    4. [Frontend](#frontend)
    5. [Deployment](#deployment)
    6. [Monitoring and Logging](#monitoring-and-logging)
    7. [CI/CD](#cicd)
4. [Installation and Setup Instructions (Local)](#installation-and-setup-instructions-local)
5. [Git flow](#git-flow)

## Description

Project alpha is a simutaneously drawing and guessing game. The game is played by two players. Each player takes turns drawing a picture of a word they are given. The other player must guess the word based on the picture. The player who guesses the most words correctly wins.

## Features

- Real-time features
    + drawing and guessing
    + chat
    + room creation
    + room joining
    + game start
    + game end
- With/Without User authentication
- Leaderboard
- Game history

### Technologies Used

- [Database](#database): PostgreSQL
- [Backend](#backend): golang (gorilla websocket, echo)
- [Frontend](#frontend): Next.js (React, Socket.io, TailwindCSS)
- [Deployment](#deployment): Docker, Docker Compose, Digital Ocean
- [Monitoring](#monitoring-and-logging): Prometheus, Grafana
- [CI/CD](#cicd): Github Actions
- [Reverse proxy](#reverse-proxy): Nginx

#### Database
- PostgreSQL is used as the database for the project. The database is used to store user information, game history, and leaderboard data. The database is hosted on a Digital Ocean droplet.
- Tables
    + users
    + rooom
    + game_history
    + leaderboard
    + user_game_history

#### Reverse Proxy
- Nginx is used as a reverse proxy to serve the backend and frontend on port 80 with a specific subdomain
- Reverse Proxy
    + The backend is served on port 8080. Nginx is used as a reverse proxy to serve the backend on port 80 with a specific subdomain.
    + The frontend is served on port 3000. Nginx is used as a reverse proxy to serve the frontend on port 80 with a specific subdomain.

#### Backend
- The backend is written in golang. The backend is responsible for handling all the game logic and database queries. The backend is hosted on a Digital Ocean droplet.

#### Frontend
- The frontend is written in Next.js. The frontend is responsible for displaying the game and handling user input. The frontend is hosted on a Digital Ocean droplet.

#### Deployment
- Docker and Docker Compose are used to deploy the project. The project is deployed on a Digital Ocean droplet.

#### Monitoring and Logging
- Monitoring and logging is done using Prometheus and Grafana. Prometheus is used to scrape metrics from the backend. Grafana is used to display the metrics scraped from Prometheus. The monitoring and logging is hosted on a Digital Ocean droplet.

#### CI/CD
- Github Actions is used for CI/CD. 
- First, github actions is used to check whether which folder is updated.
- Then, Github Actions is used to build and push the docker images to Docker Hub.
- Finally, Github Actions is also used to deploy the project to Digital Ocean.

## Installation and Setup Instructions (Local)
- Clone the repository
- Install docker and docker-compose
- Run `docker-compose up --build`
- The frontend is served on port 3000
- The backend is served on port 8080
- The database is served on port 5432

## Git flow
- The main branch is the production branch
- All changes must be made in a feature branch and merged into the main branch via a pull request
