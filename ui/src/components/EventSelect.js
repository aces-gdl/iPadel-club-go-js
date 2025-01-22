import React, { useState, useEffect } from 'react';
import { FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import axios from 'axios';

const EventSelect = ({ name, value, label, handleupdate }) => {
  const [events, setEvents] = useState([]);

  useEffect(() => {
    const fetchEvents = async () => {
      try {
        const response = await axios.get('/v1/catalogs/events');
        setEvents(response.data);
      } catch (error) {
        console.error('Error fetching events:', error);
      }
    };

    fetchEvents();
  }, []);


  return (
    <FormControl fullWidth>
      <InputLabel id={`${name}-select-label`}>{label}</InputLabel>
      <Select
        labelId={`${name}-select-label`}
        id={`${name}-select`}
        value={value}
        label={label}
        onChange={handleupdate}
        name={name}
      >
        {events.map((event) => (
          <MenuItem key={event.ID} value={event.ID}>
            {event.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

export default EventSelect;