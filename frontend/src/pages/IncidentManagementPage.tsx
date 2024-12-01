import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import IncidentsList from '@/components/incidents/IncidentsList';
import IncidentForm from '@/components/incidents/IncidentForm';
import {
  getIncidents,
  createIncident,
  updateIncident,
  deleteIncident,
} from '@/lib/api';
import { Incident } from '@/types/types';

const IncidentManagementPage: React.FC = () => {
  const [incidents, setIncidents] = useState<Incident[]>([]);
  const [selectedIncident, setSelectedIncident] = useState<
    Incident | undefined
  >();
  const [showForm, setShowForm] = useState(false);

  const loadIncidents = async () => {
    const data = await getIncidents();
    setIncidents(data);
  };

  const handleCreate = async (newIncident: Incident) => {
    await createIncident(newIncident);
    await loadIncidents();
    setShowForm(false);
  };

  const handleUpdate = async (updatedIncident: Incident) => {
    await updateIncident(updatedIncident.id, updatedIncident);
    await loadIncidents();
    setShowForm(false);
  };

  const handleDelete = async (id: string) => {
    await deleteIncident(id);
    await loadIncidents();
  };

  useEffect(() => {
    loadIncidents();
  }, []);

  return (
    <div className='container mx-auto p-6'>
      <h1 className='text-3xl font-bold mb-6'>Incident Management</h1>
      <Button onClick={() => setShowForm(true)} className='mb-4'>
        Add New Incident
      </Button>
      {showForm && (
        <IncidentForm
          initialData={selectedIncident}
          onSubmit={selectedIncident ? handleUpdate : handleCreate}
          onCancel={() => setShowForm(false)}
        />
      )}
      <IncidentsList
        incidents={incidents}
        onEdit={(incident: Incident) => {
          setSelectedIncident(incident);
          setShowForm(true);
        }}
        onDelete={handleDelete}
      />
    </div>
  );
};

export default IncidentManagementPage;
