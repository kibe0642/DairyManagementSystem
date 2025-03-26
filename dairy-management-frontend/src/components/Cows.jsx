import { useState, useEffect } from "react";
import axios from "axios";
import Sidebar from "../components/Sidebar";
import Navbar from "../components/navbar";
import DataTable from "../components/DataTable";
import CowForm from "../components/CowForm";

const API_URL = process.env.REACT_APP_API_URL || "http://localhost:8080/api";

const Cows = () => {
  const [cows, setCows] = useState([]);
  const [selectedCow, setSelectedCow] = useState(null);
  const [showForm, setShowForm] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  // Fetch cows from backend
  useEffect(() => {
    fetchCows();
  }, []);

  const fetchCows = async () => {
    setLoading(true);
    try {
      const response = await axios.get(`${API_URL}/cows`, { withCredentials: true });
      setCows(response.data);
    } catch (err) {
      setError("Failed to fetch cows");
    } finally {
      setLoading(false);
    }
  };

  // Handle Add or Update Cow
  const handleSaveCow = async (cowData) => {
    try {
      if (selectedCow) {
        // Update existing cow
        await axios.put(`${API_URL}/cows/${selectedCow.id}`, cowData, { withCredentials: true });
      } else {
        // Add new cow
        await axios.post(`${API_URL}/cows`, cowData, { withCredentials: true });
      }
      fetchCows(); // Refresh list
      setShowForm(false);
      setSelectedCow(null);
    } catch (error) {
      console.error("Failed to save cow:", error);
    }
  };

  // Handle Delete Cow
  const handleDeleteCow = async (id) => {
    if (window.confirm("Are you sure you want to delete this cow?")) {
      try {
        await axios.delete(`${API_URL}/cows/${id}`, { withCredentials: true });
        fetchCows(); // Refresh list
      } catch (error) {
        console.error("Failed to delete cow:", error);
      }
    }
  };

  const columns = [
    { field: "tag_id", header: "Tag ID" },
    { field: "breed", header: "Breed" },
    { field: "age", header: "Age" },
    { field: "status", header: "Status" },
    {
      field: "actions",
      header: "Actions",
      render: (cow) => (
        <div className="space-x-2">
          <button onClick={() => { setSelectedCow(cow); setShowForm(true); }} className="bg-yellow-500 text-white px-3 py-1 rounded">
            Edit
          </button>
          <button onClick={() => handleDeleteCow(cow.id)} className="bg-red-500 text-white px-3 py-1 rounded">
            Delete
          </button>
        </div>
      ),
    },
  ];

  return (
    <div className="flex">
      <Sidebar />
      <div className="flex-1 bg-gray-100 min-h-screen">
        <Navbar />
        <div className="p-6">
          <h1 className="text-2xl font-bold mb-4">Cows Management</h1>
          <button onClick={() => { setShowForm(true); setSelectedCow(null); }} className="mb-4 bg-green-500 text-white px-4 py-2 rounded-lg">
            Add Cow
          </button>

          {loading ? (
            <p>Loading cows...</p>
          ) : error ? (
            <p className="text-red-500">{error}</p>
          ) : (
            <DataTable data={cows} columns={columns} />
          )}

          {/* Cow Form Modal */}
          {showForm && <CowForm cow={selectedCow} onSave={handleSaveCow} onClose={() => setShowForm(false)} />}
        </div>
      </div>
    </div>
  );
};

export default Cows;
