// routes
import Router from './routes';
// theme
import ThemeProvider from './theme';
// components
import ScrollToTop from './components/ScrollToTop';
import { useEffect, useState } from 'react';
import { fetchToken, onMessageListener } from './utils/firebase';
import toast, { Toaster } from 'react-hot-toast';
import { AppCtx } from './pages/Subscribe';
import React from 'react';

export const LanguageContext = React.createContext({
  language: { title: '', body: '' },
  setLanguage: () => {},
});

export interface XXInterface {
  title: string;
  body: string;
}
// ----------------------------------------------------------------------
export default function App() {
  const setLanguage = (language: XXInterface) => {
    setNotification(language);
  };

  const [notification, setNotification] = useState({ title: '', body: '' });
  const [isTokenFound, setTokenFound] = useState(false);
  useEffect(() => {
    async function fetchMyAPI() {
      await fetchToken(setTokenFound);
    }

    fetchMyAPI();
  }, []);
  const [n, setN] = React.useContext(AppCtx);

  useEffect(() => {
    async function fetchMyAPI() {
      await fetchToken(setTokenFound);
    }
    fetchMyAPI();
    console.log(isTokenFound);
  }, [isTokenFound]);

  return (
    <ThemeProvider>
      <Toaster />
      <ScrollToTop />
      <Router />
    </ThemeProvider>
  );
}
