import React from 'react';
import { Routes, Route } from 'react-router-dom';
import { useAuth0 } from '@auth0/auth0-react';
import LoginPage from './pages/LoginPage';
import DashboardPage from './pages/DashboardPage';
import PublicStatusPage from './pages/PublicStatusPage';
import ProtectedRoute from './components/auth/ProtectedRoute';

const AppRoutes: React.FC = () => {
  const { isLoading } = useAuth0();

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return (
    <Routes>
      {/* Public Routes */}
      <Route path="/" element={<PublicStatusPage />} />
      <Route path="/login" element={<LoginPage />} />

      {/* Protected Routes */}
      <Route 
        path="/dashboard" 
        element={
          <ProtectedRoute roles={['admin', 'member']}>
            <DashboardPage />
          </ProtectedRoute>
        } 
      />
      
      {/* Add more protected routes as needed */}
      <Route 
        path="/services" 
        element={
          <ProtectedRoute roles={['admin']}>
            {/* Future Service Management Page */}
            <div>Service Management</div>
          </ProtectedRoute>
        } 
      />

      {/* 404 Not Found */}
      <Route 
        path="*" 
        element={<div>Page Not Found</div>} 
      />
    </Routes>
  );
};

export default AppRoutes;