import { useState, useEffect } from 'react';
import { Card, CardHeader, CardTitle, CardContent } from '@/components/ui/card';
import { Badge } from '@/components/ui/badge';
import { AlertCircle, CheckCircle2, XCircle, AlertTriangle } from 'lucide-react';

const StatusBadgeMap = {
  operational: { 
    icon: <CheckCircle2 className="text-green-500" />, 
    label: 'Operational', 
    className: 'bg-green-100 text-green-800' 
  },
  degraded: { 
    icon: <AlertTriangle className="text-yellow-500" />, 
    label: 'Degraded Performance', 
    className: 'bg-yellow-100 text-yellow-800' 
  },
  partial_outage: { 
    icon: <AlertCircle className="text-orange-500" />, 
    label: 'Partial Outage', 
    className: 'bg-orange-100 text-orange-800' 
  },
  major_outage: { 
    icon: <XCircle className="text-red-500" />, 
    label: 'Major Outage', 
    className: 'bg-red-100 text-red-800' 
  }
};

const PublicStatusPage = () => {
  const [services, setServices] = useState([]);
  const [incidents, setIncidents] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    // TODO: Replace with actual API call
    const fetchStatusData = async () => {
      try {
        // Simulate API call
        const mockServices = [
          { id: 1, name: 'Website', status: 'operational' },
          { id: 2, name: 'API', status: 'degraded' },
          { id: 3, name: 'Database', status: 'major_outage' }
        ];
        
        const mockIncidents = [
          { 
            id: 1, 
            title: 'API Performance Issues', 
            service: 'API', 
            status: 'investigating', 
            updatedAt: new Date().toISOString() 
          }
        ];

        setServices(mockServices);
        setIncidents(mockIncidents);
        setLoading(false);
      } catch (error) {
        console.error('Failed to fetch status', error);
        setLoading(false);
      }
    };

    fetchStatusData();
  }, []);

  if (loading) return <div>Loading status...</div>;

  return (
    <div className="container mx-auto p-6 max-w-4xl">
      <h1 className="text-3xl font-bold mb-6">System Status</h1>
      
      {/* Services Status */}
      <Card className="mb-6">
        <CardHeader>
          <CardTitle>Service Status</CardTitle>
        </CardHeader>
        <CardContent>
          {services.map(service => {
            const { icon, label, className } = StatusBadgeMap[service.status];
            return (
              <div key={service.id} className="flex justify-between items-center py-3 border-b last:border-b-0">
                <span className="font-medium">{service.name}</span>
                <div className="flex items-center space-x-2">
                  {icon}
                  <Badge className={className}>{label}</Badge>
                </div>
              </div>
            );
          })}
        </CardContent>
      </Card>

      {/* Active Incidents */}
      <Card>
        <CardHeader>
          <CardTitle>Active Incidents</CardTitle>
        </CardHeader>
        <CardContent>
          {incidents.length === 0 ? (
            <p className="text-gray-500">No active incidents</p>
          ) : (
            incidents.map(incident => (
              <div key={incident.id} className="py-3 border-b last:border-b-0">
                <div className="flex justify-between items-center">
                  <div>
                    <h3 className="font-semibold">{incident.title}</h3>
                    <p className="text-sm text-gray-500">
                      {incident.service} | {incident.status}
                    </p>
                  </div>
                  <span className="text-sm text-gray-500">
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