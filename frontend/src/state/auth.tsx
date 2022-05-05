import { useEffect } from "react";
import { atom, useRecoilState } from "recoil";
import { BackendUser, TokenStateType } from "../types/api";
import { getUser } from "../api";

export const tokenState = atom<TokenStateType | null>({
  key: "token",
  default: null
});

export const userState = atom<BackendUser | null>({
  key: "user",
  default: null,
});

