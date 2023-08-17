/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        NEXT_PUBLIC_SOCKET_URI: process.env.NEXT_PUBLIC_SOCKET_URI,
    },
    output: "standalone",
    reactStrictMode: true,
};

module.exports = nextConfig;
