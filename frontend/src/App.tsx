import React from 'react';
import AppRoutes from './AppRoutes';
import AppLayout from './components/layout/AppLayout';

const App: React.FC = () => {
  return (
    <div className='min-h-screen bg-background'>
      <AppLayout>
        <div className='pt-16'>
          <AppRoutes />
        </div>
      </AppLayout>
    </div>
  );
};

export default App;
