export interface Service {
  id: string;
  name: string;
  description?: string; // Optional description of the service
  status: 'operational' | 'degraded' | 'partial_outage' | 'major_outage';
  createdAt?: string;
  updatedAt?: string;
}

export interface Incident {
  id: string;
  title: string;
  description: string;
  service_id: string; // Name or ID of the associated service
  status: 'investigating' | 'active' | 'monitoring' | 'resolved';
  CreatedAt: string;
  UpdatedAt: string;
  ResolvedAt?: string; 
  priority: 'low' | 'medium' | 'high' | 'critical';
}

export interface User {
  id: string;
  email: string;
  name: string;
  role: 'admin' | 'user'; // Define roles as needed
  organizationId: string; // ID of the organization the user belongs to
}

export interface Organization {
  id: string;
  name: string;
  description?: string;
  createdAt: string;
  updatedAt: string;
}

export interface Team {
  id: string;
  name: string;
  description?: string;
  members: User[]; // Array of user objects
  createdAt: string;
  updatedAt: string;
}

export interface WebSocketMessage {
  type: 'service-update' | 'incident-update' | 'team-update'; // Define possible message types
  payload: unknown; // Payload structure depends on the message type
}

export interface ApiResponse<T> {
  data: T;
  message?: string; // Optional message
  error?: string; // Error details if the request fails
}

export interface AuthToken {
  token: string;
  expiresAt: string; // ISO timestamp for token expiry
}

export interface AuthContextType {
  user: User | null;
  token: string | null;
  login: () => Promise<void>;
  logout: () => void;
  isAuthenticated: boolean;
}

export interface Status {
  serviceId: string;
  currentStatus: 'operational' | 'degraded' | 'partial_outage' | 'major_outage';
  lastUpdated: string; // ISO timestamp for the last update
}

export interface TeamMember {
  id: string;
  name: string;
  email: string;
  role: 'admin' | 'member';
}

export interface TeamManagement {
  teamId: string;
  name: string;
  members: TeamMember[];
}
