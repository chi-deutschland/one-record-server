import { Navigate, useRoutes } from 'react-router-dom';
// layouts
import DashboardLayout from './layouts/dashboard';
//
import Overview from './pages/Overview';
import Subscribe from './pages/Subscribe';

// ----------------------------------------------------------------------

export default function Router() {
  return useRoutes([
    {
      path: 'static',
      element: <DashboardLayout />,
      children: [{ path: '', element: <Overview /> },{ path: 'subscribe', element: <Subscribe /> }],
    },
    { path: '*', element: <Navigate to="/static/404" replace /> },
  ]);
}
