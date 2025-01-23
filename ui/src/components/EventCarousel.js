import React from 'react';
import { Paper, Button, Typography, Box } from '@mui/material';
import Carousel from 'react-material-ui-carousel'
const EventCarousel = ({ events, selectedEventId, onEventSelect }) => {
  return (
    <Carousel>
      {events.map((event) => (
        <Paper
          key={event.ID}
          elevation={3}
          sx={{
            height: 300,
            width:300,
            position: 'relative',
            backgroundRepeat: 'no-repeat',
            backgroundImage: `url('/v1/images/${event?.image_url}')`,
            backgroundSize: 'contain',
            backgroundPosition: 'center',
          }}
        >
          <Box
            sx={{
              position: 'absolute',
              bottom: 0,
              left: 0,
              right: 0,
              bgcolor: 'rgba(0, 0, 0, 0.5)',
              color: 'white',
              padding: 2,
            }}
          >
            <Typography variant="h6">{event.name}</Typography>
            <Typography variant="body2">{event.date}</Typography>
            <Button
              variant="contained"
              color={selectedEventId === event.ID ? "secondary" : "primary"}
              onClick={() => onEventSelect(event.ID)}
              sx={{ mt: 1 }}
            >
              {selectedEventId === event.ID ? "Seleccionado" : "Seleccionar Evento"}
            </Button>
          </Box>
        </Paper>
      ))}
    </Carousel>
  );
};

export default EventCarousel;