import React, { useState } from "react";
import { Link } from "react-router-dom";
import useUser from "../hooks/useUser";
import JobItem from "./JobItem";

export default function Home() {
  const [search, setSearch] = useState("");
  const { isLogged, errorMessage, findJob, jobInfo, jwt } = useUser();

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log(errorMessage);
    if (search.length > 0) findJob({ job: search });
  };

  return (
    <div>
      {!isLogged ? (
        <div className="alert alert-info" role="alert">
          <strong className="mr-sm-2">Log in for more features!</strong>
          <Link to="/login">
            <button type="button" className="btn btn-success ">
              Log in for more features
            </button>
          </Link>
        </div>
      ) : null}
      <h3 className="mx-auto">Welcome</h3>
      <form
        onSubmit={handleSubmit}
        className="form-inline d-flex justify-content-center md-form form-sm mt-0 my-2"
      >
        <i className="fas fa-search" aria-hidden="true" />
        <input
          className="form-control form-control-sm ml-3 w-75 mr-sm-2"
          type="text"
          onChange={(e) => setSearch(e.target.value)}
          value={search}
          placeholder="Search your next Job"
          aria-label="Search"
        />
        <div className="form-group">
          <button className="btn btn-outline-success my-2 my-sm-0">
            Search
          </button>
        </div>
      </form>
      {errorMessage === "" ? (
        !jobInfo.objective ? null : (
          <JobItem
            isLogged={isLogged}
            jwt={jwt}
            id={jobInfo.id}
            requirements={jobInfo.requirements}
            objective={jobInfo.objective}
            organizations={jobInfo.organizations}
            stableOn={jobInfo.stableOn}
            deadline={jobInfo.deadline}
            status={jobInfo.status}
          ></JobItem>
        )
      ) : (
        <div className="alert alert-danger" role="alert">
          {errorMessage}
        </div>
      )}
    </div>
  );
}
