import React, { useState } from "react";
import useUser from "../hooks/useUser";
import UserItem from "./UserItem";

export default function FindUser() {
  const [username, setUsername] = useState("");
  const { findUser, bioUser, errorMessage } = useUser();

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(errorMessage);
    if (username.length > 0) findUser({ username });
  };

  return (
    <div>
      <h3 className="mx-auto">Find a registered user in Torre</h3>
      <form
        onSubmit={handleSubmit}
        className="form-inline d-flex justify-content-center md-form form-sm mt-0 my-2"
      >
        <i className="fas fa-search" aria-hidden="true" />
        <input
          className="form-control form-control-sm ml-3 w-75 mr-sm-2"
          type="text"
          onChange={(e) => setUsername(e.target.value)}
          value={username}
          placeholder="Username"
          aria-label="Search"
        />
        <div className="form-group">
          <button className="btn btn-outline-success my-2 my-sm-0">
            Search
          </button>
        </div>
      </form>
      {errorMessage === "" ? (
        !bioUser.publicId ? null : (
          <UserItem
            username={bioUser.publicId}
            picture={bioUser.picture}
            name={bioUser.name}
            professionalHeadline={bioUser.professionalHeadline}
            summaryOfBio={bioUser.summaryOfBio}
          ></UserItem>
        )
      ) : (
        <div className="alert alert-danger" role="alert">
          {errorMessage}
        </div>
      )}
    </div>
  );
}
