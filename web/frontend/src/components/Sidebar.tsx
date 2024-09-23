import React from 'react';
import { Link } from 'react-router-dom';

const Sidebar: React.FC = () => {
  return (
    <aside>
      <Link to="/dashboard">Dashboard Home</Link>
      <Link to="/dashboard/reports">Reports</Link>
    </aside>
  );
};

export default Sidebar;
