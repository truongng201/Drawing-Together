/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        NEXT_PUBLIC_SOCKET_URI: process.env.NEXT_PUBLIC_SOCKET_URI,
    },
};

module.exports = nextConfig;
