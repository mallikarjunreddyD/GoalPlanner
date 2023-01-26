import React, { Component } from "react";
import {Link, Routes, Route, useNavigate} from 'react-router-dom';


const SigninSuccess = () => {
    const navigate = useNavigate();
    const handleSetNewGoal = () =>{
        navigate('/setNewGoal')
    }
    const handleReviewGoal = () =>{
        navigate('/ReviewGoal')
    }
    return (
        <div>
              <button onClick={handleSetNewGoal} className="btn" type="submit">
              Set New Goal
            </button>
            <button onClick={handleReviewGoal} className="btn" type="submit">
              Review Goal
            </button>
        </div>
    )
}

export default SigninSuccess;