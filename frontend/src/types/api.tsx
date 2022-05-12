export type User = {
  ID: number;
  username: string;
  email: string;
  points: number;
  admin: boolean;
};

export type TaskType = {
  title: string;
  content: string;
  points: number;
  release_date: string;
};

export type TokenStateType = {
  access: string;
  refresh: string;
};

export type ScoreBoardUser = {
  username: string;
  points: number;
};
