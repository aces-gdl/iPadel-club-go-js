import React, { useState } from 'react'
import MainCard from '../layout/MainCard'
import ClubSelect from '../components/ClubSelect';
import EventSelect from '../components/EventSelect';
import CategorySelect from '../components/CategorySelect';

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
    <MainCard title="Bienvenido a iPadel Club" minHeigth={'300px'}>
      <ClubSelect 
        name='ClubID'
        value={values.ClubID}
        label="Clubs"
        handleupdate={handleUpdate}
      />
      <EventSelect 
        name='EventID'
        value={values.EventID}
        label="Eventos"
        handleupdate={handleUpdate}
      />
      <CategorySelect 
        name='CategoryID'
        value={values.CategoryID}
        label="Categorías"
        handleupdate={handleUpdate}
      />
    </MainCard>
  )
}

export default Home