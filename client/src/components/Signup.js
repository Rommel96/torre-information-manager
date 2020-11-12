import React, { useState, useEffect } from "react";
import { Link, useHistory } from "react-router-dom";
import { validateEmail } from "../utils/validation";
import useUser from "../hooks/useUser";

export default function Signup() {
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { signup, isLogged, errorMessage, setErrorMessage } = useUser();
  const history = useHistory();

  useEffect(() => {
    if (isLogged) history.push("/");
    console.log(isLogged);
  }, [isLogged]);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (!validateEmail(email)) setErrorMessage("Invalid Email. Try again!");
    else if (!name) setErrorMessage("All fields are required!");
    else if (password.length < 8)
      setErrorMessage("The password must be at least 8 characters!");
    else signup({ name, email, password });
  };

  return (
    <div className="container mt-5">
      <div className="row">
        <div className="col-md-8 offset-md-3">
          <h2>Create New Account </h2>
          {errorMessage === "" ? null : (
            <div className="alert alert-danger" role="alert">
              {errorMessage}
            </div>
          )}
          <form onSubmit={handleSubmit}>
            <div className="form-group">
              <label>Name</label>
              <input
                type="text"
                placeholder="Name"
                onChange={(e) => setName(e.target.value)}
                value={name}
                className="form-control"
              />
            </div>
            <div className="form-group">
              <label>Email</label>
              <input
                type="text"
                onChange={(e) => setEmail(e.target.value)}
                value={email}
                placeholder="Email"
                name="email"
                className="form-control"
              />
            </div>
            <div className="form-group">
              <label>Password</label>
              <input
                onChange={(e) => setPassword(e.target.value)}
                value={password}
                type="password"
                placeholder="Password"
                name="password"
                className="form-control"
              />
            </div>
            <div className="form-group">
              <button className="btn btn-primary">Sign Up</button>
            </div>
          </form>
          <div className="mx-auto" style={{ width: "200px" }}>
            <Link to="/login" onClick={() => setErrorMessage("")}>
              <button type="button" className="btn btn-success ">
                Sign in instead
              </button>
            </Link>
          </div>
        </div>
      </div>
    </div>
    //<div>ASasdasd</div>
  );
}
