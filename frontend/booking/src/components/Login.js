import { useState } from "react";
import './styles/AddNewForm.css';

export default function Login(){

  const[username, setUserName] = useState('');
  const[password, setPassword] = useState('');
  const[user1, setUser1] = useState()

  

  const handleSubmit = (e) =>{
e.preventDefault()
    fetch("http://localhost:8000/users/login/" + username + "/" +password ,{
      method:"get",
      headers:{
          'Content-Type': 'application/json',
          'Accept': 'application/json',
      },
      
    }).then(res => res.json()).then((result)=>
    {
      //result.addHeader("Access-Control-Allow-Origin")
      setUser1(result)

    console.log(result)

    }
    )
  };
 

  return(
    <body>
      <div className="wrapper">
        <form onSubmit={handleSubmit}>
          <h1> Booking Platform</h1>
          <fieldset>
                <label>
                    <p>User Name</p>
                    <input id="username" name="username" onChange={(e)=>setUserName(e.target.value)}/>
                </label>
            </fieldset>
            <fieldset>
                <label>
                    <p>Password</p>
                    <input id="password" type="password" name="password" onChange={(e)=>setPassword(e.target.value)}/>
                </label>
            </fieldset>
            <button type="submit">Login</button>
        </form>
      </div>
      <div className="wrapper">
          Create an account? <a >Sing Up</a>
      </div>
    </body>

  );

}