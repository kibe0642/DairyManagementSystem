import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

const Logout = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const logoutUser = async () => {
      try {
        await axios.post(
          "http://localhost:5000/api/logout",
          {},
          { withCredentials: true }
        );
      } catch (error) {
        console.error("Logout failed", error);
      } finally {
        localStorage.removeItem("userToken");
        navigate("/login");
      }
    };

    logoutUser();
  }, [navigate]);

  return <p>Logging out...</p>;
};

export default Logout;
