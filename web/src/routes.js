import { Navigate, useRoutes } from 'react-router-dom';
// layouts
import DashboardLayout from './layouts/dashboard';
//
import Overview from './pages/Overview';

// ----------------------------------------------------------------------

export default function Router() {
  return useRoutes([
    {
      path: '/',
      element: <DashboardLayout />,
      children: [{ path: '/', element: <Overview /> }],
    },
    { path: '*', element: <Navigate to="/404" replace /> },
  ]);
}
