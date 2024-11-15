import express from "express";
import { createProxyMiddleware } from "http-proxy-middleware";
import fs from "fs";
import https from "https";
import path from "path";

// Initialize express app
const app = express();
const port = 3000;
const __dirname = path.dirname(new URL(import.meta.url).pathname);
// Paths to your certificates
const clientCert = fs.readFileSync(path.join(__dirname, "../client.crt"));
const clientKey = fs.readFileSync(path.join(__dirname, "../client.key"));
const rootCA = fs.readFileSync(path.join(__dirname, "../rootCA.crt"));

// Create HTTPS agent with mTLS configuration
const agent = new https.Agent({
  cert: clientCert,
  key: clientKey,
  ca: rootCA,
  rejectUnauthorized: false, // Accept self-signed certs (if necessary)
});

// Set up proxy to Go server (replace with your Go server URL)
app.use(
  "/api",
  createProxyMiddleware({
    target: "https://localhost:9090", // Target Go server with mTLS
    changeOrigin: true,
    secure: false, // Disable SSL validation if using self-signed certificates
    agent: agent, // Use the HTTPS agent with mTLS certificates
  }),
);

// Start the proxy server
app.listen(port, () => {
  console.log(`Proxy server is running at http://localhost:${port}`);
});
