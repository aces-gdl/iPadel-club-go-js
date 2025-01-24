// App.js
import React from 'react';
import { Box } from '@mui/material';
import Routes from './routes';
import { ResponseInterceptor } from 'api/ResponseInterceptor';
//import { ResponseInterceptor } from 'api/ResponseInterceptor';
const App = () => {

  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      minHeight="100vh"
      minWidth={"100vw"}

    >
      <ResponseInterceptor />  {/* Intercepta y maneja las respuestas de la API */}  
      <Routes />
    </Box>

  );
};

export default App;
