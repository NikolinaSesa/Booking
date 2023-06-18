import React, { useEffect, useState } from 'react'
import './styles/AddNewForm.css'
import { useNavigate } from 'react-router-dom'



const GuestHomepage = () => {
  const [hosts, setHosts] = useState([])






  useEffect(()=>{

    fetch("http://localhost:8000/users/getAll",{
      headers : { 
        'Content-Type': 'application/json',
        'Accept': 'application/json',
        "method":"GET",
       },
    })
    .then(res =>res.json())
    .then((result)=>
    {
        setHosts(result);
    }
    )
  }, [])

  const navigate = useNavigate();

const deleteOrder = (e) =>{
  var Id = JSON.parse(localStorage.getItem('Id'))
    var test = JSON.parse(localStorage.getItem('testToken'))
    fetch("http://localhost:8081/api/orders/delete/"+Id,{
        method:"DELETE",
        headers : { 
          'Content-Type': 'application/json',
           Authorization: `Bearer ${test.accessToken}`,
         },
        body:JSON.stringify(test.accessToken)
    
      }).then(() =>{
        window.location.href = "/Orders"
      })
   
}


  return(
   

    
        <body>
            <div class="topnav">
                <a href="/UserHomepage">Home Page</a>
                <a class="active" href="/Orders">Your orders</a>
                <a  href="/ContractDeals">Contracts</a>
                <a href="/UserProfileUpdate">Profile</a>
             
            </div>
           
            <div className='wrapper'>
                <table>
                    <tr>
                        <th>Order id</th>
                        <th>Delivery address</th>
                        <th>City</th>
                        <th>Delivery date</th>
                        <th>Total price of order</th>
                        <th>Order status</th>
                    </tr>
                    {hosts.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.users.id}</td>
                              
                                <td>
                                    <button onClick={(e) => {
                                        e.preventDefault()

                                        localStorage.setItem('Id', val.id)

                                        navigate(`/SeeOrder/`+val.id);
                                        
                                    }}>View</button>
                                </td>
                                <td>
                                    <button onClick={(e) => {
                                        

                                        localStorage.setItem('Id', val.id)

                                        navigate(`/UpdateOrder/`+val.id);
                                        
                                    }}>Update</button>
                                </td>
                                <td>
                                    <button onClick={(e) =>{
                                         localStorage.setItem('Id', val.id);
                                        deleteOrder();
                                       
                                    }}
                                    >Cancel</button>
                                </td>
                            </tr>
                        )
                    })}
                </table>
                
            </div>
            
            <button className='request' onClick={(e) =>{
                                     navigate(`/CreatePdf`);  
                                       
                                    }}
                                    >Generate report</button>
        </body>
    
  )
 
}

export default GuestHomepage
