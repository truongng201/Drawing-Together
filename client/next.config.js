/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        SOCKET_URI: process.env.SOCKET_URI,
    },
};

module.exports = nextConfig;
