import React, { useEffect, useState } from 'react'
import './styles/AddNewForm.css';
import { useParams } from 'react-router-dom';



function UserProfile(){
    let { id } = useParams();
    var [host, setHost] = useState('');
    var [address, setAddress] = useState('');
    var [city, setCity] = useState('');
    var [date, setDate] = useState('');
    var [userId, setUserId] = useState('');
    var[quantity,setQuantity]=useState('');
    var [productId, setProductId] = useState('');



   

    useEffect(() =>{
        var test = JSON.parse(localStorage.getItem('testToken'))
        fetch("http://localhost:8000/users/user/" + id, {
          method:"GET",
          headers:{
            "Content-Type":"application/json",
            'Accept': 'application/json',
            Authorization: `Bearer ${test.accessToken}`
                }
    
        })
        .then(res => res.json())
        .then((result) =>
        {
            setHost(result);
            
        }
        )
      },[])
    
      

    return(
        <body>
            
        <div class="topnav">
        <a href="/UserHomepage">Home Page</a>
                <a class="active" href="/Orders">Your orders</a>
                <a >Contracts</a>
                <a href="/UserProfileUpdate">Profile</a>
               
        </div>
        <div className="wrapper">
          <form className='1' >
          <h1>User profile</h1>
        
            <fieldset>
                
            </fieldset>

            <fieldset>
                <label>
                    <p>First name</p>
                    <input  onChange={(e)=>setAddress(e.target.value)}
                    placeholder={host.user.firstName}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Last name</p>
                    <input onChange={(e)=>setCity(e.target.value)}
                    placeholder={host.user.lastName}/>
                </label>
            </fieldset>
           
            
            <button  className='update' type="submit">Submit</button>
            </form>
            
        </div>
        </body>
    )
    }

export default UserProfile;