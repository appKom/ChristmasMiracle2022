import { atom } from "recoil";
import { User, TokenStateType } from "../types/api";

export const tokenState = atom<TokenStateType | null>({
  key: "token",
  default: null,
});

export const userState = atom<User | null>({
  key: "user",
  default: null,
});
