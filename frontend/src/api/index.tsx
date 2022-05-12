import { User, ScoreBoardUser, TaskType, TokenStateType } from "../types/api";

export const API_BASE =
  process.env.REACT_APP_API_BASE || "http://localhost:8000";

export const API_URL = `${API_BASE}/api/v1`;
export const AUTH_URL = `${API_URL}/auth`;

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
      Authorization: `Bearer ${token}`,
    },
  });

export const fetchProfile = async (
  token: TokenStateType | null
): Promise<User | null> => {
  if (!token) {
    return null;
  }

  const response = await authorizedGet({
    url: `${API_URL}/profile`,
    token: token.access,
  });
  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return null;
};

export const fetchTasks = async (): Promise<TaskType[]> => {
  const response = await get({
    url: `${API_URL}/tasks`,
  });

  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return [];
};

export const fetchTask = async (id: number): Promise<TaskType | null> => {
  const response = await get({
    url: `${API_URL}/tasks/${id}`,
  });

  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return null;
};

export const fetchScoreBoard = async (
  token: TokenStateType | null
): Promise<ScoreBoardUser[]> => {
  if (!token) {
    return [];
  }

  const response = await authorizedGet({
    url: `${API_URL}/scoreboard`,
    token: token.access,
  });

  if (response.status === 200) {
    const data = await response.json();
    return data;
  }
  return [];
};
