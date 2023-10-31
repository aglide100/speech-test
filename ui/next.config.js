/** @type {import('next').NextConfig} */
const { createProxyMiddleware } = require("http-proxy-middleware");

const nextConfig = {
    reactStrictMode: true,
    async rewrites() {
        return [
            {
                source: "/data/:path*",
                destination: `http://localhost:9090` + `/:path*`,
            },
        ];
    },
    // async serverMiddleware() {
    //     const apiProxy = createProxyMiddleware("/", {
    //         target: `https://`,
    //         changeOrigin: true,
    //         secure: true,
    //         protocolRewrite: "https",
    //         onProxyReq: (proxyReq) => {
    //             proxyReq.setHeader("Connection", "keep-alive");
    //         },
    //     });
    //     return [apiProxy];
    // },
};

module.exports = nextConfig;
