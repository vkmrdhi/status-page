import React, { useState } from 'react';
import { Service } from '@/types/types';
import { Button } from '@/components/ui/button';

interface ServiceFormProps {
  onSubmit: (service: Service) => void;
  initialData?: Service;
}

const ServiceForm: React.FC<ServiceFormProps> = ({ onSubmit, initialData }) => {
  const [name, setName] = useState(initialData?.name || '');
  const [description, setDescription] = useState(initialData?.description || '');
  const [status, setStatus] = useState<Service['status']>(initialData?.status || 'operational');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({ id: initialData?.id || Date.now(), name, description, status });
  };

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <div>
        <label htmlFor="name" className="block text-sm font-medium">
          Service Name
        </label>
        <input
          id="name"
          value={name}
          onChange={(e) => setName(e.target.value)}
          className="w-full border p-2 rounded"
          required
        />
      </div>

      <div>
        <label htmlFor="description" className="block text-sm font-medium">
          Description
        </label>
        <textarea
          id="description"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className="w-full border p-2 rounded"
        />
      </div>

      <div>
        <label htmlFor="status" className="block text-sm font-medium">
          Status
        </label>
        <select
          id="status"
          value={status}
          onChange={(e) => setStatus(e.target.value as Service['status'])}
          className="w-full border p-2 rounded"
        >
          <option value="operational">Operational</option>
          <option value="degraded">Degraded Performance</option>
          <option value="partial_outage">Partial Outage</option>
          <option value="major_outage">Major Outage</option>
        </select>
      </div>

      <Button type="submit">Save Service</Button>
    </form>
  );
};

export default ServiceForm;
