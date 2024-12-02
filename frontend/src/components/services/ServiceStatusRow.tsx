import React from 'react';
import { Badge } from '@/components/ui/badge';
import { ServiceStatusBadgeMap } from '@/lib/constants';

interface ServiceStatusRowProps {
  name: string;
  status: string;
  onClick: (newStatus: string) => void;
}

const ServiceStatusRow: React.FC<ServiceStatusRowProps> = ({
  name,
  status,
  onClick,
}) => {
  return (
    <div className='flex flex-col p-4 border rounded shadow-sm space-y-2'>
      <div className='flex justify-between items-center'>
        <span className='font-medium'>{name}</span>
      </div>
      <div className='flex space-x-2'>
        {(
          Object.keys(ServiceStatusBadgeMap) as Array<keyof typeof ServiceStatusBadgeMap>
        ).map((key) => (
          <Badge
            key={key}
            onClick={() => onClick(key)}
            className={`cursor-pointer space-x-2 ${
              key === status
                ? ServiceStatusBadgeMap[key].className
                : 'bg-gray-200 text-gray-500'
            } ${
              key === status ? 'opacity-100' : 'opacity-50 hover:opacity-75'
            }`}
          >
            {key === status && <span>{ServiceStatusBadgeMap[key].icon}</span>}
            <span>{ServiceStatusBadgeMap[key].label}</span>
          </Badge>
        ))}
      </div>
    </div>
  );
};

export default ServiceStatusRow;
