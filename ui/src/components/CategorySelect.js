import React, { useState, useEffect } from 'react';
import { FormControl, InputLabel, Select, MenuItem } from '@mui/material';
import axios from 'axios';

const CategorySelect = ({ name, value, label, handleupdate }) => {
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        const response = await axios.get('/v1/catalogs/categories');
        setCategories(response.data);
      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    };

    fetchCategories();
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
        {categories.map((category) => (
          <MenuItem key={category.ID} value={category.ID}>
            {category.description}
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
};

export default CategorySelect;