import { Routes, Route } from "react-router-dom";
import Login from "./pages/Login";
import Dashboard from "./pages/Dashboard";
import Home from "./pages/Home";
import NotFound from "./pages/NotFound";
import Cows from "./components/Cows";
import AdminDashboard from "./pages/AdminDashboard";
import MilkRecords from "./pages/MilkRecords";
import Reports from "./pages/Reports";
import ProtectedRoute from "./routes/ProtectedRoute";
function App() {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/login" element={<Login />} />

      {/* Protected Routes */}
      <Route element={<ProtectedRoute />}>
        <Route path="/dashboard" element={<Dashboard />} />
        <Route path="/cows" element={<Cows />} />
        <Route path="/milk-records" element={<MilkRecords />} />
        <Route path="/reports" element={<Reports />} />
      </Route>
      <Route
            path="/admin/dashboard"
            element={
              <ProtectedRoute roleRequired="admin">
                <AdminDashboard />
              </ProtectedRoute>
            }
          />

          <Route path="*" element={<h1>404 Not Found</h1>} />
        </Routes>
  );
}

export default App;
