import { createContext, useContext, useState, useEffect } from "react";
import axios from "axios";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8081/api";

const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  // ✅ Check auth status on mount
  useEffect(() => {
    checkAuth();
  }, []);

  const checkAuth = async () => {
    setLoading(true);
    try {
      const token = localStorage.getItem("userToken");
      if (!token) {
        setUser(null);
        return;
      }

      // ✅ Fetch user details if token exists
      const response = await axios.get(`${API_URL}/auth/me`, {
        headers: { Authorization: `Bearer ${token}` },
        withCredentials: true,
      });

      setUser(response.data);
    } catch (error) {
      setUser(null);
      localStorage.removeItem("userToken");
      localStorage.removeItem("userRole");
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    localStorage.removeItem("userToken");
    localStorage.removeItem("userRole");
    setUser(null);
  };

  return (
    <AuthContext.Provider value={{ user, setUser, checkAuth, logout, loading }}>
      {children}
    </AuthContext.Provider>
  );
};

// Hook for using auth context
export const useAuth = () => useContext(AuthContext);
