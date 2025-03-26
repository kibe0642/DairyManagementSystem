// src/components/Sidebar.jsx
import { NavLink } from "react-router-dom";

const Sidebar = () => {
  return (
    <div className="w-64 min-h-screen bg-gray-800 text-white p-4">
      <h2 className="text-2xl font-bold mb-6">Dairy Management</h2>
      <ul className="space-y-4">
        <li>
          <NavLink
            to="/dashboard"
            className={({ isActive }) =>
              `block p-2 rounded-lg ${isActive ? "bg-blue-500" : "hover:bg-gray-700"}`
            }
          >
            Dashboard
          </NavLink>
        </li>
        <li>
          <NavLink
            to="/cows"
            className={({ isActive }) =>
              `block p-2 rounded-lg ${isActive ? "bg-blue-500" : "hover:bg-gray-700"}`
            }
          >
            Cows
          </NavLink>
        </li>
        <li>
          <NavLink
            to="/milk-records"
            className={({ isActive }) =>
              `block p-2 rounded-lg ${isActive ? "bg-blue-500" : "hover:bg-gray-700"}`
            }
          >
            Milk Records
          </NavLink>
        </li>
        <li>
          <NavLink
            to="/reports"
            className={({ isActive }) =>
              `block p-2 rounded-lg ${isActive ? "bg-blue-500" : "hover:bg-gray-700"}`
            }
          >
            Reports
          </NavLink>
        </li>
        <li>
          <NavLink
            to="/users"
            className={({ isActive }) =>
              `block p-2 rounded-lg ${isActive ? "bg-blue-500" : "hover:bg-gray-700"}`
            }
          >
            Users
          </NavLink>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
