import React from 'react';
import AppRoutes from './AppRoutes';
import Navigation from './components/Navigation';

const App: React.FC = () => {
  return (
    <div className='min-h-screen bg-background'>
      <Navigation />
      <div className='pt-16'>
        {' '}
        {/* Offset for fixed navigation */}
        <AppRoutes />
      </div>
    </div>
  );
};

export default App;
