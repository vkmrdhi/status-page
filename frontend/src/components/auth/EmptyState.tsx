import React from 'react';
import { Package } from 'lucide-react';

interface EmptyStateProps {
  icon?: React.ElementType;
  title?: string;
  description?: string;
  action?: React.ReactNode;
}

const EmptyState: React.FC<EmptyStateProps> = ({
  icon: Icon = Package,
  title = 'No items found',
  description = 'There are currently no items to display.',
  action
}) => {
  return (
    <div className="flex flex-col items-center justify-center p-6 bg-gray-50 rounded-lg text-center">
      <Icon className="w-16 h-16 text-muted-foreground mb-4" />
      <h2 className="text-xl font-semibold text-gray-700 mb-2">{title}</h2>
      <p className="text-muted-foreground mb-4">{description}</p>
      {action && (
        <div className="mt-4">
          {action}
        </div>
      )}
    </div>
  );
};

export default EmptyState;