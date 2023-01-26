import React, { Component } from "react";
import "./App.css";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import SignUp from "../pages/signup";
import Landing from "../pages/landing";
import SignIn  from "../pages/signin";
import SigninSuccess from "../pages/signinSuccess";
import SetNewGoal from "../pages/setNewGoal";
import SetGoalSuccess from "../pages/setGoalSuccess";
class App extends Component {
  async componentWillMount() {}
  async storeData(userName, password) {}

  constructor(props) {
    super(props);
    this.state = { };
  }
  render() {
    return (
      <Router>
        <Landing />
        <Routes>
          <Route path="signup" element={<SignUp />} />
          <Route path="signin" element={<SignIn />} />
          <Route path="signinSuccess" element={<SigninSuccess />} />
          <Route path="setNewGoal" element={<SetNewGoal />} />
          <Route path="setGoalSuccess" element={<SetGoalSuccess />} />
        </Routes>
      </Router>
    );
  }
}
export default App;
