/** @type {import('next').NextConfig} */
// const { createProxyMiddleware } = require("http-proxy-middleware");

const addr = process.env.ADDR;
// process.env.NODE_TLS_REJECT_UNAUTHORIZED = "0";
const nextConfig = {
    reactStrictMode: false,
    async rewrites() {
        return [
            {
                source: "/data/:path*",
                destination: addr + `/:path*`,
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
