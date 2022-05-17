import PropTypes from 'prop-types';
// icons
import { Icon, IconProps } from '@iconify/react';
// @mui
import { Box, Button, Stack, SxProps, TextField, Theme, Typography } from '@mui/material';
import { useState } from 'react';
import { SubscribeTopic } from '../network/notification';
import toast from 'react-hot-toast';
import { LoadingButton } from '@mui/lab';

// ----------------------------------------------------------------------

interface Props {}

export default function Subscribe({}: Props) {
  const [topic, setTopic] = useState<string>('');
  const [loading, setLoading] = useState<boolean>(false);
  const onClick = async () => {
    const token = localStorage.getItem('pnk');
    if (token == undefined) {
      toast.error('Token Not Found');
    } else {
      await SubscribeTopic(token, topic)
        .then(() => {
          setLoading(false);
          toast.success('Accepted');
        })
        .catch(() => {
          setLoading(false);
          toast.error('Cannot subscribe the topic');
        });
    }
  };

  return (
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
          Contained
        </LoadingButton>
      </Stack>
    </Box>
  );
}
