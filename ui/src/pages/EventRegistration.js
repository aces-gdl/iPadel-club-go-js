import React, { useState } from 'react';
import {
    TextField,
    Button,
    Box,
    Typography,
    Container,
    Paper,
    Snackbar,
    Alert
} from '@mui/material';
import { useFormik } from 'formik';
import * as Yup from 'yup';
import MainCard from '../layout/MainCard';
import EventSelect from '../components/EventSelect';
import CategorySelect from '../components/CategorySelect';
import axios from 'axios';

const validationSchema = Yup.object({
    EventID: Yup.string().required('El evento es requerido'),
    CategoryID: Yup.string().required('La categoría es requerida'),
    partnerPhone: Yup.string().required('El teléfono del compañero es requerido').min(10, 'El teléfono debe tener al menos 10 dígitos'),
});
const EventRegistration = () => {
    const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });
    const [person2_id, setPerson2_id] = useState(0);

    const formik = useFormik({
        initialValues: {
            EventID: '',
            CategoryID: '',
            partnerPhone: '',
            partnerName: '',
        },
        validationSchema: validationSchema,
        onSubmit: async (values) => {
            try {
                let payload = {
                    event_id: values.EventID,
                    category_id: values.CategoryID,
                    person1_id: JSON.parse(localStorage.getItem('user')).ID,
                    person2_id: person2_id,
                }
                await axios.post('/v1/inscription', payload);
                setSnackbar({ open: true, message: 'Inscripción exitosa', severity: 'success' });
                formik.resetForm();
            } catch (error) {
                setSnackbar({ open: true, message: 'Error en la inscripción', severity: 'error' });
            }
        },
    });

    const handlePartnerPhoneChange = async (e) => {
        const { value } = e.target;
        formik.setFieldValue('partnerPhone', value);

        if (value.length >= 10) {
            try {
                const response = await axios.get(`/v1/person/is-available?EventID=${formik.values.EventID}&Phone=${value}&CategoryID=${formik.values.CategoryID}`);    
                if (response.data && response.data.data.name) {
                    formik.setFieldValue('partnerName', `${response.data.data.name} ${response.data.data.lastName}`);
                    setPerson2_id(response.data.data.ID);
                } else {
                    formik.setFieldValue('partnerName', '');
                }
            } catch (error) {
                console.error('Error buscando usuario:', error.data.error);
                setSnackbar({ open: true, message: 'Error buscando usuario', severity: 'error' });
            }
        }
    };

    return (
        <MainCard title="Inscripción a Evento" minHeight={'300px'}>
            <Container maxWidth="sm">
                <Paper elevation={3} style={{ padding: '2rem', marginTop: '2rem' }}>
                    <Typography variant="h6" gutterBottom>
                        Completa tu inscripción
                    </Typography>
                    <form onSubmit={formik.handleSubmit}>
                        <EventSelect
                            name='EventID'
                            value={formik.values.EventID}
                            label="Evento"
                            handleupdate={formik.handleChange}
                            error={formik.touched.EventID && Boolean(formik.errors.EventID)}
                            helperText={formik.touched.EventID && formik.errors.EventID}
                        />
                        <Box mt={2}>
                            <CategorySelect
                                name='CategoryID'
                                value={formik.values.CategoryID}
                                label="Categoría"
                                handleupdate={formik.handleChange}
                                error={formik.touched.CategoryID && Boolean(formik.errors.CategoryID)}
                                helperText={formik.touched.CategoryID && formik.errors.CategoryID}
                            />
                        </Box>
                        <TextField
                            fullWidth
                            margin="normal"
                            name="partnerPhone"
                            label="Teléfono del compañero"
                            value={formik.values.partnerPhone}
                            onChange={handlePartnerPhoneChange}
                            error={formik.touched.partnerPhone && Boolean(formik.errors.partnerPhone)}
                            helperText={formik.touched.partnerPhone && formik.errors.partnerPhone}
                        />
                        <TextField
                            fullWidth
                            margin="normal"
                            name="partnerName"
                            label="Nombre del compañero"
                            value={formik.values.partnerName}
                            InputProps={{
                                readOnly: true,
                            }}
                        />
                        <Button
                            type="submit"
                            variant="contained"
                            color="primary"
                            fullWidth
                            style={{ marginTop: '1rem' }}
                        >
                            Inscribirse
                        </Button>
                    </form>
                </Paper>
            </Container>
            <Snackbar
                open={snackbar.open}
                autoHideDuration={6000}
                onClose={() => setSnackbar({ ...snackbar, open: false })}
            >
                <Alert onClose={() => setSnackbar({ ...snackbar, open: false })} severity={snackbar.severity}>
                    {snackbar.message}
                </Alert>
            </Snackbar>
        </MainCard>
    );
};

export default EventRegistration;