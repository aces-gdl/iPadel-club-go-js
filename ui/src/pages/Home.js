import React, { useState } from 'react'
import MainCard from '../layout/MainCard'
import ClubSelect from '../components/ClubSelect';
import EventSelect from '../components/EventSelect';
import CategorySelect from '../components/CategorySelect';
import { Container } from '@mui/material';

const Home = () => {
  const [values, setValues] = useState({
    ClubID: '1',
    EventID: '1',
    CategoryID: '1',
  })
  const handleUpdate = (e) => {
    const { name, value } = e.target;
    setValues({ ...values, [name]: value });
  };

  return (
    <MainCard title="Bienvenido a iPadel Club" >
      <Container fixed>
        aqui
      </Container>    
    </MainCard>
  )
}

export default Home