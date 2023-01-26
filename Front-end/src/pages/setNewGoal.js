import React, { Component } from "react";
import { Link, Routes, Route, useNavigate } from "react-router-dom";
import { useState } from "react";

const SetNewGoal = () => {
  const navigate = useNavigate();
  const [goalTitle, setGoalTitle] = useState("");
  const [goalPurpose1, setGoalPurpose1] = useState("");
  const [goalPurpose2, setGoalPurpose2] = useState("");
  const [specific, setIsSpecific] = useState("");
  const [measurable, setIsMeasurable] = useState("");
  const [achievable, setIsAchievable] = useState("");
  const [relevant, setIsRelevant] = useState("");
  const [timeBound, setIsTimeBound] = useState("");
  const [submitted, setSubmitted] = useState(false);
  const [error, setError] = useState(false);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (goalTitle === "" || goalPurpose1 === "" || goalPurpose2 === "") {
      setError(true);
    } else {
      setSubmitted(true);
      navigate("/setGoalSuccess");
      setError(false);
    }
  };

  return (
    <div>
      <form>
        {/* Labels and inputs for form data */}

        <label className="label">Goal Title</label>
        <input
          className="input"
          value={goalTitle}
          type="text"
          onChange={(e) => {
            setGoalTitle(e.target.value);
          }}
        />

        <label className="label">Why this Goal is important</label>
        <input
          className="input"
          value={goalPurpose1}
          type="text"
          onChange={(e) => {
            setGoalPurpose1(e.target.value);
          }}
        />

        <label className="label">What if this Goal is not set</label>
        <input
          className="input"
          value={goalPurpose2}
          type="text"
          onChange={(e) => {
            setGoalPurpose2(e.target.value);
          }}
        />
        <label>
          <input type="checkbox" name="specific" value={specific} onChange={(e) => {
            setIsSpecific(e.target.value)
          }}/>
          is Specific?
        </label>
        <label>
          <input type="checkbox" name="measurable" value={measurable} onChange={(e) => {
            setIsMeasurable(e.target.value)}}/>
          is Measurable?
        </label>
        <label>
          <input type="checkbox" name="achievable" value={achievable} onChange={(e) => {
            setIsAchievable(e.target.value)}}/>
          is Achievable?
        </label>
        <label>
          <input type="checkbox" name="relevant" value={relevant} onChange={(e) => {
            setIsRelevant(e.target.value)}}/>
          is Relevant?
        </label>
        <label>
          <input type="checkbox" name="timeBound" value={timeBound} onChange={(e) => {
            setIsTimeBound(e.target.value)}}/>
          is Time Bound?
        </label>

        <button onClick={handleSubmit} className="btn" type="submit">
          Submit
        </button>
      </form>
    </div>
  );
};

export default SetNewGoal;
