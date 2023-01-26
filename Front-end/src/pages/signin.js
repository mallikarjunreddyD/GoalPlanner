import React from "react";
import { useState } from 'react';
import {Link, Routes, Route, useNavigate} from 'react-router-dom';


const SignIn = () => {
  const navigate = useNavigate();
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState(false);
  
  
  const handleSubmit = (e) => {
    e.preventDefault();
    if (email === '' || password === '') {
      setError(true);
    } else {
      setSubmitted(true);
   navigate('/signinSuccess')

      setError(false);
    }
  };

    return (
        <div className="form">
          <div>
            <h1>User Sign In</h1>
          </div>
     
          <form>
            {/* Labels and inputs for form data */}
     
            <label className="label">Email</label>
            <input className="input"
              value={email} type="email"  onChange={(e)=> {
                setEmail(e.target.value)
            }} />
     
            <label className="label">Password</label>
            <input className="input"
              value={password} type="password" onChange={(e)=> {
                setPassword(e.target.value)}}/>
     
            <button onClick={handleSubmit} className="btn" type="submit">
              Submit
            </button>
          </form>
        </div>
      );
}
export default SignIn