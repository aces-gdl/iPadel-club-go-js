/* eslint-disable  no-unused-vars */
import React, { useContext, useState } from 'react';
import { Button, Typography, Box, AppBar, Toolbar, IconButton, Menu, MenuItem, Snackbar, Alert } from '@mui/material';
import { Outlet, useNavigate } from 'react-router-dom';
import MenuIcon from '@mui/icons-material/Menu';
import { IconMoon, IconSun } from '@tabler/icons-react';
import { ThemeContext } from '../themes/ThemeContext';
import axios from 'axios';


const MainLayout = () => {
  const navigate = useNavigate();
  const [anchorEl, setAnchorEl] = React.useState(null);
  const { mode, toggleTheme } = useContext(ThemeContext);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });



  const handleLogout = async () => {
    try {
      const response = await axios.post('/v1/auth/logout');
      localStorage.removeItem('user');
      navigate('/auth/login');
      setSnackbar({ open: true, message: 'Sesión cerrada con exito.', severity: 'Success' });
    } catch (err) {
      setSnackbar({ open: true, message: 'Error al cerrar sesión. Inténtalo de nuevo.', severity: 'error' });
    }
  };
  const handleMenu = (event) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const menuItems = [
    { label: 'Inicio', action: () => navigate('/') },
    { label: 'Ranking', action: () => navigate('/ranking') },
    { label: 'Eventos', action: () => navigate('/inscriptions') },
    { label: 'Nosotros', action: () => navigate('/aboutus') },
    { label: 'Mi Cuenta', action: () => navigate('/update-password') },
    { label: 'Imagenes', action: () => navigate('/images') },
    { label: 'Cerrar Sesión', action: handleLogout },

  ];



  return (
    <Box sx={{ display: 'flex', flexDirection: 'column', height: '100vh', width: '100vw' }}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>

            iPadel Club, Hola : {JSON.parse(localStorage.getItem('user'))?.name} 
          </Typography>
          <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
            {menuItems.map((item) => (
              <Button key={item.label} color="inherit" onClick={item.action}>
                {item.label}
              </Button>
            ))}
            <MenuItem onClick={toggleTheme}>{mode === 'light' ? <IconMoon /> : <IconSun />}</MenuItem>

          </Box>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ display: { xs: 'flex', md: 'none' } }}
            onClick={handleMenu}
          >
            <MenuIcon />
          </IconButton>
          <Menu
            id="menu-appbar"
            anchorEl={anchorEl}
            anchorOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            keepMounted
            transformOrigin={{
              vertical: 'top',
              horizontal: 'right',
            }}
            open={Boolean(anchorEl)}
            onClose={handleClose}
          >
            {menuItems.map((item) => (
              <MenuItem key={item.label} onClick={() => { item.action(); handleClose(); }}>
                {item.label}
              </MenuItem>
            ))}
            <MenuItem onClick={toggleTheme}><IconSun /></MenuItem>
          </Menu>
        </Toolbar>
      </AppBar>
      <Box
        sx={{
          paddingTop: '12rem', 
          paddingLeft: '2rem',
          paddingRight: '2rem',
         
          flexGrow: 1,
        /*   backgroundImage: `url(${process.env.PUBLIC_URL}/images/padel-background.avif)`,
         */  backgroundSize: 'cover',
          backgroundPosition: 'center',
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'center',
          overflow: 'auto',
        }}
      >
        <Outlet />
      </Box>
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
      >
        <Alert onClose={() => setSnackbar({ ...snackbar, open: false })} severity={snackbar.severity}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Box>
  );
};

export default MainLayout;
