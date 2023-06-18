import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import Login from './components/Login';
import HostHomepage from './components/HostHomepage';
import GuestHomepage from './components/GuestHomepage';
import GuestProfile from './components/GuestProfile';

// const root = ReactDOM.createRoot(document.getElementById('root'));
// root.render(
//   <React.StrictMode>
//     <App />
//   </React.StrictMode>
// );

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
// reportWebVitals();
ReactDOM.render(
  <Router>
    <Routes>
        <Route path='/' element={<Login/>}/>
        <Route path='/GuestHomepage' element = {<GuestHomepage/>}/>
        <Route path='/HostHomepage' element = {<HostHomepage/>}/>
        <Route path='/GuestProfile' element = {<GuestProfile/>}/>


        
    </Routes>
  </Router>,

 document.getElementById('root')
);