import React , { useState , useEffect}  from 'react';
import Axios from 'axios';
export default function UserLandingPage()
{
    const [fetchImages , setFetchedImageData] = useState([])
    const [delay , setDelay] = useState("")
    const obj = Object.entries(fetchImages)
    const obj1 = Object.entries(obj[0][1])
     useEffect(()=> {
         
        Axios.get("http://localhost:8888/fetchImages").then((response)=>{
            
          setFetchedImageData([response.data])
        //   setDelay("Check")
        //   console.log(response.data);
        }); 
      },[]);
    return(
        <div>
           {/* { {fetchImages.data.map((val)=> {
                return <h1> ImageName: {val.data}  </h1> 
            })}  } */}
            {/* {setDelay("again")} */}
            <h1> hshhshd</h1>
            {console.log(obj1[0][1])}

           {/* {fetchImages.map((c)=> {
               return <h5>{Object.values(c)}</h5>
           })} */}

            {
                obj1[0][1].map((val)=> {
                    return <img src={val.wImageUrl}/>
                    

                    
                }
                )
            }
                
            
            
        </div> 
        
            
    );
}