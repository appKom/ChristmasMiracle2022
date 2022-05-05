import { AUTH_URL, post } from "./index";
import { TokenStateType } from "../types/api";

export type LoginArguments = {
  email: string;
  password: string;
};

export const loginUser = async ({
  email,
  password,
}: LoginArguments): Promise<TokenStateType | null> => {
  const response = await post({
    url: AUTH_URL + "/login",
    body: { email, password },
  });

  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return null;
};
