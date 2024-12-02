import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import {
  getIncidents,
  createIncident,
  updateIncident,
  deleteIncident,
} from '@/lib/api';
import { Incident } from '@/types/types';
import IncidentForm from '@/components/incidents/IncidentForm';
import IncidentList from '@/components/incidents/IncidentsList';

const IncidentManagementPage: React.FC = () => {
  const [incidents, setIncidents] = useState<Incident[]>([]);
  const [selectedIncident, setSelectedIncident] = useState<Incident | null>(
    null
  );
  const [showForm, setShowForm] = useState(false);
  const [loading, setLoading] = useState<boolean>(false); // Add loading state

  const loadIncidents = async () => {
    setLoading(true); // Start loading
    try {
      const data = await getIncidents();
      setIncidents(data);
    } catch (error) {
      console.error('Error fetching incidents:', error);
    } finally {
      setLoading(false); // End loading
    }
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

  const handleDelete = async (id: number) => {
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
      {showForm ? (
        <IncidentForm
          incident={selectedIncident}
          onSave={selectedIncident ? handleUpdate : handleCreate}
          onCancel={() => setShowForm(false)}
        />
      ) : (
        <IncidentList
          incidents={incidents}
          onEdit={(incident) => {
            console.log(incident);
            setSelectedIncident(incident);
            setShowForm(true);
          }}
          onDelete={handleDelete}
          loading={loading} // Pass the loading state here
        />
      )}
    </div>
  );
};

export default IncidentManagementPage;
