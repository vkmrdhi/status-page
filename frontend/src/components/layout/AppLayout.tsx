import React from 'react';
import Header from './Header';
import SidebarNavigation from './SidebarNavigation';

interface AppLayoutProps {
  children: React.ReactNode;
}

const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  return (
    <div className='h-screen flex flex-col'>
      <Header />

      <div className='flex flex-grow overflow-hidden'>
        <SidebarNavigation />

        <div className='flex-grow bg-gray-100 p-4 overflow-auto'>
          {children}
        </div>
      </div>
    </div>
  );
};

export default AppLayout;
