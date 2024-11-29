import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Button } from '@/components/ui/button';
import Profile from '@/components/auth/Profile';
import LogoutButton from '@/components/auth/LogoutButton';
import { Plus, Server } from 'lucide-react';

const DashboardPage: React.FC = () => {
  const { user } = useAuth0();

  return (
    <div className="container mx-auto p-6">
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {/* Profile Section */}
        <div className="md:col-span-1">
          <Profile />
          <Card className="mt-4">
            <CardHeader>
              <CardTitle>Actions</CardTitle>
            </CardHeader>
            <CardContent className="space-y-2">
              <Button className="w-full" variant="outline">
                <Plus className="mr-2 h-4 w-4" />
                Add New Service
              </Button>
              <Button className="w-full" variant="outline">
                <Server className="mr-2 h-4 w-4" />
                Manage Services
              </Button>
              <LogoutButton />
            </CardContent>
          </Card>
        </div>

        {/* Services Overview */}
        <div className="md:col-span-2">
          <Card>
            <CardHeader>
              <CardTitle>Your Services</CardTitle>
            </CardHeader>
            <CardContent>
              {/* TODO: Implement service list */}
              <p className="text-muted-foreground">
                No services configured. Click "Add New Service" to get started.
              </p>
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
};

export default DashboardPage;