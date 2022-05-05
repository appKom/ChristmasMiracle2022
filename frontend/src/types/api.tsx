export type User = {
  ID: number;
  Username: string;
  Email: string;
  Points: number;
  Admin: boolean;
};

export type TaskType = {
  Title: string;
  Content: string;
  Points: number;
};

export type TokenStateType = {
  Access: string;
  Refresh: string;
};

export type ScoreBoardUser = {
  Username: string;
  Points: number;
};
