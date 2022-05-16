import PropTypes from 'prop-types';
import { Link as RouterLink } from 'react-router-dom';
// @mui
import { SxProps, Theme, useTheme } from '@mui/material/styles';
import { Box } from '@mui/material';

// ----------------------------------------------------------------------

interface Props {
  disabledLink?: boolean;
  sx?: SxProps<Theme>;
}

export default function Logo({ disabledLink = false, sx }: Props) {
  const theme = useTheme();

  const logo = (
    <Box
      component="img"
      src="https://upload.wikimedia.org/wikipedia/de/e/ee/Fraunhofer-Gesellschaft_2009_logo.svg"
      sx={{ width: 226, height: 40, ...sx }}
    />
  );

  if (disabledLink) {
    return <>{logo}</>;
  }

  return <RouterLink to="/">{logo}</RouterLink>;
}
