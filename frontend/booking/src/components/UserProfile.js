import React, { useEffect, useState } from 'react'
import './styles/AddNewForm.css';
import { useParams } from 'react-router-dom';



function UserProfile(){
    let { id } = useParams();
    var [order, setOrder] = useState('');
    var [address, setAddress] = useState('');
    var [city, setCity] = useState('');
    var [date, setDate] = useState('');
    var [userId, setUserId] = useState('');
    var[quantity,setQuantity]=useState('');
    var [productId, setProductId] = useState('');



   

    useEffect(() =>{
        var test = JSON.parse(localStorage.getItem('testToken'))
        fetch("http://localhost:8081/api/orders/" + id, {
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
            setOrder(result);
            console.log(result)
            
        }
        )
      },[])
    
      
    const handleClick = (e) =>{
        var test = JSON.parse(localStorage.getItem('testToken'))
        e.preventDefault()
        var admin = order;
       
        admin.city = city;
        admin.date = date;
        admin.address = address;
        
        
    var orderedProducts = [

        {
          productId: productId,
          quantity: quantity,
        }
      ]
    var orderedProductsDTO = orderedProducts;
    const new_order = {address, city, date,orderedProductsDTO}
        console.log(admin);
        fetch("http://localhost:8081/api/orders/update/" + id,{ 
        method:"PUT",
        headers:{
            "Content-Type":"application/json",
            'Accept': 'application/json',
            Authorization: `Bearer ${test.accessToken}`,
        },
        body:JSON.stringify(new_order)
    
      }).then(() =>{
        console.log("Order updated")
        console.log(admin);
        window.location.href = "/Orders"
      })
    }
      
    
    

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
                    placeholder={order.address}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Last name</p>
                    <input onChange={(e)=>setCity(e.target.value)}
                    placeholder={order.city}/>
                </label>
            </fieldset>
            <fieldset>
               <label>
                    <p>Date</p>
                    <input type="date"  id="date" name="date" onChange={(e)=>setDate(e.target.value)}
                    placeholder={order.date}/>
                </label>
            </fieldset>
            
            <button onClick={handleClick} className='update' type="submit">Submit</button>
            </form>
            
        </div>
        </body>
    )
    }

export default UserProfile;