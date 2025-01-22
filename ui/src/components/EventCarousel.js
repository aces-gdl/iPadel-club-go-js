import React from 'react';
import Carousel from 'react-material-ui-carousel';
import { Paper, Button, Typography, Box } from '@mui/material';

const EventCarousel = ({ events, selectedEventId, onEventSelect }) => {
  return (
    <Carousel>
      {events.map((event) => (
        <Paper
          key={event.id}
          elevation={3}
          sx={{
            height: 300,
            position: 'relative',
            backgroundImage: `url(${event.imageUrl})`,
            backgroundSize: 'cover',
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
              color={selectedEventId === event.id ? "secondary" : "primary"}
              onClick={() => onEventSelect(event.id)}
              sx={{ mt: 1 }}
            >
              {selectedEventId === event.id ? "Selected" : "Select Event"}
            </Button>
          </Box>
        </Paper>
      ))}
    </Carousel>
  );
};

export default EventCarousel;