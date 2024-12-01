import React from 'react';
import { Service } from '@/types/types';
import { Card, CardHeader, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';

interface ServiceListProps {
  services: Service[];
  onEdit: (service: Service) => void;
  onDelete: (id: number) => void;
}

const ServiceList: React.FC<ServiceListProps> = ({ services, onEdit, onDelete }) => (
  <Card>
    <CardHeader>
      <h2 className="text-xl font-bold">Services</h2>
    </CardHeader>
    <CardContent>
      {services.length === 0 ? (
        <p className="text-gray-500">No services available.</p>
      ) : (
        services.map((service) => (
          <div key={service.id} className="flex justify-between py-3 border-b last:border-b-0">
            <span className="font-medium">{service.name}</span>
            <div className="flex items-center space-x-2">
              <Badge>{service.status}</Badge>
              <button onClick={() => onEdit(service)} className="text-blue-500 hover:underline">
                Edit
              </button>
              <button onClick={() => onDelete(service.id)} className="text-red-500 hover:underline">
                Delete
              </button>
            </div>
          </div>
        ))
      )}
    </CardContent>
  </Card>
);

export default ServiceList;
