import React, { useEffect, useState } from 'react';
import {
    Button,
    Typography,
    Paper,
    Snackbar,
    Alert,
} from '@mui/material';
import MainCard from '../layout/MainCard';
import axios from 'axios';
import EventCarousel from 'components/EventCarousel';

const EventRegistration = () => {
    const [snackbar, setSnackbar] = useState({ open: false, message: '', severity: 'success' });

    const [events, setEvents] = useState([]);

    const getEvents = async () => {
        try {
            const response = await axios.get('/v1/catalogs/events');
            setEvents(response.data);
        } catch (error) {
            console.error(error);
        }
    }

    useEffect(() => {
        getEvents();
    }, []);

    const handleEventSelection = (event) => {
        console.log('Event selected:', event);
        const selectedEvent = events.find((item) => event === item.ID);

        setSnackbar({ open: true, message: 'El evento seleccionado es:'+ selectedEvent.name, severity:'success' });

    }

    return (
        <MainCard title="Inscripción a Evento" >
                <Paper elevation={3} style={{ padding: '2rem', marginTop: '2rem', minWidth:'400px' }}>
                    <Typography variant="h6" gutterBottom>
                        Completa tu inscripción
                    </Typography>
                    <EventCarousel
                        events={events}
                        selectedEventId={1}
                        onEventSelect={(e) =>handleEventSelection(e) }
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
        </MainCard>
    );
};

export default EventRegistration;