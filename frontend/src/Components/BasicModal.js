
import React, { useState, useEffect } from "react";
import axios from "axios";
import Box from '@mui/material/Box';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Modal from '@mui/material/Modal';

const style = {
  position: 'absolute',
  top: '50%',
  left: '50%',
  transform: 'translate(-50%, -50%)',
  width: 400,
  bgcolor: 'background.paper',
  border: '2px solid #000',
  boxShadow: 24,
  p: 4,
};
const cols =  8;
const rows =  2;
function srcset(image, width, height, rows = 1, cols = 1) {
  return {
    src: `${image}?w=${width * cols}&h=${height * rows}&fit=crop&auto=format`,
    srcSet: `${image}?w=${width * cols}&h=${
      height * rows
    }&fit=crop&auto=format&dpr=2 2x`,
  };
}


export default function BasicModal(props) {
  
  const [open, setOpen] = useState(false);
  const handleOpen = () => setOpen(true);
  const handleClose = () =>  {console.log("here") 
  setOpen(false);}
  const [fetchImages, setFetchedImageData] = useState([]);
  useEffect(() => {
    axios.get(`http://localhost:8085/fetchProductInfo/${props.imageId}`).then((response) => {
      
      setFetchedImageData(response.data.data);
      // console.log(fetchImages);
    });
  }, []);

  return (
    <div>
      <Button onClick={handleOpen}>Open modal</Button>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box sx={style}>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            {fetchImages.title}
          </Typography>
          <img
                            {...srcset(fetchImages.wImageUrl, 125, 100, rows, cols)}
                            alt={fetchImages.title}
                            loading="lazy"
                          />
            <Typography id="modal-modal-title" variant="h6" component="h2">
            PRICE = {fetchImages.price}
          </Typography>
          <Typography id="modal-modal-title" variant="h6" component="h2">
           Description =  {fetchImages.description}
          </Typography>
          <Typography id="modal-modal-title" variant="h6" component="h2">
            Upload Date = {fetchImages.uploadedAt}
          </Typography>
          <button onClick={handleClose}>Close</button>
        </Box>
      </Modal>
      {console.log(open)}
    </div>
  );
}
