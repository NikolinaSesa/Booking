import React, { useEffect, useState } from 'react'
import './styles/AddNewForm.css'
import 'bootstrap/dist/css/bootstrap.min.css'
import { useNavigate } from 'react-router-dom'



const GuestHomepage = () => {
  const [orders, setOrders] = useState([])







  const navigate = useNavigate();
  
  const navigateToAddNew = (e) =>{
    e.preventDefault()
    window.location.href = "/CreateNewOrder"
}

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


const rowEvent = {
    onClick: (e, row) => {
     
        navigate(`/SeeOrder/`+row.id);
    },
  }; 
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
                        <th>First name</th>
                        <th>Last name</th>
                        <th>Status</th>
                        <th>Average grade</th>
                        
                    </tr>
                    {orders.map((val, key) => {
                        return(
                            <tr key={key} >
                                <td>{val.id}</td>
                                <td>{val.address}</td>
                                <td>{val.city}</td>
                                <td>{val.date}</td>
                                <td>{val.totalPrice}</td>
                                <td>{val.status}</td>
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
