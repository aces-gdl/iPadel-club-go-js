import React, { useState, useEffect } from 'react';
import {  Typography, Card, CardMedia, CardContent, CircularProgress, Grid2 } from '@mui/material';
import axios from 'axios';
import MainCard from '../../layout/MainCard';

const ImageViewerPage = () => {
  const [images, setImages] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchImages();
  }, []);

  const fetchImages = async () => {
    try {
      const response = await axios.get('/v1/images/list');
      setImages(response.data.images);
      setLoading(false);
    } catch (error) {
      console.error('Error fetching images:', error);
      setLoading(false);
    }
  };

  return (
    <MainCard title="Imágenes Subidas">
      <Grid2 container spacing={2}>
        <Grid2 item xs={12}>
          <Typography variant="h6" gutterBottom>
            Galería de Imágenes
          </Typography>
        </Grid2>
        {loading ? (
          <Grid2 item xs={12} style={{ display: 'flex', justifyContent: 'center' }}>
            <CircularProgress />
          </Grid2>
        ) : images.length > 0 ? (
          images.map((image, index) => (
            <Grid2 item xs={12} sm={6} md={4} lg={3} key={index}>
              <Card>
                <CardMedia
                  component="img"
                  height="140"
                  image={`/v1/images/${image}`}
                  alt={`Image ${index + 1}`}
                />
                <CardContent>
                  <Typography variant="body2" color="text.secondary">
                    {image}
                  </Typography>
                </CardContent>
              </Card>
            </Grid2>
          ))
        ) : (
          <Grid2 item xs={12}>
            <Typography variant="body1">No hay imágenes para mostrar.</Typography>
          </Grid2>
        )}
      </Grid2>
    </MainCard>
  );
};

export default ImageViewerPage;