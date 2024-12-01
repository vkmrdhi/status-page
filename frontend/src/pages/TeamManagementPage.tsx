import React, { useState, useEffect } from 'react';
import { Button } from '@/components/ui/button';
import TeamList from '@/components/teams/TeamList';
import TeamForm from '@/components/teams/TeamForm';
import { getTeams, createTeam, updateTeam, deleteTeam } from '@/lib/api';
import { Team } from '@/types/types';

const TeamManagementPage: React.FC = () => {
  const [teams, setTeams] = useState<Team[]>([]);
  const [selectedTeam, setSelectedTeam] = useState<Team | null>(null);
  const [showForm, setShowForm] = useState(false);

  const loadTeams = async () => {
    const data = await getTeams();
    setTeams(data);
  };

  const handleCreate = async (newTeam: Team) => {
    await createTeam(newTeam);
    await loadTeams();
    setShowForm(false);
  };

  const handleUpdate = async (updatedTeam: Team) => {
    await updateTeam(updatedTeam.id, updatedTeam);
    await loadTeams();
    setShowForm(false);
  };

  const handleDelete = async (id: number) => {
    await deleteTeam(id);
    await loadTeams();
  };

  useEffect(() => {
    loadTeams();
  }, []);

  return (
    <div className='container mx-auto p-6'>
      <h1 className='text-3xl font-bold mb-6'>Team Management</h1>
      <Button onClick={() => setShowForm(true)} className='mb-4'>
        Add New Team
      </Button>
      {showForm && (
        <TeamForm
          team={selectedTeam}
          onSave={selectedTeam ? handleUpdate : handleCreate}
          onCancel={() => setShowForm(false)}
        />
      )}
      <TeamList
        teams={teams}
        onEdit={(team: Team) => {
          setSelectedTeam(team);
          setShowForm(true);
        }}
        onDelete={handleDelete}
      />
    </div>
  );
};

export default TeamManagementPage;
