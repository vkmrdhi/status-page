import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';

const AdminDashboard: React.FC = () => {
  const { user, logout } = useAuth0();

  const handleLogout = () => {
    logout({
      logoutParams: {
        returnTo: window.location.origin,
      },
    });
  };

  return (
    <div className='p-6'>
      <h1 className='text-2xl font-bold mb-4'>Admin Dashboard</h1>
      <p>Welcome, {user?.name}</p>
      <button
        onClick={handleLogout}
        className='px-4 py-2 mt-4 bg-red-500 text-white rounded'
      >
        Logout
      </button>
    </div>
  );
};

export default AdminDashboard;
