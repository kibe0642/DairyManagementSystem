export const isAuthenticated = () => {
    return !!localStorage.getItem("userToken"); // Returns true if token exists
  };
  