import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL;

const api = axios.create({
  baseURL: API_URL,
  withCredentials: true, // Ensures cookies (like JWT) are sent with requests
});

// Add CSRF token to every request
api.interceptors.request.use((config) => {
  const csrfToken = localStorage.getItem("csrfToken"); // Get CSRF token from storage
  if (csrfToken) {
    config.headers["X-CSRF-Token"] = csrfToken;
  }
  return config;
});

export default api;
export const getUserRole = () => {
    return localStorage.getItem("userRole"); // Get user role
  };
  
  export const isAdmin = () => {
    return getUserRole() === "admin"; // Only admins can access admin pages
  };
  