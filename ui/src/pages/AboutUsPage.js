import React from 'react';
import {
  Container,
  Typography,
  Box,
  Paper,
  Grid,
  useTheme,
  useMediaQuery
} from '@mui/material';
import { IconBuildingSkyscraper, IconCode, IconCalendarEvent } from '@tabler/icons-react';

const AboutUsPage = () => {
  const theme = useTheme();
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  return (
    <Container maxWidth="lg">
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom align="center">
          Quiénes Somos
        </Typography>
        <Paper elevation={3} sx={{ p: 4, mt: 4 }}>
          <Grid container spacing={4}>
            <Grid item xs={12}>
              <Typography variant="body1" paragraph>
                Somos una compañía con más de 30 años de experiencia en el mundo de la tecnología, 
                especializada en crear soluciones innovadoras a través de páginas web. Nuestra 
                trayectoria nos ha permitido evolucionar y adaptarnos constantemente a las nuevas 
                tendencias y necesidades del mercado digital.
              </Typography>
            </Grid>
            <Grid item xs={12} md={4}>
              <Box display="flex" flexDirection="column" alignItems="center">
                <IconBuildingSkyscraper size={isMobile ? 48 : 64} color={theme.palette.primary.main} />
                <Typography variant="h6" align="center" sx={{ mt: 2 }}>
                  30+ Años de Experiencia
                </Typography>
                <Typography variant="body2" align="center">
                  Tres décadas liderando en soluciones tecnológicas
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={4}>
              <Box display="flex" flexDirection="column" alignItems="center">
                <IconCode size={isMobile ? 48 : 64} color={theme.palette.primary.main} />
                <Typography variant="h6" align="center" sx={{ mt: 2 }}>
                  Soluciones Web Innovadoras
                </Typography>
                <Typography variant="body2" align="center">
                  Creamos páginas web a la medida de tus necesidades
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12} md={4}>
              <Box display="flex" flexDirection="column" alignItems="center">
                <IconCalendarEvent size={isMobile ? 48 : 64} color={theme.palette.primary.main} />
                <Typography variant="h6" align="center" sx={{ mt: 2 }}>
                  Eventos de Pádel en México
                </Typography>
                <Typography variant="body2" align="center">
                  Nuestra nueva pasión: administrar eventos de pádel
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={12}>
              <Typography variant="body1" paragraph>
                Hoy, estamos emocionados de anunciar nuestra incursión en un nuevo y apasionante 
                campo: la administración de eventos relacionados con el pádel en México. Esta nueva 
                aventura nos permite combinar nuestra experiencia en tecnología con nuestra pasión 
                por el deporte, creando soluciones únicas para la comunidad del pádel.
              </Typography>
              <Typography variant="body1">
                Nuestro objetivo es revolucionar la forma en que se organizan y gestionan los 
                eventos de pádel, aprovechando nuestra experiencia en desarrollo web para ofrecer 
                plataformas intuitivas y eficientes. Estamos comprometidos con el crecimiento 
                del pádel en México y esperamos contribuir significativamente a su desarrollo 
                y popularización.
              </Typography>
            </Grid>
          </Grid>
        </Paper>
      </Box>
    </Container>
  );
};

export default AboutUsPage;