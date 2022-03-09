import { BackendUser } from "../types/api";

export const API_BASE =
  process.env.REACT_APP_API_BASE || "http://localhost:8000";

export const API_URL = `${API_BASE}/api`;
export const AUTH_URL = `${API_BASE}/api/token/`;
export const REFRESH_URL = `${API_BASE}/api/token/refresh/`;
export const VALIDATE_URL = `${API_BASE}/api/token/validate/`;

export type AJAXArguments = {
  url: string;
  body?: Record<string, unknown> | string;
  headers?: HeadersInit;
  token?: string;
};

export const get = ({ url, body, headers }: AJAXArguments): Promise<Response> =>
  fetch(url, {
    method: "GET",
    body: JSON.stringify(body),
    headers: headers,
  });

export const authorizedGet = ({
  url,
  body,
  headers,
  token,
}: AJAXArguments): Promise<Response> => {
  return get({
    url,
    body,
    headers: {
      ...headers,
      Authorization: `Bearer ${token}`,
      "Content-Type": "application/json",
    },
  });
};

export const post = ({
  url,
  body,
  headers,
}: AJAXArguments): Promise<Response> =>
  fetch(url, {
    method: "POST",
    body: typeof body === "string" ? body : JSON.stringify(body),
    headers,
  });

export const authorizedPost = ({
  url,
  body,
  headers,
  token,
}: AJAXArguments): Promise<Response> =>
  post({
    url,
    body,
    headers: {
      ...headers,
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

export const getUser = async (
  token: string | null
): Promise<BackendUser | null> => {
  if (!token) {
    return null;
  }

  const response = await authorizedGet({
    url: `${API_URL}/profile/`,
    token,
  });
  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return null;
};
