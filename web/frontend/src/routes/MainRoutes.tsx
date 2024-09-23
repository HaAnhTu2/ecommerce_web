import React from 'react';
import { Routes, Route } from 'react-router-dom';
import UserRoutes from './UserRoutes';
import DashboardRoutes from './DashboardRoutes';

const MainRoutes: React.FC = () => {
  return (
    <Routes>
      <Route path="/*" element={<UserRoutes />} />
      <Route path="/dashboard/*" element={<DashboardRoutes />} />
    </Routes>
  );
};

export default MainRoutes;
