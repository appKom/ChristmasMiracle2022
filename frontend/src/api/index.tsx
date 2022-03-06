export const API_BASE =
  process.env.REACT_APP_API_BASE || "http://localhost:8000";

export const API_URL = `${API_BASE}api`;
export const AUTH_URL = `${API_BASE}auth`;

export type AJAXArguments = {
  url: string;
  body?: Record<string, unknown> | string;
  headers?: HeadersInit;
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
}: AJAXArguments): Promise<Response> =>
  get({
    url,
    body,
    headers: {
      ...headers,
      Authorization: `Token ${"token"}`,
      "Content-Type": "application/json",
    },
  });

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
}: AJAXArguments): Promise<Response> =>
  post({
    url,
    body,
    headers: {
      ...headers,
      "Content-Type": "application/json",
      Authorization: `Token ${"token"}`,
    },
  });
