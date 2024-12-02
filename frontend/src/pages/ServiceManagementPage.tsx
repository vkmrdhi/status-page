import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import ServiceList from '@/components/services/ServiceList';
import ServiceForm from '@/components/services/ServiceForm';
import {
  getServices,
  createService,
  updateService,
  deleteService,
} from '@/lib/api';
import { Service } from '@/types/types';

const ServiceManagementPage: React.FC = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [selectedService, setSelectedService] = useState<Service | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [loading, setLoading] = useState<boolean>(false); // Add loading state

  const loadServices = async () => {
    setLoading(true); // Start loading
    try {
      const data = await getServices();
      setServices(data);
    } catch (error) {
      console.error('Error fetching services:', error);
    } finally {
      setLoading(false); // End loading
    }
  };

  const handleCreate = async (newService: Service) => {
    await createService(newService);
    await loadServices();
    setShowForm(false);
  };

  const handleUpdate = async (updatedService: Service) => {
    await updateService(updatedService.id, updatedService);
    await loadServices();
    setShowForm(false);
  };

  const handleDelete = async (id: number) => {
    await deleteService(id);
    await loadServices();
  };

  useEffect(() => {
    loadServices();
  }, []);

  return (
    <div className='container mx-auto p-6'>
      <h1 className='text-3xl font-bold mb-6'>Service Management</h1>
      <Button onClick={() => setShowForm(true)} className='mb-4'>
        Add New Service
      </Button>
      {showForm ? (
        <ServiceForm
          service={selectedService}
          onSave={selectedService ? handleUpdate : handleCreate}
          onCancel={() => setShowForm(false)}
        />
      ) : (
        <ServiceList
          services={services}
          onEdit={(service) => {
            setSelectedService(service);
            setShowForm(true);
          }}
          onDelete={handleDelete}
          loading={loading} // Pass the loading prop here
        />
      )}
    </div>
  );
};

export default ServiceManagementPage;
