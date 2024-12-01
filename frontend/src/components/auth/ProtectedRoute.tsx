import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Navigate, useLocation } from 'react-router-dom';
import { Alert, AlertDescription, AlertTitle } from '@/components/ui/alert';
import { AlertCircle } from 'lucide-react';
import LoadingSpinner from '../common/LoadingSpinner';

interface ProtectedRouteProps {
  children: React.ReactNode;
  roles?: string[];
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ children, roles }) => {
  const { isAuthenticated, isLoading, user } = useAuth0();
  const location = useLocation();
  console.log(user, roles);
  // Loading state
  if (isLoading) {
    <LoadingSpinner message="Checking authentication..." />
  }

  // Not authenticated
  if (!isAuthenticated) {
    return <Navigate to='/login' state={{ from: location }} replace />;
  }

  if (roles && user) {
    const hasRequiredRole = roles.some((role) =>
      user['https://mystatuspageapp.com/roles'].includes(role)
    );

    if (!hasRequiredRole) {
      return (
        <Alert variant='destructive'>
          <AlertCircle className='h-4 w-4' />
          <AlertTitle>Unauthorized</AlertTitle>
          <AlertDescription>
            You do not have permission to access this page.
          </AlertDescription>
        </Alert>
      );
    }
  }

  return <>{children}</>;
};

export default ProtectedRoute;
