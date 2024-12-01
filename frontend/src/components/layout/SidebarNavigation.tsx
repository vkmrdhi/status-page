import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import { 
  Settings 
} from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils';
import { NavItems } from '@/lib/constants';


const SidebarNavigation: React.FC = () => {
  const location = useLocation();

  return (
    <div className="hidden md:flex flex-col w-64 bg-white border-r h-screen p-4">
      <Link 
        to="/" 
        className="text-2xl font-bold text-primary mb-8 text-center"
      >
        StatusHub
      </Link>

      <nav className="space-y-2 flex-grow">
        {NavItems.map((item) => {
          const Icon = item.icon;
          const isActive = location.pathname.startsWith(item.path);

          return (
            <Link key={item.path} to={item.path}>
              <Button
                variant={isActive ? 'secondary' : 'ghost'}
                className={cn(
                  'w-full justify-start space-x-2',
                  isActive && 'bg-secondary text-secondary-foreground'
                )}
              >
                <Icon className="w-4 h-4" />
                <span>{item.name}</span>
              </Button>
            </Link>
          );
        })}
      </nav>

      <div className="border-t pt-4 space-y-2">
        <Button variant="outline" className="w-full">
          Invite Team
        </Button>
        <Button variant="outline" className="w-full">
          <Settings className="w-4 h-4 mr-2" />
          Account Settings
        </Button>
      </div>
    </div>
  );
};

export default SidebarNavigation;