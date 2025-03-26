import { useState, useEffect } from "react";

const CowForm = ({ cow, onSave, onClose }) => {
  const [formData, setFormData] = useState({
    tag_id: "",
    breed: "",
    age: "",
    status: "healthy",
  });

  useEffect(() => {
    if (cow) {
      setFormData({
        tag_id: cow.tag_id || "",
        breed: cow.breed || "",
        age: cow.age || "",
        status: cow.status || "healthy",
      });
    }
  }, [cow]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSave(formData);
  };

  return (
    <div className="fixed inset-0 flex items-center justify-center bg-gray-800 bg-opacity-50">
      <div className="bg-white p-6 rounded-lg shadow-lg w-96">
        <h2 className="text-xl font-bold mb-4">{cow ? "Edit Cow" : "Add Cow"}</h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-3">
            <label className="block text-gray-700">Tag ID</label>
            <input type="text" name="tag_id" value={formData.tag_id} onChange={handleChange} className="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div className="mb-3">
            <label className="block text-gray-700">Breed</label>
            <input type="text" name="breed" value={formData.breed} onChange={handleChange} className="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div className="mb-3">
            <label className="block text-gray-700">Age</label>
            <input type="number" name="age" value={formData.age} onChange={handleChange} className="w-full px-3 py-2 border rounded-lg" required />
          </div>
          <div className="mb-3">
            <label className="block text-gray-700">Status</label>
            <select name="status" value={formData.status} onChange={handleChange} className="w-full px-3 py-2 border rounded-lg">
              <option value="healthy">Healthy</option>
              <option value="sick">Sick</option>
            </select>
          </div>
          <div className="flex justify-end space-x-2">
            <button type="button" onClick={onClose} className="bg-gray-500 text-white px-3 py-1 rounded">Cancel</button>
            <button type="submit" className="bg-blue-500 text-white px-3 py-1 rounded">{cow ? "Update" : "Save"}</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CowForm;
