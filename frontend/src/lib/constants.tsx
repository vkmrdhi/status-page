import {
  CheckCircle2,
  AlertTriangle,
  AlertCircle,
  XCircle,
  Home,
  Server,
  Users,
} from 'lucide-react';

// Status Badge Map
export const StatusBadgeMap = {
  operational: {
    icon: <CheckCircle2 className='text-green-500' />,
    label: 'Operational',
    className: 'bg-green-100 text-green-800',
  },
  degraded: {
    icon: <AlertTriangle className='text-yellow-500' />,
    label: 'Degraded Performance',
    className: 'bg-yellow-100 text-yellow-800',
  },
  partial_outage: {
    icon: <AlertCircle className='text-orange-500' />,
    label: 'Partial Outage',
    className: 'bg-orange-100 text-orange-800',
  },
  major_outage: {
    icon: <XCircle className='text-red-500' />,
    label: 'Major Outage',
    className: 'bg-red-100 text-red-800',
  },
};

export const NavItems = [
  { 
    name: 'Dashboard', 
    path: '/dashboard', 
    icon: Home
  },
  { 
    name: 'Services', 
    path: '/services', 
    icon: Server
  },
  { 
    name: 'Incidents', 
    path: '/incidents', 
    icon: AlertTriangle 
  },
  { 
    name: 'Team', 
    path: '/team', 
    icon: Users
  }
];