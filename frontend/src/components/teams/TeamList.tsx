import React from 'react';
import { Team } from '@/types/types';

interface TeamListProps {
  teams: Team[];
  onEdit: (team: Team) => void;
  onDelete: (id: string) => void;
}

const TeamList: React.FC<TeamListProps> = ({ teams, onEdit, onDelete }) => (
  <div className="space-y-4">
    {teams.length === 0 ? (
      <p className="text-gray-500">No teams available.</p>
    ) : (
      teams.map((team) => (
        <div key={team.id} className="flex justify-between items-center border p-4 rounded">
          <div>
            <h3 className="font-bold">{team.name}</h3>
            <p className="text-sm text-gray-500">{team.description}</p>
          </div>
          <div className="flex items-center space-x-4">
            <button
              onClick={() => onEdit(team)}
              className="text-blue-500 hover:underline"
            >
              Edit
            </button>
            <button
              onClick={() => onDelete(team.id)}
              className="text-red-500 hover:underline"
            >
              Delete
            </button>
          </div>
        </div>
      ))
    )}
  </div>
);

export default TeamList;