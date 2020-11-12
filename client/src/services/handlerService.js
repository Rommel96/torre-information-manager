import Axios from "axios";
const ENDPOINT = "https://torre-information-manager.herokuapp.com/user/torre-";
const TEST = "https://torre-information-manager.herokuapp.com/search/";
const FAVORITES = "https://torre-information-manager.herokuapp.com/user/job";
export async function findUserService({ username, jwt }) {
  const config = {
    headers: { Authorization: `Bearer ${jwt}` },
    validateStatus: null,
  };
  const instance = Axios.create({
    validateStatus: null,
  });
  const res = await instance.get(ENDPOINT + "user/" + username, config);
  if (res.status === 200) return res.data;
  return null;
}

export async function findJobService({ job }) {
  const config = {
    validateStatus: null,
  };
  const res = await Axios.get(TEST + job, config);
  if (res.status === 200) return res.data;
  return null;
}

export async function getSavedJobsService({ jwt }) {
  const config = {
    headers: { Authorization: `Bearer ${jwt}` },
    validateStatus: null,
  };
  const instance = Axios.create({
    validateStatus: null,
  });
  const res = await instance.get(FAVORITES, config);
  if (res.status === 200) return res.data;
  return null;
}
