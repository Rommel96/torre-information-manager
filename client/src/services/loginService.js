import Axios from "axios";
const ENDPOINT = "https://torre-information-manager.herokuapp.com/auth/login";

export default async function login({ email, password }) {
  const instance = Axios.create({
    validateStatus: null,
  });
  const res = await instance.post(ENDPOINT, {
    email,
    password,
  });
  return res.data;
}
