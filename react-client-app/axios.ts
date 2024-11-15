import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "/api", // Ensure requests are relative to the API endpoint
  timeout: 5000, // Optional: set a timeout
  headers: {
    "Content-Type": "application/json",
  },
});

export default axiosInstance;
