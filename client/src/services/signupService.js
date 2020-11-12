import Axios from "axios";
const ENDPOINT = "https://torre-information-manager.herokuapp.com/auth/signup";

export default async function signup({ name, email, password }) {
  const instance = Axios.create({
    validateStatus: null,
  });
  const res = await instance.post(ENDPOINT, {
    name,
    email,
    password,
  });
  return res.data;
}
