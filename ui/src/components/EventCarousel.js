import React from 'react';
import { Paper, Button, Typography, Box } from '@mui/material';
import Carousel from 'react-material-ui-carousel'
import dayjs from 'dayjs';
const EventCarousel = ({ events, selectedEventId, onEventSelect }) => {
  let myDate = dayjs(events[0]?.start_date).format('dddd, MMMM D, YYYY');;
  return (
    <Carousel
      autoPlay={false}
      navButtonsAlwaysVisible={true}
      height={600}
      
    >
      {events.map((event) => (
        <Box paddingTop={1}>
          <Box
            key={event.ID}
            elevation={3}
            sx={{
              minHeight: 430,
              minWidth: '100%',
              backgroundRepeat: 'no-repeat',
              backgroundImage: `url('/v1/images/${event?.image_url}')`,
              backgroundSize: 'contain',
              backgroundPosition: 'center',
            }}
          >

          </Box>
          <Box
            paddingTop={5}
            marginX={2}
          >
            <Typography variant="h6" fontWeight={800}>{event.name}</Typography>
            <Box display={'flex'} flexDirection={'row'}>
              <Typography variant='body2' fontWeight={800} component={'p'}>Club: </Typography>
              <Typography variant="body2" > {event.club.name}</Typography>
            </Box>       
            <Box display={'flex'} flexDirection={'row'}>
            <Typography variant='body2' fontWeight={800} component={'p'}>Del: </Typography>
              <Typography variant="body2" > {dayjs(event.start_date).format("DD/MM")}</Typography>
              <Typography variant='body2' fontWeight={800} component={'p'}> al: </Typography>
              <Typography variant="body2" > {dayjs(event.end_date).format("DD/MM")}</Typography>
            </Box>
            <Button
              fullWidth
              variant="contained"
              color={"primary"}
              onClick={() => onEventSelect(event.ID)}
              sx={{ mt: 1 }}
            >
              Inscribirme
            </Button>
          </Box>
        </Box>
      ))}
    </Carousel>
  );
};

export default EventCarousel;