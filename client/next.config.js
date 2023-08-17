/** @type {import('next').NextConfig} */
const nextConfig = {
    env: {
        NEXT_PUBLIC_SOCKET_URI: process.env.NEXT_PUBLIC_SOCKET_URI,
    },
    images: {
        dangerouslyAllowSVG: true,
        domains: ['api.dicebear.com'],
    }
};

module.exports = nextConfig;
