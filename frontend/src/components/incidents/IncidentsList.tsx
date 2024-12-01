import React from 'react';
import { Incident } from '@/types/types';

interface IncidentListProps {
  incidents: Incident[];
  onEdit: (incident: Incident) => void;
  onDelete: (id: string) => void;
}

const IncidentList: React.FC<IncidentListProps> = ({ incidents }) => (
  <div>
    {incidents.length === 0 ? (
      <p className='text-gray-500'>No active incidents.</p>
    ) : (
      incidents.map((incident) => (
        <div key={incident.id} className='py-4 border-b last:border-b-0'>
          <h3 className='font-bold'>{incident.title}</h3>
          <p className='text-sm text-gray-500'>{incident.description}</p>
          <span className='text-sm text-gray-400'>
            {new Date(incident.updatedAt).toLocaleString()}
          </span>
        </div>
      ))
    )}
  </div>
);

export default IncidentList;
