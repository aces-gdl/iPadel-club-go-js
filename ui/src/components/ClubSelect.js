import React, { useState, useEffect } from 'react';
import { FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import axios from 'axios';

const ClubSelect = (props) => {
  const [clubs, setClubs] = useState([]);

  useEffect(() => {
    const fetchClubs = async () => {
      try {
        const response = await axios.get('/v1/catalogs/clubs');
        setClubs(response.data);
      } catch (error) {
        console.error('Error fetching clubs:', error);
      }
    };

    fetchClubs();
  }, []);



  return (
    <FormControl fullWidth>
      <InputLabel id="club-select-label">Club</InputLabel>
      <Select
       labelId="CategoryL"
       id={props.name}
       name={props.name}
       value={props.value}
       label="Club"
       onChange={props.handleupdate}>
        {clubs.map((club) => (
          <MenuItem key={club.ID} value={club.ID}>
            {club.name}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

export default ClubSelect;