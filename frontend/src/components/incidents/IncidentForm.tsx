import React, { useState, useEffect } from 'react';
import { Incident } from '@/types/types';
import { Button } from '@/components/ui/button';

interface IncidentFormProps {
  onSave: (incident: Incident) => void;
  initialData?: Incident;
  onCancel: () => void;
}

const IncidentForm: React.FC<IncidentFormProps> = ({
  onSave,
  onCancel,
  initialData,
}) => {
  const [title, setTitle] = useState<string>(initialData?.title || '');
  const [description, setDescription] = useState<string>(
    initialData?.description || ''
  );
  const [status, setStatus] = useState<Incident['status']>(
    initialData?.status || 'investigating'
  );
  const [service, setService] = useState<string>(initialData?.service_id || '');

  // Reset the form when initialData changes
  useEffect(() => {
    if (initialData) {
      setTitle(initialData.title);
      setDescription(initialData.description);
      setStatus(initialData.status);
      setService(initialData.service_id);
    }
  }, [initialData]);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onSave({
      id: initialData?.id || Date.now().toString(),
      title,
      description,
      status,
      service_id: service,
      UpdatedAt: new Date().toISOString(),
      CreatedAt: initialData?.CreatedAt || new Date().toISOString(),
      priority: 'low'
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

      <div className='flex space-x-2'>
        <Button type='submit'>Save Incident</Button>
        <Button type='button' onClick={onCancel} variant='outline'>
          Cancel
        </Button>
      </div>
    </form>
  );
};

export default IncidentForm;
