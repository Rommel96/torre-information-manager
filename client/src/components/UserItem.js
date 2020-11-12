import React from "react";

export default function UserItem(props) {
  const { picture, name, professionalHeadline, summaryOfBio, username } = props;
  return (
    <div className="card mx-auto" style={{ width: "18rem" }}>
      <img className="card-img-top" src={picture} alt={name} />
      <div className="card-body">
        <h5 className="card-title" style={{ color: "brown" }}>
          {name}
        </h5>
        <h6 className="card-title text-info">{professionalHeadline}</h6>
        <p className="card-text">{summaryOfBio}</p>
        <a
          href={"https://bio.torre.co/en/" + username}
          className="btn btn-primary"
          target="_blank"
        >
          See Profile
        </a>
      </div>
    </div>
  );
}
