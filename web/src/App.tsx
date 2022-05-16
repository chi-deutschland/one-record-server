// routes
import Router from './routes';
// theme
import ThemeProvider from './theme';
// components
import ScrollToTop from './components/ScrollToTop';
import { useEffect, useState } from 'react';
import { fetchToken, onMessageListener } from './utils/firebase';
import toast, { Toaster } from 'react-hot-toast';

// ----------------------------------------------------------------------

export default function App() {
  const [notification, setNotification] = useState({ title: '', body: '' });
  const [isTokenFound, setTokenFound] = useState(false);
  fetchToken(setTokenFound);
  useEffect(() => {
    console.log(isTokenFound);
  }, [isTokenFound]);
  onMessageListener()
    .then((payload: any) => {
      setNotification({ title: payload.notification.title, body: payload.notification.body });
      console.log('Received background message ', payload);

      toast(`${payload.notification.title}: ${payload.notification.body}`);
    })
    .catch((err) => console.log('failed: ', err));

  return (
    <ThemeProvider>
      <Toaster />
      <ScrollToTop />
      <Router />
    </ThemeProvider>
  );
}
