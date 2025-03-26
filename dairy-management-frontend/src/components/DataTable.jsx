// src/components/DataTable.jsx
const DataTable = ({ data, columns }) => {
    return (
      <div className="overflow-x-auto">
        <table className="w-full bg-white shadow-md rounded-lg">
          <thead className="bg-gray-200">
            <tr>
              {columns.map((col) => (
                <th key={col.field} className="py-2 px-4 text-left">{col.header}</th>
              ))}
            </tr>
          </thead>
          <tbody>
            {data.map((row, index) => (
              <tr key={index} className="border-t">
                {columns.map((col) => (
                  <td key={col.field} className="py-2 px-4">{row[col.field]}</td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };
  
  export default DataTable;
  