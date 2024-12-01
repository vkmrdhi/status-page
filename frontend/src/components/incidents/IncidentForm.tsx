import React, { useState } from 'react';
import { Incident } from '@/types/types';
import { Button } from '@/components/ui/button';

interface IncidentFormProps {
  onSubmit: (incident: Incident) => void;
  initialData?: Incident;
  onCancel: () => void;
}

const IncidentForm: React.FC<IncidentFormProps> = ({
  onSubmit,
  initialData,
}) => {
  const [title, setTitle] = useState(initialData?.title || '');
  const [description, setDescription] = useState(
    initialData?.description || ''
  );
  const [status, setStatus] = useState<Incident['status']>(
    initialData?.status || 'investigating'
  );
  const [service, setService] = useState(initialData?.service || '');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSubmit({
      id: initialData?.id || Date.now(),
      title,
      description,
      status,
      service,
      updatedAt: new Date().toISOString(),
      createdAt: initialData?.createdAt || new Date().toISOString(),
    });
  };

  return (
    <form onSubmit={handleSubmit} className='space-y-4'>
      <div>
        <label htmlFor='title' className='block text-sm font-medium'>
          Title
        </label>
        <input
          id='title'
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className='w-full border p-2 rounded'
          required
        />
      </div>

      <div>
        <label htmlFor='description' className='block text-sm font-medium'>
          Description
        </label>
        <textarea
          id='description'
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          className='w-full border p-2 rounded'
        />
      </div>

      <div>
        <label htmlFor='service' className='block text-sm font-medium'>
          Affected Service
        </label>
        <input
          id='service'
          value={service}
          onChange={(e) => setService(e.target.value)}
          className='w-full border p-2 rounded'
          required
        />
      </div>

      <div>
        <label htmlFor='status' className='block text-sm font-medium'>
          Status
        </label>
        <select
          id='status'
          value={status}
          onChange={(e) => setStatus(e.target.value as Incident['status'])}
          className='w-full border p-2 rounded'
        >
          <option value='investigating'>Investigating</option>
          <option value='identified'>Identified</option>
          <option value='monitoring'>Monitoring</option>
          <option value='resolved'>Resolved</option>
        </select>
      </div>

      <Button type='submit'>Save Incident</Button>
    </form>
  );
};

export default IncidentForm;
