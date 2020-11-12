import React, { createContext, useState } from "react";

export const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [jwt, setJWT] = useState(null);
  const [bioUser, setBioUser] = useState({});
  const [jobInfo, setJobInfo] = useState({});
  const [errorMessage, setErrorMessage] = useState("");
  const [favorites, setFavorites] = useState([]);
  return (
    <UserContext.Provider
      value={{
        jwt,
        setJWT,
        errorMessage,
        setErrorMessage,
        bioUser,
        setBioUser,
        jobInfo,
        setJobInfo,
        favorites,
        setFavorites,
      }}
    >
      {children}
    </UserContext.Provider>
  );
};
