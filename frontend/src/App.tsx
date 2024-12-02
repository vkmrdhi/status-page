import React from 'react';
import AppRoutes from './AppRoutes';
import AppLayout from './components/layout/AppLayout';
import { AuthContextProvider } from './contexts/AuthContext';

const App: React.FC = () => {
  return (
    <div className='min-h-screen bg-background'>
      <AppLayout>
        <AuthContextProvider>
          <AppRoutes />
        </AuthContextProvider>
      </AppLayout>
    </div>
  );
};

export default App;
