import React from "react";
import { Link, useHistory } from "react-router-dom";
import Axios from "axios";

export default function JobItem(props) {
  const {
    isLogged,
    id,
    objective,
    organizations,
    requirements,
    stableOn,
    deadline,
    status,
    jwt,
  } = props;
  const history = useHistory();

  const addFavorite = async () => {
    console.log("jwt: ", jwt);
    const config = {
      headers: { Authorization: `Bearer ${jwt}` },
      validateStatus: null,
    };
    const instance = Axios.create({
      validateStatus: null,
    });
    const body = {
      jobId: id,
      objective,
      stableOn,
      deadline,
      status,
    };
    const res = await instance.post(
      "https://torre-information-manager.herokuapp.com/user/job",
      body,
      {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      }
    );
    if (res.status === 200) history.push("/favorites");
    return null;
  };

  return (
    <div className="mx-auto card w-75">
      <div className="card-body">
        <div className="row h-50">
          <img
            src={
              organizations
                ? organizations.picture
                : "https://picsum.photos/200/300.jpg"
            }
            alt={objective}
            style={{
              verticalAlign: "middle",
              width: 50,
              height: 50,
              borderRadius: "50%",
            }}
          />
          <h5 className="card-title align-self-center text-center my-2 mx-2 ">
            {objective}
          </h5>
        </div>
        <p className="card-text my-2">{requirements}</p>
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
          isLogged ? (
            <div>
              <a
                href={"https://torre.co/en/jobs/" + id}
                className="btn btn-success mx-1"
                target="_blank"
              >
                Apply
              </a>
              <a
                onClick={addFavorite}
                className="btn btn-info mx-1"
                target="_blank"
              >
                Saved
              </a>
            </div>
          ) : (
            <Link to="/login" className="btn btn-success">
              Login to Apply
            </Link>
          )
        ) : null}
      </div>
    </div>
  );
}
