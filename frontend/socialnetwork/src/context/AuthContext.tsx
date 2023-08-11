import React, { createContext, useState } from "react";

type AuthUser = {
  userId: string;
  displayName: string;
  firstName: string;
  lastName: string;
  isLoggedIn: number;
  role: number;
  exp: number;
  authToken: string;
};

export type UserContextType = {
  user: AuthUser | null;
  setUser: React.Dispatch<React.SetStateAction<AuthUser | null>>;
};

type UserContextProviderType = {
  children: React.ReactNode;
};

export const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
});

export const UserContextProvider = ({ children }: UserContextProviderType) => {
  const [user, setUser] = useState<AuthUser | null>(null);
  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
};
