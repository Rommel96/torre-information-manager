import React, { useEffect } from "react";
import FavoriteItem from "./FavoriteItem";
import useUser from "../hooks/useUser";

export default function Favorites() {
  const { getSavedJobs, favorites } = useUser();
  useEffect(() => {
    console.log("fav: ", favorites);
    getSavedJobs();
  }, []);

  return (
    <div>
      {favorites[0].map((item) => (
        <FavoriteItem key={item._id} job={item}></FavoriteItem>
      ))}
    </div>
  );
}
