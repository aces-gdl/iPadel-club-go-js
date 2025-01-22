import React, { useState } from 'react';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import {
  TextField,
  Button,
  Box,
  Typography,
  Container,
  Paper,
  IconButton,
  InputAdornment,
  Snackbar,
  Alert,
} from '@mui/material';
import { IconEye, IconEyeOff, IconLock } from '@tabler/icons-react';
import axios from 'axios';

const UpdatePasswordPage = () => {
  const [showCurrentPassword, setShowCurrentPassword] = useState(false);
  const [showNewPassword, setShowNewPassword] = useState(false);
  const [showConfirmPassword, setShowConfirmPassword] = useState(false);
  const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });

  const formik = useFormik({
    initialValues: {
      currentPassword: '',
      newPassword: '',
      confirmPassword: '',
    },
    validationSchema: Yup.object({
      currentPassword: Yup.string().required('La contraseña actual es requerida'),
      newPassword: Yup.string()
        .min(8, 'La contraseña debe tener al menos 8 caracteres')
        .matches(/[a-zA-Z]/, 'La contraseña debe contener al menos una letra')
        .matches(/[0-9]/, 'La contraseña debe contener al menos un número')
        .required('La nueva contraseña es requerida'),
      confirmPassword: Yup.string()
        .oneOf([Yup.ref('newPassword'), null], 'Las contraseñas deben coincidir')
        .required('Confirma tu nueva contraseña'),
    }),
    onSubmit: async (values) => {
      try {
        await axios.post('/v1/security/update-password', {
          currentPassword: values.currentPassword,
          newPassword: values.newPassword,
        });
        setSnackbar({ open: true, message: 'Contraseña actualizada con éxito', severity: 'success' });
        formik.resetForm();
      } catch (error) {
        setSnackbar({ open: true, message: 'Error al actualizar la contraseña', severity: 'error' });
      }
    },
  });

  const handleTogglePasswordVisibility = (field) => {
    switch (field) {
      case 'current':
        setShowCurrentPassword(!showCurrentPassword);
        break;
      case 'new':
        setShowNewPassword(!showNewPassword);
        break;
      case 'confirm':
        setShowConfirmPassword(!showConfirmPassword);
        break;
      default:
        break;
    }
  };

  return (
    <Container component="main" maxWidth="xs">
      <Paper elevation={3} sx={{ mt: 8, p: 4 }}>
        <Typography component="h1" variant="h5" align="center" gutterBottom>
          Actualizar Contraseña
        </Typography>
        <Box component="form" onSubmit={formik.handleSubmit} noValidate sx={{ mt: 1 }}>
          <TextField
            margin="normal"
            required
            fullWidth
            name="currentPassword"
            label="Contraseña Actual"
            type={showCurrentPassword ? 'text' : 'password'}
            id="currentPassword"
            autoComplete="current-password"
            value={formik.values.currentPassword}
            onChange={formik.handleChange}
            error={formik.touched.currentPassword && Boolean(formik.errors.currentPassword)}
            helperText={formik.touched.currentPassword && formik.errors.currentPassword}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <IconLock size={20} />
                </InputAdornment>
              ),
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton onClick={() => handleTogglePasswordVisibility('current')}>
                    {showCurrentPassword ? <IconEyeOff size={20} /> : <IconEye size={20} />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="newPassword"
            label="Nueva Contraseña"
            type={showNewPassword ? 'text' : 'password'}
            id="newPassword"
            autoComplete="new-password"
            value={formik.values.newPassword}
            onChange={formik.handleChange}
            error={formik.touched.newPassword && Boolean(formik.errors.newPassword)}
            helperText={formik.touched.newPassword && formik.errors.newPassword}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <IconLock size={20} />
                </InputAdornment>
              ),
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton onClick={() => handleTogglePasswordVisibility('new')}>
                    {showNewPassword ? <IconEyeOff size={20} /> : <IconEye size={20} />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
          />
          <TextField
            margin="normal"
            required
            fullWidth
            name="confirmPassword"
            label="Confirmar Nueva Contraseña"
            type={showConfirmPassword ? 'text' : 'password'}
            id="confirmPassword"
            autoComplete="new-password"
            value={formik.values.confirmPassword}
            onChange={formik.handleChange}
            error={formik.touched.confirmPassword && Boolean(formik.errors.confirmPassword)}
            helperText={formik.touched.confirmPassword && formik.errors.confirmPassword}
            InputProps={{
              startAdornment: (
                <InputAdornment position="start">
                  <IconLock size={20} />
                </InputAdornment>
              ),
              endAdornment: (
                <InputAdornment position="end">
                  <IconButton onClick={() => handleTogglePasswordVisibility('confirm')}>
                    {showConfirmPassword ? <IconEyeOff size={20} /> : <IconEye size={20} />}
                  </IconButton>
                </InputAdornment>
              ),
            }}
          />
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Actualizar Contraseña
          </Button>
        </Box>
      </Paper>
      <Snackbar
        open={snackbar.open}
        autoHideDuration={6000}
        onClose={() => setSnackbar({ ...snackbar, open: false })}
        anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
      >
        <Alert onClose={() => setSnackbar({ ...snackbar, open: false })} severity={snackbar.severity} sx={{ width: '100%' }}>
          {snackbar.message}
        </Alert>
      </Snackbar>
    </Container>
  );
};

export default UpdatePasswordPage;