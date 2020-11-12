import React from "react";
import { Link } from "react-router-dom";

export default function FavoriteItem(props) {
  const { job } = props;
  const { jobId, objective, stableOn, deadline, status } = job;
  return (
    <div className="mx-auto card w-75">
      <div className="card-body">
        <div className="row h-50">
          <h5 className="card-title align-self-center text-center my-2 mx-2 ">
            {objective}
          </h5>
        </div>
        <div className="row my-2 d-flex flex-row-reverse">
          <div className="btn btn-secondary mx-1">
            Status <span className="badge badge-light ml-1">{status}</span>
          </div>
          <div className="btn btn-secondary mx-1">
            Deadline <span className="badge badge-light ml-1">{deadline}</span>
          </div>
          <div className="btn btn-secondary mx-1">
            Available from{" "}
            <span className="badge badge-light ml-1">{stableOn}</span>
          </div>
        </div>
        {status !== "closed" ? (
          <a
            href={"https://torre.co/en/jobs/" + jobId}
            className="btn btn-success mx-1"
            target="_blank"
          >
            Apply
          </a>
        ) : null}
      </div>
    </div>
  );
}
