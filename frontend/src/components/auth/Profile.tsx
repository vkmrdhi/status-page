import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';
import { Avatar, AvatarFallback, AvatarImage } from '@/components/ui/avatar';
import { Badge } from '@/components/ui/badge';

const Profile: React.FC = () => {
  const { user, isAuthenticated, isLoading } = useAuth0();

  if (isLoading) {
    return <div>Loading ...</div>;
  }

  if (!isAuthenticated) {
    return null;
  }

  return (
    <Card className="max-w-md mx-auto">
      <CardHeader className="flex flex-row items-center space-x-4">
        <Avatar>
          <AvatarImage 
            src={user?.picture} 
            alt={user?.name || 'User profile'}
          />
          <AvatarFallback>
            {user?.name ? user.name.charAt(0).toUpperCase() : 'UN'}
          </AvatarFallback>
        </Avatar>
        <div>
          <CardTitle>{user?.name}</CardTitle>
          <p className="text-sm text-muted-foreground">{user?.email}</p>
        </div>
      </CardHeader>
      <CardContent>
        <div className="space-y-2">
          <div className="flex justify-between items-center">
            <span>Auth Provider</span>
            <Badge variant="secondary">{user?.sub?.split('|')[0]}</Badge>
          </div>
          {user?.email_verified !== undefined && (
            <div className="flex justify-between items-center">
              <span>Email Verified</span>
              <Badge 
                variant={user.email_verified ? 'default' : 'destructive'}
              >
                {user.email_verified ? 'Yes' : 'No'}
              </Badge>
            </div>
          )}
        </div>
      </CardContent>
    </Card>
  );
};

export default Profile;