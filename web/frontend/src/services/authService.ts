import { useState, useEffect, createContext, useContext } from 'react';
import api from './api';

interface AuthContextProps {
  isAuthenticated: boolean;
  user: any;
  login: (email: string, password: string) => Promise<void>;
  register: (email: string, password: string) => Promise<void>;
  logout: () => void;
}

const AuthContext = createContext<AuthContextProps | undefined>(undefined);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<any>(null);

  useEffect(() => {
    // Fetch user data if logged in
  }, []);

  const login = async (email: string, password: string) => {
    const response = await api.post('/login', { email, password });
    setUser(response.data.user);
    setIsAuthenticated(true);
  };

  const register = async (email: string, password: string) => {
    await api.post('/register', { email, password });
    // Handle registration success
  };

  const logout = () => {
    setUser(null);
    setIsAuthenticated(false);
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, login, register, logout }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
};
