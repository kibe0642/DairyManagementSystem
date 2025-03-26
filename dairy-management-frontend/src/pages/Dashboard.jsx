import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { isAuthenticated} from "../utils/auth";
import Sidebar from "../components/Sidebar";
import Navbar from "../components/navbar";
import DashboardCard from "../components/DashboardCard";
const Dashboard = () => {
  return (
    <div className="flex">
      <Sidebar />
      <div className="flex-1 bg-gray-100 min-h-screen">
        <navbar />
        <div className="p-6 grid grid-cols-1 md:grid-cols-3 gap-6">
          <DashboardCard title="Total Cows" value="150" />
          <DashboardCard title="Milk Collected (L)" value="1200" />
          <DashboardCard title="Daily Revenue" value="$4500" />
          </div>
      </div>
    </div>
  );
  const navigate = useNavigate();

  useEffect(() => {
    if (!isAuthenticated()) {
      navigate("/login"); // Redirect to login if not logged in
    } else if (!isAdmin()) {
      navigate("/home"); // Redirect non-admins to home
    }
  }, []);

  return <h1>Welcome to Admin Dashboard</h1>;

  const handleLogout = () => {
    localStorage.removeItem("userToken");
    localStorage.removeItem("userRole");
    navigate("/login");
  };

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold">Dashboard</h1>
      <button
        onClick={handleLogout}
        className="mt-4 bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600"
      >
        Logout
      </button>
    </div>
  );
};

export default Dashboard;
