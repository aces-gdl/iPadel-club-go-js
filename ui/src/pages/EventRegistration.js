import React, { useEffect, useState } from 'react';
import {
    Button,
    Typography,
    Paper,
    Snackbar,
    Alert,
    Grid2,
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

    const handleEventSelection = async (event) => {
        console.log('Event selected:', event);
        const selectedEvent = events.find((item) => event === item.ID);
        try {
            let response = await axios.post('/v1/inscriptions', { 
                EventId: selectedEvent.ID, 
                CategoryID: 1,  
                PersonID: JSON.parse(localStorage.getItem('user')).ID, 
            });
             console.log('Response:', response.data);
            setSnackbar({ open: true, message: 'El evento seleccionado es:' + selectedEvent.name, severity: 'success' });
        } catch (error) {

        }

    }

    return (
        <MainCard title="Eventos disponibles" >

            <Paper elevation={3} style={{ minWidth: '350px', minHeight: '620px', marginTop: 1 }}>
                <EventCarousel
                    events={events}
                    selectedEventId={1}
                    onEventSelect={(e) => handleEventSelection(e)}
                />

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