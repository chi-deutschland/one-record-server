import PropTypes from 'prop-types';
// icons
import { Icon, IconProps } from '@iconify/react';
// @mui
import { Box, SxProps, Theme } from '@mui/material';

// ----------------------------------------------------------------------

interface Props {
  icon: IconProps | any;
  sx?: SxProps<Theme>;
  width?: any;
  height?: any;
}

export default function Iconify({ icon, sx, ...other }: Props) {
  return <Box component={Icon} icon={icon} sx={{ ...sx }} {...other} />;
}
