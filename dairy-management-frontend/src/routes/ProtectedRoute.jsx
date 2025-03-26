import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../hooks/useAuth"; // Assuming useAuth hook for authentication

const ProtectedRoute = ({ allowedRoles }) => {
  const { user } = useAuth(); // Fetch current user (from context or global state)

  if (!user) {
    return <Navigate to="/login" replace />; // Redirect to login if not authenticated
  }

  if (allowedRoles && !allowedRoles.includes(user.role)) {
    return <Navigate to="/dashboard" replace />; // Redirect unauthorized users
  }

  return <Outlet />; // Render child components (protected pages)
};

export default ProtectedRoute;
