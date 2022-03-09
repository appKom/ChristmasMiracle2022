export type BackendUser = {
  pk: number;
  username: string;
  email: string;
  solvedTasks: TaskType[];
};

export type TaskType = {
  title: string;
  description: string;
  points: number;
};

export type TokenStateType = {
  access: string;
  refresh: string;
};
