import PropTypes from 'prop-types';
import { useMemo } from 'react';
// material
import { CssBaseline } from '@mui/material';
import { ThemeProvider as MUIThemeProvider, createTheme, StyledEngineProvider } from '@mui/material/styles';
//
// import palette from './palette';
// import typography from './typography';
// import componentsOverride from './overrides';
// import shadows, { customShadows } from './shadows';

// ----------------------------------------------------------------------



interface Props {
  children: React.ReactNode
}

export default function ThemeProvider({children}:Props) {
  const themeOptions = useMemo(
    () => ({
      shape: { borderRadius: 8 },
    }),
    []
  );

  const theme = createTheme(themeOptions);
  
  return (
    <StyledEngineProvider injectFirst>
      <MUIThemeProvider theme={theme}>
        <CssBaseline />
        {children}
      </MUIThemeProvider>
    </StyledEngineProvider>
  );
}
