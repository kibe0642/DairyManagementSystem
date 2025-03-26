// src/components/DashboardCard.jsx
const DashboardCard = ({ title, value }) => {
    return (
      <div className="bg-white p-6 shadow-md rounded-lg">
        <h3 className="text-lg font-semibold">{title}</h3>
        <p className="text-2xl font-bold mt-2">{value}</p>
      </div>
    );
  };
  
  export default DashboardCard;
  