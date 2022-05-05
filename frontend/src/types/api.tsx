export type BackendUser = {
  pk: number;
  username: string;
  email: string;
  solvedTasks: TaskType[];
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
