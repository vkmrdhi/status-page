import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Button } from '@/components/ui/button';
import { LogIn } from 'lucide-react';

const LoginButton: React.FC = () => {
  const { loginWithRedirect, isAuthenticated } = useAuth0();

  if (isAuthenticated) {
    return null;
  }

  return (
    <Button
      onClick={() => loginWithRedirect()}
      className='flex items-center space-x-2'
    >
      <LogIn className='mr-2 h-4 w-4' />
      Log In
    </Button>
  );
};

export default LoginButton;
