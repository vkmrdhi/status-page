import { useState, useEffect } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import {
  AlertCircle,
  CheckCircle2,
  XCircle,
  AlertTriangle,
} from 'lucide-react';
import { useAuth0 } from '@auth0/auth0-react';

interface Service {
  id: number;
  name: string;
  status: keyof typeof StatusBadgeMap;
}

interface Incident {
  id: number;
  title: string;
  service: string;
  status: string;
  updatedAt: string;
}

const StatusBadgeMap = {
  operational: {
    icon: <CheckCircle2 className='text-green-500' />,
    label: 'Operational',
    className: 'bg-green-100 text-green-800',
  },
  degraded: {
    icon: <AlertTriangle className='text-yellow-500' />,
    label: 'Degraded Performance',
    className: 'bg-yellow-100 text-yellow-800',
  },
  partial_outage: {
    icon: <AlertCircle className='text-orange-500' />,
    label: 'Partial Outage',
    className: 'bg-orange-100 text-orange-800',
  },
  major_outage: {
    icon: <XCircle className='text-red-500' />,
    label: 'Major Outage',
    className: 'bg-red-100 text-red-800',
  },
};

const PublicStatusPage: React.FC = () => {
  const { getAccessTokenSilently } = useAuth0();

  const [services, setServices] = useState<Service[]>([]);
  const [incidents, setIncidents] = useState<Incident[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [socket, setSocket] = useState<WebSocket | null>(null);

  // Fetch initial status data
  const fetchStatusData = async () => {
    try {
      const mockServices: Service[] = [
        { id: 1, name: 'Website', status: 'operational' },
        { id: 2, name: 'API', status: 'degraded' },
        { id: 3, name: 'Database', status: 'major_outage' },
      ];

      const mockIncidents: Incident[] = [
        {
          id: 1,
          title: 'API Performance Issues',
          service: 'API',
          status: 'investigating',
          updatedAt: new Date().toISOString(),
        },
      ];

      setServices(mockServices);
      setIncidents(mockIncidents);
    } catch (error) {
      console.error('Failed to fetch status', error);
    } finally {
      setLoading(false);
    }
  };

  // Initialize WebSocket connection
  const setupWebSocket = async () => {
    try {
      const token = await getAccessTokenSilently();
      const ws = new WebSocket(
        `ws://localhost:8080/status-updates?token=${token}`
      );

      ws.onopen = () => console.log('Connected to WebSocket');
      ws.onmessage = (event: MessageEvent) => {
        const data = JSON.parse(event.data);
        if (data.type === 'service-update') {
          setServices((prevServices) =>
            prevServices.map((service) =>
              service.id === data.service.id ? data.service : service
            )
          );
        }

        if (data.type === 'incident-update') {
          setIncidents((prevIncidents) =>
            prevIncidents.map((incident) =>
              incident.id === data.incident.id ? data.incident : incident
            )
          );
        }
      };
      ws.onerror = (error: Event) => console.error('WebSocket error:', error);
      ws.onclose = () => console.log('WebSocket connection closed');

      setSocket(ws);
    } catch (error) {
      console.error('Failed to connect to WebSocket', error);
    }
  };

  useEffect(() => {
    fetchStatusData();
  }, []);

  useEffect(() => {
    setupWebSocket();
    return () => {
      socket?.close();
    };
  }, []);

  // Optionally, send a message via the WebSocket
  const sendMessage = (message: string) => {
    if (socket?.readyState === WebSocket.OPEN) {
      socket.send(message);
    }
  };

  if (loading) return <div>Loading status...</div>;

  return (
    <div className='container mx-auto p-6 max-w-4xl'>
      <h1 className='text-3xl font-bold mb-6'>System Status</h1>

      {/* Services Status */}
      <Card className='mb-6'>
        <CardHeader>
          <CardTitle>Service Status</CardTitle>
        </CardHeader>
        <CardContent>
          {services.map((service) => {
            const { icon, label, className } = StatusBadgeMap[service.status];
            return (
              <div
                key={service.id}
                className='flex justify-between items-center py-3 border-b last:border-b-0'
              >
                <span className='font-medium'>{service.name}</span>
                <div className='flex items-center space-x-2'>
                  {icon}
                  <Badge className={className}>{label}</Badge>
                </div>
              </div>
            );
          })}
        </CardContent>
      </Card>

      <Card>
        <CardHeader>
          <CardTitle>Active Incidents</CardTitle>
        </CardHeader>
        <CardContent>
          {incidents.length === 0 ? (
            <p className='text-gray-500'>No active incidents</p>
          ) : (
            incidents.map((incident) => (
              <div key={incident.id} className='py-3 border-b last:border-b-0'>
                <div className='flex justify-between items-center'>
                  <div>
                    <h3 className='font-semibold'>{incident.title}</h3>
                    <p className='text-sm text-gray-500'>
                      {incident.service} | {incident.status}
                    </p>
                  </div>
                  <span className='text-sm text-gray-500'>
                    {new Date(incident.updatedAt).toLocaleString()}
                  </span>
                </div>
              </div>
            ))
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default PublicStatusPage;
