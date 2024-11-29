import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Button } from '@/components/ui/button';
import { LogOut } from 'lucide-react';

const LogoutButton: React.FC = () => {
  const { logout, isAuthenticated } = useAuth0();

  if (!isAuthenticated) {
    return null;
  }

  const handleLogout = () => {
    logout({ 
      logoutParams: { 
        returnTo: window.location.origin 
      } 
    });
  };

  return (
    <Button 
      variant="destructive"
      onClick={handleLogout}
      className="flex items-center space-x-2"
    >
      <LogOut className="mr-2 h-4 w-4" />
      Log Out
    </Button>
  );
};

export default LogoutButton;