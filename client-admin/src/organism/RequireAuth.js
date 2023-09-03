import { Navigate } from 'react-router';

export function RequireAuth({ children }) {
  if(!localStorage.getItem('access_token')) {
    return <Navigate to='/login' />;
  }
  return children;
}

export function Authenticated({ children }) {
  if(localStorage.getItem('access_token')) {
    return <Navigate to='/' />;
  }
  return children;
}
