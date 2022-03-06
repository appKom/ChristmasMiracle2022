import { AUTH_URL, post } from "./index";
import { LoginResponseType } from "../types/api";

export type LoginArguments = {
  username: string;
  password: string;
};

export const loginUser = async ({
  username,
  password,
}: LoginArguments): Promise<LoginResponseType | null> => {
  const response = await post({
    url: `${AUTH_URL}/login/`,
    body: { username, password },
    headers: {
      "Content-Type": "application/json",
    },
  });

  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return null;
};

export const logoutUser = async (): Promise<boolean> => {
  const response = await post({
    url: `${AUTH_URL}logout/`,
  });

  if (response.status === 200) {
    return true;
  }
  return false;
};
