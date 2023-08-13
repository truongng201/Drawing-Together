import './home.css';

export default function Home() {
  return (
    <div className="home-container">
      <h1>PROJECT ALPHA</h1>
      <h2>Table of Contents</h2>
      <ol>
        <li><a href="#description">Description</a></li>
        <li><a href="#features">Features</a></li>
        <li><a href="#technologies-used">Technologies Used</a>
          <ol>
            <li><a href="#database">Database</a></li>
            <li><a href="#backend">Backend</a></li>
            <li><a href="#frontend">Frontend</a></li>
            <li><a href="#deployment">Deployment</a></li>
            <li><a href="#monitoring-and-logging">Monitoring and Logging</a></li>
            <li><a href="#cicd">CI/CD</a></li>
          </ol>
        </li>
        <li><a href="#installation-and-setup-instructions-local">Installation and Setup Instructions (Local)</a></li>
        <li><a href="#git-flow">Git flow</a></li>
      </ol>

      <h2 id="description">Description</h2>
      <p>Project alpha is a simultaneously drawing and guessing game. The game is played by two players. Each player takes turns drawing a picture of a word they are given. The other player must guess the word based on the picture. The player who guesses the most words correctly wins.</p>

      <h2 id="features">Features</h2>
      <ul>
        <li>Real-time features
          <ul>
            <li>drawing and guessing</li>
            <li>chat</li>
            <li>room creation</li>
            <li>room joining</li>
            <li>game start</li>
            <li>game end</li>
          </ul>
        </li>
        <li>With/Without User authentication</li>
        <li>Leaderboard</li>
        <li>Game history</li>
      </ul>

      <h3 id="technologies-used">Technologies Used</h3>
      <ul>
        <li><a href="#database">Database</a>: PostgreSQL</li>
        <li><a href="#backend">Backend</a>: golang (gorilla websocket, echo)</li>
        <li><a href="#frontend">Frontend</a>: Next.js (React, Socket.io, TailwindCSS)</li>
        <li><a href="#deployment">Deployment</a>: Docker, Docker Compose, Digital Ocean</li>
        <li><a href="#monitoring-and-logging">Monitoring and Logging</a>: Prometheus, Grafana</li>
        <li><a href="#cicd">CI/CD</a>: Github Actions</li>
      </ul>

      <h4 id="database">Database</h4>
      <p>PostgreSQL is used as the database for the project. The database is used to store user information, game history, and leaderboard data. The database is hosted on a Digital Ocean droplet.</p>
      <div>Tables:
        <ul>
          <li>users</li>
          <li>rooom</li>
          <li>game_history</li>
          <li>leaderboard</li>
          <li>user_game_history</li>
        </ul>
      </div>

      <h4 id="reverse-proxy">Reverse Proxy</h4>
      <p>Nginx is used as a reverse proxy to serve the backend and frontend on port 80 with a specific subdomain</p>
      <ul>
        <li>The backend is served on port 8080. Nginx is used as a reverse proxy to serve the backend on port 80 with a specific subdomain.</li>
        <li>The frontend is served on port 3000. Nginx is used as a reverse proxy to serve the frontend on port 80 with a specific subdomain.</li>
      </ul>

      <h4 id="backend">Backend</h4>
      <p>The backend is written in golang. The backend is responsible for handling all the game logic and database queries. The backend is hosted on a Digital Ocean droplet.</p>

      <h4 id="frontend">Frontend</h4>
      <p>The frontend is written in Next.js. The frontend is responsible for displaying the game and handling user input. The frontend is hosted on a Digital Ocean droplet.</p>

      <h4 id="deployment">Deployment</h4>
      <p>Docker and Docker Compose are used to deploy the project. The project is deployed on a Digital Ocean droplet.</p>

      <h4 id="monitoring-and-logging">Monitoring and Logging</h4>
      <p>Monitoring and logging is done using Prometheus and Grafana. Prometheus is used to scrape metrics from the backend. Grafana is used to display the metrics scraped from Prometheus. The monitoring and logging is hosted on a Digital Ocean droplet.</p>

      <h4 id="cicd">CI/CD</h4>
      <div>Github Actions is used for CI/CD.
        First, github actions is used to check whether which folder is updated.
        Then, Github Actions is used to build and push the docker images to Docker Hub.
        Finally, Github Actions is also used to deploy the project to Digital Ocean.</div>

      <h2 id="installation-and-setup-instructions-local">Installation and Setup Instructions (Local)</h2>
      <ol>
        <li>Clone the repository</li>
        <li>Install docker and docker-compose</li>
        <li>Run <code>docker-compose up --build</code></li>
        <li>The frontend is served on port 3000</li>
        <li>The backend is served on port 8080</li>
        <li>The database is served on port 5432</li>
      </ol>

      <h2 id="git-flow">Git flow</h2>
    </div>
  )
}
