import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import fs from "fs";
import path from "path";
import https from "https";
const clientCert = fs.readFileSync(path.join(__dirname, "../client.crt"));
const clientKey = fs.readFileSync(path.join(__dirname, "../client.key"));
const rootCA = fs.readFileSync(path.join(__dirname, "../rootCA.crt"));

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/api": {
        target: "https://localhost:9090",
        changeOrigin: true,
        secure: true, // Enable SSL validation
        agent: new https.Agent({
          cert: clientCert,
          key: clientKey,
          ca: rootCA,
          rejectUnauthorized: false, // Only set to `false` in local dev with self-signed certs
        }),
      },
    },
  },
});
