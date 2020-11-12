import React from "react";
import { Link } from "react-router-dom";
import useUser from "../hooks/useUser";

export default function NavBar() {
  const { isLogged, logout } = useUser();
  //const isLogged = false;

  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light">
      <Link className="navbar-brand" to="/">
        Torre Manager
      </Link>
      <button
        className="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon" />
      </button>
      <div className="collapse navbar-collapse" id="navbarNav">
        <ul className="navbar-nav">
          <li className="nav-item">
            {isLogged ? (
              <Link className="nav-link" to="/find-user">
                Find User
              </Link>
            ) : null}
          </li>
          <li className="nav-item">
            {isLogged ? (
              <Link className="nav-link" to="/favorites">
                Favorites
              </Link>
            ) : null}
          </li>
          <li className="nav-item active">
            {!isLogged ? (
              <Link className="nav-link" to="/login">
                Login
              </Link>
            ) : (
              <Link className="nav-link" to="/" onClick={logout}>
                Logout
              </Link>
            )}
          </li>
        </ul>
      </div>
    </nav>
  );
}
