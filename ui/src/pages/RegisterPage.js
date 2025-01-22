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
    useTheme,
    useMediaQuery,
} from '@mui/material';

import { DatePicker, LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';

import { useNavigate } from 'react-router-dom';
import {
    IconEye,
    IconEyeOff,
    IconPhone,
    IconLock,
    IconUser,
    IconUsers,
    IconCalendar,
    IconShirt,
} from '@tabler/icons-react';
import axios from 'axios';

const validationSchema = Yup.object({
    phone: Yup.string()
        .matches(/^[0-9]+$/, "Debe ser solo números")
        .min(9, "Debe tener al menos 9 dígitos")
        .max(15, "No debe exceder 15 dígitos")
        .required("El teléfono es requerido"),
    password: Yup.string()
        .min(8, "La contraseña debe tener al menos 8 caracteres")
        .required("La contraseña es requerida"),
    name: Yup.string().required("El nombre es requerido"),
    lastName: Yup.string().required("Los apellidos son requeridos"),
    birthDate: Yup.date().required("La fecha de nacimiento es requerida"),
    clothingSize: Yup.string().required("La talla de ropa es requerida"),
});

const RegisterPage = () => {
    const [showPassword, setShowPassword] = useState(false);
    const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });
    const navigate = useNavigate();
    const theme = useTheme();
    const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

    const formik = useFormik({
        initialValues: {
            phone: '',
            password: '',
            name: '',
            lastName: '',
            birthDate: null,
            clothingSize: '',
        },
        validationSchema: validationSchema,
        onSubmit: async (values) => {
            try {
                await axios.post('/v1/auth/register', values);
                setSnackbar({ open: true, message: 'Registro exitoso. Inicia sesión para continuar.', severity: 'success' });
                setTimeout(() => navigate('/auth/login'), 1500);
            } catch (err) {
                setSnackbar({ open: true, message: 'Error en el registro. Inténtalo de nuevo.', severity: 'error' });
            }
        },
    });

    return (
        <LocalizationProvider dateAdapter={AdapterDayjs}>
            <Container maxWidth="sm">
                <Paper elevation={3} style={{ padding: isMobile ? '1rem' : '2rem', marginTop: isMobile ? '1rem' : '2rem' }}>
                    <Box display="flex" flexDirection="column" alignItems="center">
                        <Typography variant={isMobile ? "h5" : "h4"} gutterBottom>iPadel Club</Typography>
                        <Typography variant={isMobile ? "h6" : "h5"} gutterBottom>Registro</Typography>
                        <form onSubmit={formik.handleSubmit} style={{ width: '100%' }}>
                            {/* Phone TextField */}
                            <TextField
                                fullWidth
                                id="phone"
                                name="phone"
                                label="Teléfono"
                                value={formik.values.phone}
                                onChange={formik.handleChange}
                                error={formik.touched.phone && Boolean(formik.errors.phone)}
                                helperText={formik.touched.phone && formik.errors.phone}
                                margin="normal"
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <IconPhone size={isMobile ? 16 : 20} />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            {/* Password TextField */}
                            <TextField
                                fullWidth
                                id="password"
                                name="password"
                                label="Contraseña"
                                type={showPassword ? 'text' : 'password'}
                                value={formik.values.password}
                                onChange={formik.handleChange}
                                error={formik.touched.password && Boolean(formik.errors.password)}
                                helperText={formik.touched.password && formik.errors.password}
                                margin="normal"
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <IconLock size={isMobile ? 16 : 20} />
                                        </InputAdornment>
                                    ),
                                    endAdornment: (
                                        <InputAdornment position="end">
                                            <IconButton onClick={() => setShowPassword(!showPassword)}>
                                                {showPassword ? <IconEyeOff size={isMobile ? 16 : 20} /> : <IconEye size={isMobile ? 16 : 20} />}
                                            </IconButton>
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            {/* Name TextField */}
                            <TextField
                                fullWidth
                                id="name"
                                name="name"
                                label="Nombre"
                                value={formik.values.name}
                                onChange={formik.handleChange}
                                error={formik.touched.name && Boolean(formik.errors.name)}
                                helperText={formik.touched.name && formik.errors.name}
                                margin="normal"
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <IconUser size={isMobile ? 16 : 20} />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            {/* Last Name TextField */}
                            <TextField
                                fullWidth
                                id="lastName"
                                name="lastName"
                                label="Apellidos"
                                value={formik.values.lastName}
                                onChange={formik.handleChange}
                                error={formik.touched.lastName && Boolean(formik.errors.lastName)}
                                helperText={formik.touched.lastName && formik.errors.lastName}
                                margin="normal"
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <IconUsers size={isMobile ? 16 : 20} />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            {/* Birth Date Picker */}
                            <DatePicker
                                label="Fecha de Nacimiento"
                                value={formik.values.birthDate}
                                onChange={(value) => formik.setFieldValue('birthDate', value)}
                                renderInput={(params) => (
                                    <TextField
                                        {...params}
                                        fullWidth
                                        margin="normal"
                                        error={formik.touched.birthDate && Boolean(formik.errors.birthDate)}
                                        helperText={formik.touched.birthDate && formik.errors.birthDate}
                                        InputProps={{
                                            startAdornment: (
                                                <InputAdornment position="start">
                                                    <IconCalendar size={isMobile ? 16 : 20} />
                                                </InputAdornment>
                                            ),
                                        }}
                                    />
                                )}
                            />
                            {/* Clothing Size TextField */}
                            <TextField
                                fullWidth
                                id="clothingSize"
                                name="clothingSize"
                                label="Talla de Ropa"
                                value={formik.values.clothingSize}
                                onChange={formik.handleChange}
                                error={formik.touched.clothingSize && Boolean(formik.errors.clothingSize)}
                                helperText={formik.touched.clothingSize && formik.errors.clothingSize}
                                margin="normal"
                                InputProps={{
                                    startAdornment: (
                                        <InputAdornment position="start">
                                            <IconShirt size={isMobile ? 16 : 20} />
                                        </InputAdornment>
                                    ),
                                }}
                            />
                            <Button 
                                color="primary" 
                                variant="contained" 
                                fullWidth 
                                type="submit" 
                                style={{ marginTop: '1rem' }}
                                size={isMobile ? "small" : "medium"}
                            >
                                Registrarse
                            </Button>
                        </form>
                        {/* Login Link */}
                        <Button 
                            color="secondary" 
                            style={{ marginTop: '0.5rem' }} 
                            onClick={() => navigate('/auth/login')}
                            size={isMobile ? "small" : "medium"}
                        >
                            ¿Ya tienes cuenta? Inicia sesión
                        </Button>
                    </Box>
                </Paper>
                {/* Snackbar for notifications */}
                <Snackbar
                    open={snackbar.open}
                    autoHideDuration={3000}
                    onClose={() => setSnackbar({ ...snackbar, open: false })}
                    anchorOrigin={{ vertical: 'bottom', horizontal: 'center' }}
                >
                    <Alert onClose={() => setSnackbar({ ...snackbar, open: false })} severity={snackbar.severity}>
                        {snackbar.message}
                    </Alert>
                </Snackbar>
            </Container>
        </LocalizationProvider>
    );
}

export default RegisterPage;
