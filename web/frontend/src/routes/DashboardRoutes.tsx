import React from 'react';
import { Routes, Route } from 'react-router-dom';
import DashboardPage from '../dashboard/DashboardPage';
import ReportsPage from '../dashboard/ReportsPage';

const DashboardRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/" element={<DashboardPage />} />
      <Route path="/reports" element={<ReportsPage />} />
    </Routes>
  );
};

export default DashboardRoutes;
