// App.js
import React from 'react';
import { Box } from '@mui/material';
import Routes from './routes';
import axios from 'axios';
//import { ResponseInterceptor } from 'api/ResponseInterceptor';
const App = () => {
/* 
  if (process.env.NODE_ENV === "production") {
    axios.defaults.baseURL = process.env.REACT_APP_API_BASE_URL ;
    console.log('Prod -> Base URL :', axios.defaults.baseURL)
  } else {
    //axios.defaults.baseURL = 'http://localhost:3500';
    //axios.defaults.baseURL = 'https://ipadel-club-js-development.up.railway.app';
    console.log('Dev -> Base URL :', axios.defaults.baseURL)
  } */
  return (
    <Box
      display="flex"
      flexDirection="column"
      alignItems="center"
      justifyContent="center"
      minHeight="100vh"
      minWidth={"100vw"}

    >
      <Routes />
    </Box>

  );
};

export default App;
