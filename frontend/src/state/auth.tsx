import { useEffect } from "react";
import { atom, useRecoilState } from "recoil";
import { BackendUser, TokenStateType } from "../types/api";
import { getUser } from "../api";

export const tokenState = atom<TokenStateType>({
  key: "token",
  default: {
    access: "",
    refresh: "",
  },
});

export const userState = atom<BackendUser | null>({
  key: "user",
  default: null,
});

export const useUser = (): BackendUser | null => {
  const [user, setUser] = useRecoilState(userState);
  const [token] = useRecoilState(tokenState);

  useEffect(() => {
    const fetchProfile = async () => {
      const user = await getUser(token.access);
      if (user) {
        setUser(user);
      }
    };

    if (!user) {
      fetchProfile();
    }
  }, [token, user, setUser]);
  return user;
};
