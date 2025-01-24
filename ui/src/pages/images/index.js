import { Grid2 } from '@mui/material'
import React from 'react'
import ImageUploadPage from './ImageUploadPage'
import ImageViewerPage from './ImageViewerPage'
import ImageRemovePage from './ImageRemovePage'

const index = () => {
  return (
    <Grid2 container spacing={2} >
        <Grid2 item size={{ xs: 12, md: 6 }} >
          <ImageUploadPage/>
        </Grid2>
        <Grid2 item size={{ xs: 12, md: 6 }}>
          <ImageViewerPage/>
        </Grid2>
        <Grid2 item size={{ xs: 12, md: 6 }}>
          <ImageRemovePage/>
        </Grid2>
      </Grid2>
  )
}

export default index