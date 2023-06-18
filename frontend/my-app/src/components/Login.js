import { useState } from "react";
import './styles/AddNewForm.css';

export default function Login(){

  const[username, setUserName] = useState('');
  const[password, setPassword] = useState('');
  const[token, setToken] = useState([])

  const handleSubmit = (e) =>{
    e.preventDefault()
    let user = {username, password}
    console.log(user)
    fetch("http://localhost:8081/auth/login", {
      headers : { 
        'Content-Type': 'application/json',
        'Accept': 'application/json',
      },
      method:"POST",
      body:JSON.stringify(user),
    }).then(res => res.json()).then((result)=>
    {
      setToken(result);

      let testToken = {accessToken : "", expiresIn : 0}

      testToken.accessToken = result.accessToken;
      testToken.expiresIn = result.expiresIn;

      localStorage.setItem('testToken', JSON.stringify(testToken));
      localStorage.setItem('user_userName', username);

      console.log(JSON.parse(localStorage.getItem('testToken')));
      if(localStorage.getItem('user_userName')==="marko123"){
        window.location.href='/Orders';
      }else

      window.location.href='/DeliverManagerOrders';
    }
    )
  };
 

  return(
    <body>
      <div className="wrapper">
        <form >
          <h1>Chocolate Factory</h1>
          <fieldset>
                <label>
                    <p>User Name</p>
                    <input id="userName" name="userName" onChange={(e)=>setUserName(e.target.value)}/>
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
          Create an account? <a href="/DeliveryManagerHome">Sing Up</a>
      </div>
    </body>

  );

}