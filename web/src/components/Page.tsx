import PropTypes from 'prop-types';
import { Helmet } from 'react-helmet-async';
import React, { forwardRef } from 'react';
// @mui
import { Box, SxProps, Theme } from '@mui/material';

// ----------------------------------------------------------------------
interface Props {
  children: React.ReactNode;
  title: string;
  meta: React.ReactNode;
  sx?: SxProps<Theme>;
  ref?: React.Ref<unknown>;
}

const Page = forwardRef(({ children, title = '', meta, ref, ...other }: Props) => (
  <>
    <Helmet>
      <title>{`${title} | Minimal-UI`}</title>
      {meta}
    </Helmet>

    <Box ref={ref} {...other}>
      {children}
    </Box>
  </>
));

export default Page;
