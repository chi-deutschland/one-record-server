import PropTypes from 'prop-types';
// icons
import { Icon, IconProps } from '@iconify/react';
// @mui
import { Alert, AlertColor, Box, Button, Stack, SxProps, TextField, Theme, Typography } from '@mui/material';
import { useEffect, useState } from 'react';
import { SubscribeTopic } from '../network/notification';
import toast from 'react-hot-toast';
import { LoadingButton } from '@mui/lab';
import React from 'react';
import { onMessageListener } from '../utils/firebase';

// ----------------------------------------------------------------------

interface Props {}

export default function Subscribe({}: Props) {
  const [topic, setTopic] = useState<string>('');
  const [notifs, setNotifs] = useState<AppContextInterface[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const onClick = async () => {
    const token = localStorage.getItem('pnk');
    if (token == undefined) {
      toast.error('Token Not Found');
    } else {
      await SubscribeTopic(token, topic)
        .then((a) => {
          setLoading(false);
          toast.success('Accepted');
          const d = JSON.parse(JSON.stringify(notifs));
          d.push({ name: 'Topic: ' + topic, url: `successfully Subscribed`, severity: 'success' });
          setNotifs(d);
        })
        .catch(() => {
          setLoading(false);
          toast.error('Cannot subscribe the topic');
        });
    }
  };

  onMessageListener()
    .then((payload: any) => {
      // setNotification({ title: payload.notification.title, body: payload.notification.body });
      console.log('Received background message ', payload);
      const d = JSON.parse(JSON.stringify(notifs));
      d.push({
        name: 'Topic: ' + payload.notification.title,
        url: 'New Security Declaration on piece ' + payload.notification.body,
        severity: 'info',
      });
      // appContext = d;
      setNotifs(d);
      console.log(d);

      toast(`${payload.notification.title}: ${payload.notification.body}`);
    })
    .catch((err) => console.log('failed: ', err));

  useEffect(() => {
    console.log('----____>', notifs);
  }, [notifs]);

  return (
    <>
      <Box
        sx={{
          m: '0 auto',
          maxWidth: 400,
        }}
      >
        <Stack direction="column" spacing={2}>
          <Typography>Enter your favorite topic</Typography>
          <TextField
            value={topic}
            onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
              setTopic(event.target.value as string);
            }}
            label="Topic"
            variant="outlined"
          />
          <LoadingButton
            loading={loading}
            variant="contained"
            onClick={() => {
              onClick();
            }}
          >
            Subscribe
          </LoadingButton>
        </Stack>
      </Box>
      <Stack sx={{ m: 2 }} spacing={2}>
        {notifs.map((s) => {
          return <Alert severity={s.severity}>{s.name + ' ' + s.url}</Alert>;
        })}
      </Stack>
    </>
  );
}

interface AppContextInterface {
  name: string;
  url: string;
  severity: AlertColor;
}

export const AppCtx = React.createContext<AppContextInterface[]>([]);
