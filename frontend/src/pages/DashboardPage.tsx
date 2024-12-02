import React, { useState, useEffect } from 'react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import { Server } from 'lucide-react';
import { useNavigate } from 'react-router-dom';
import LoadingSpinner from '@/components/common/LoadingSpinner';
import { getServices, updateService } from '@/lib/api';
import ServiceStatusRow from '@/components/services/ServiceStatusRow';

interface Service {
  id: string;
  name: string;
  status: string;
}

const DashboardPage: React.FC = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchServices = async () => {
      try {
        const data = await getServices();
        setServices(data);
      } catch (error) {
        console.error('Error fetching services:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchServices();
  }, []);

  const handleStatusChange = async (id: string, newStatus: string) => {
    try {
      const updatedService = await updateService(id, { status: newStatus });
      setServices((prevServices) =>
        prevServices.map((service) =>
          service.id === id ? { ...service, ...updatedService } : service
        )
      );
    } catch (error) {
      console.error('Error updating service status:', error);
    }
  };

  return (
    <div className='container mx-auto p-6'>
      <div className='grid grid-cols-1 gap-6'>
        <div className='flex justify-end'>
          <Button onClick={() => navigate('/services')}>
            <Server className='mr-2 h-4 w-4' />
            Manage Services
          </Button>
        </div>
        <Card>
          <CardHeader>
            <CardTitle>Your Services</CardTitle>
          </CardHeader>
          <CardContent>
            {loading ? (
              <LoadingSpinner message='Loading services...' />
            ) : services.length === 0 ? (
              <p className='text-muted-foreground'>
                No services configured. Click "Add New Service" to get started.
              </p>
            ) : (
              <ul className='space-y-4'>
                {services.map((service) => (
                  <ServiceStatusRow
                    key={service.id}
                    name={service.name}
                    status={service.status}
                    onClick={(newStatus) =>
                      handleStatusChange(service.id, newStatus)
                    }
                  />
                ))}
              </ul>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default DashboardPage;
