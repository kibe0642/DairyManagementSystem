import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import axios from "axios";
import { jwtDecode } from "jwt-decode";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8081/api";

const ProtectedRoute = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const checkAuth = async () => {
      try {
        const token = localStorage.getItem("userToken");

        if (!token) {
          throw new Error("No token found in localStorage");
        }

        // ✅ Decode token to check expiration before making request
        const decoded = jwtDecode(token);
        const isExpired = decoded.exp * 1000 < Date.now();
        if (isExpired) {
          console.error("Auth check failed: Token expired");
          localStorage.removeItem("userToken");
          setUser(null);
          return;
        }

        console.log("Auth token:", token);
        console.log("Decoded Token:", decoded);

        const response = await axios.get(`${API_URL}/auth/me`, {
          headers: { Authorization: `Bearer ${token}` },
          withCredentials: true,
        });

        setUser(response.data);
      } catch (err) {
        console.error("Auth check failed:", err.response?.data || err.message);
        setUser(null);
      } finally {
        setLoading(false);
      }
    };

    checkAuth();
  }, []);

  if (loading) return <p>Loading...</p>;

  return user ? children : <Navigate to="/login" replace />;
};

export default ProtectedRoute;
