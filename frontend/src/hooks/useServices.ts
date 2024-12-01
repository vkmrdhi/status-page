import { useEffect, useState } from 'react';
import { getServices } from '@/lib/api';
import { Service } from '@/types/types';

export const useServices = () => {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetch = async () => {
      const data = await getServices();
      setServices(data);
      setLoading(false);
    };
    fetch();
  }, []);

  return { services, loading };
};
