import React from "react";
import { NavLink } from "react-router-dom";

const Landing = () => {
    return ( <> <nav>
    <NavLink to="/signup"> Sign Up</NavLink>
    <NavLink to="/signin"> Sign In</NavLink>
    </nav>
    </>)
}
export default Landing