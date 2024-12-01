import React from 'react';
import SidebarNavigation from './SidebarNavigation';

interface AppLayoutProps {
  children: React.ReactNode;
}

const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  return (
    <div className='flex h-screen'>
      <SidebarNavigation />
      <div className='flex flex-col flex-1 overflow-hidden'>
        <main className='flex-1 overflow-y-auto p-6 bg-gray-50'>
          {children}
        </main>
      </div>
    </div>
  );
};

export default AppLayout;
