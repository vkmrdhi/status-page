import React from 'react';

interface MainNavigationProps {
  onToggleSidebar: () => void;
}

const MainNavigation: React.FC<MainNavigationProps> = ({ onToggleSidebar }) => {
  return (
    <header className="bg-white shadow-md p-4 flex justify-between items-center">
      <button onClick={onToggleSidebar} className="text-gray-700">
        Toggle Sidebar
      </button>
      <span className="text-xl font-bold">Application Title</span>
    </header>
  );
};

export default MainNavigation;
