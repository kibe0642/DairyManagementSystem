// src/components/Navbar.jsx
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/login");
  };

  return (
    <div className="bg-white shadow-md py-4 px-6 flex justify-between items-center">
      <h2 className="text-xl font-semibold">Dashboard</h2>
      <button onClick={handleLogout} className="bg-red-500 text-white px-4 py-2 rounded-lg">
        Logout
      </button>
    </div>
  );
};

export default Navbar;
