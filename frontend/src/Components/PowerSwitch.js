import React, { useState } from 'react';

import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';

const PowerSwitch = () => {
  return (
    <div className='control'>
     <Stack direction="row" spacing={3}>
      <Button variant="contained" color="success">
        POWER ON
      </Button>
      <Button variant="contained" color="error">
        POWER OFF
      </Button>
    </Stack>
    </div>
  );
}
export default PowerSwitch;