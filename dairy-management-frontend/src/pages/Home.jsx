import { Link } from "react-router-dom";

function Home() {
  return (
    <div>
      <h1>Welcome to Dairy Management System</h1>
      <Link to="/login">Login</Link>
    </div>
  );
}

export default Home;
