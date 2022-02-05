import React from 'react';
import '../App.css';

import { Button, Form, FormGroup, Label, Input }
  from 'reactstrap';
  



function SellerUploadPage() {
  return (
    
    
    <Form className="details">
      <h1 className="text-center">Seller Page (Picture Details)</h1>
      <FormGroup>
        <Label>Picture Title</Label>
        <Input type="text" placeholder='Picture Title'/><br></br><br></br>
        <Label>Picture Description</Label>
        <Input type="text" placeholder='Picture Description'/><br></br><br></br>
        <Label>Price</Label>
        <Input type="number" placeholder='Enter the price'/><br></br><br></br>
        <Label>Upload Image</Label>
        <Input type="file" placeholder='Upload Image'/><br></br><br></br>
        <label for="genre">Choose a genre:</label>

        <select name="genre" id="genre">
        <option value="Genre1">Genre1</option>
        <option value="Genre2">Genre2</option>
        <option value="Genre3">Genre3</option>
        <option value="Genre4">Genre4</option>
        </select>
        
        <Input type="submit" placeholder='Submit'/><br></br><br></br>
      </FormGroup>
    </Form>
  );
}

export default SellerUploadPage;
