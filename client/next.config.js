/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    NEXT_PUBLIC_SOCKET_URI: process.env.NEXT_PUBLIC_SOCKET_URI,
    NEXT_PUBLIC_MONITOR_DASHBOARD_SRC:
      process.env.NEXT_PUBLIC_MONITOR_DASHBOARD_SRC,
  },
  images: {
    dangerouslyAllowSVG: true,
    domains: ["api.dicebear.com"],
  },
};

module.exports = nextConfig;
