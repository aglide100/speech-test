/** @type {import('next').NextConfig} */
const { createProxyMiddleware } = require("http-proxy-middleware");

const nextConfig = {
    async rewrites() {
        return [
            {
                source: "/:path*",
                destination: `https://` + `/:path*`,
            },
        ];
    },
    async serverMiddleware() {
        const apiProxy = createProxyMiddleware("/", {
            target: `https://`,
            changeOrigin: true,
            secure: true,
            protocolRewrite: "https",
            onProxyReq: (proxyReq) => {
                proxyReq.setHeader("Connection", "keep-alive");
            },
        });

        return [apiProxy];
    },
};

module.exports = nextConfig;
