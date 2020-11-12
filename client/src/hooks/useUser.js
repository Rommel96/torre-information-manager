import { useCallback, useContext } from "react";
import { UserContext } from "../context/UserContext";
import signupService from "../services/signupService";
import loginService from "../services/loginService";
import {
  findUserService,
  findJobService,
  getSavedJobsService,
} from "../services/handlerService";

export default function useUser() {
  const {
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
  } = useContext(UserContext);

  const signup = useCallback(
    async ({ name, email, password }) => {
      const data = await signupService({ name, email, password });
      console.log(data);
      if (data.status !== "ERROR") {
        setJWT(data.message);
        console.log(data.message);
        setErrorMessage("");
      } else {
        setErrorMessage(data.message);
      }
      console.log("jwt:", jwt);
    },
    [setJWT]
  );

  const login = useCallback(
    async ({ email, password }) => {
      const data = await loginService({ email, password });
      console.log(data);
      if (data.status !== "ERROR") {
        setJWT(data.message);
        console.log(data.message);
        setErrorMessage("");
      } else {
        setErrorMessage(data.message);
      }
      console.log("jwt:", jwt);
    },
    [setJWT]
  );

  const logout = useCallback(() => {
    setJWT(null);
  }, [setJWT]);

  const findUser = useCallback(async ({ username }) => {
    const data = await findUserService({ username, jwt });
    if (data) {
      setBioUser(data.person);
      setErrorMessage("");
    } else {
      setErrorMessage("User not Found");
    }
  }, []);

  const findJob = useCallback(async ({ job }) => {
    console.log(job);
    const data = await findJobService({ job });
    if (data) {
      let requirements = "";
      data.details.forEach((d) => {
        if (d.code === "requirements") {
          requirements = d.content;
          console.log(d);
          return;
        }
      });
      const stableOn = new Date(Date.parse(data.stableOn));
      const deadline = new Date(Date.parse(data.deadline));

      const parse = {
        id: data.id,
        objective: data.objective,
        organizations: data.organizations[0],
        requirements: requirements,
        stableOn: stableOn.toLocaleDateString("en-US"),
        deadline: deadline.toLocaleDateString("en-US"),
        status: data.status,
      };
      setJobInfo(parse);
      setErrorMessage("");
    } else {
      setErrorMessage("Jobs not Found");
    }
  }, []);

  const getSavedJobs = useCallback(async () => {
    const data = await getSavedJobsService({ jwt });
    if (data) {
      setFavorites(data.message);
      setErrorMessage("");
    } else {
      setErrorMessage("Jobs not Found");
    }
    console.log(favorites);
  }, []);

  return {
    isLogged: Boolean(jwt),
    errorMessage: String(errorMessage),
    setErrorMessage,
    bioUser: bioUser,
    jobInfo: jobInfo,
    favorites: Array(favorites),
    signup,
    login,
    logout,
    findUser,
    findJob,
    getSavedJobs,
    jwt: jwt,
  };
}
