import { AUTH_URL, post, VALIDATE_URL } from "./index";
import { TokenStateType } from "../types/api";

export type LoginArguments = {
  username: string;
  password: string;
};

export const loginUser = async ({
  username,
  password,
}: LoginArguments): Promise<TokenStateType | null> => {
  const response = await post({
    url: AUTH_URL,
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

export const validateToken = async (token: string): Promise<boolean> => {
  const response = await post({
    url: `${AUTH_URL}validate/`,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${VALIDATE_URL}`,
    },
  });
  if (response.status === 200) {
    return true;
  }
  return false;
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
