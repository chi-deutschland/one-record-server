import { initializeApp } from 'firebase/app';
import { getMessaging, getToken, onMessage } from 'firebase/messaging';

const firebaseConfig = {
  apiKey: 'AIzaSyBuHGzBCvxDAwN9H_FjJcEqRXWcXAt4seY',
  authDomain: 'one-record.firebaseapp.com',
  projectId: 'one-record',
  storageBucket: 'one-record.appspot.com',
  messagingSenderId: '817384515289',
  appId: '1:817384515289:web:592f1f92a7d819af6240c8',
};

const firebaseApp = initializeApp(firebaseConfig);
const messaging = getMessaging(firebaseApp);

export const fetchToken = (setTokenFound: (arg0: boolean) => void) => {
  return getToken(messaging, {
    vapidKey: 'BGKouv4vrl3LQmhXjLFaGR_SSjiKd0dQ-QYEBzunmH_QKX7BG7Bq9QW1xvCqq6azYi2oIGkm18S1utxXLnJcs3Y',
  })
    .then((currentToken: any) => {
      if (currentToken) {
        console.log('current token for client: ', currentToken);
        setTokenFound(true);
        // Track the token -> client mapping, by sending to backend server
        // show on the UI that permission is secured
      } else {
        console.log('No registration token available. Request permission to generate one.');
        setTokenFound(false);
        // shows on the UI that permission is required
      }
    })
    .catch((err: any) => {
      console.log('An error occurred while retrieving token. ', err);
      // catch error while creating client token
    });
};

export const onMessageListener = () =>
  new Promise((resolve) => {
    onMessage(messaging, (payload: unknown) => {
      console.log('Message received. ', payload);

      resolve(payload);
    });
  });
