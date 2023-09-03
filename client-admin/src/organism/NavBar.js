import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function NavBar({ toggleMenu }) {
  const [activeMenu, setActiveMenu] = useState('dashboard')
  let navigate = useNavigate();
  let menuActive = 'w-full text-gray-800 flex items-center pl-6 p-2 my-2 transition-colors duration-200 justify-start border-l-4 border-purple-500';
  let menuNormal = 'w-full text-gray-400 flex items-center pl-6 p-2 my-2 transition-colors duration-200 justify-start hover:text-gray-800 border-l-4 border-transparent';

  let menuClass = 'h-screen lg:block shadow-lg relative w-80'
  toggleMenu ? menuClass.replace('hidden', '') : menuClass += ' hidden';

  return (
    <div>
      <div className={menuClass}>
        <div className="bg-white h-full dark:bg-gray-700">
          <div className="flex items-center justify-start pt-6 ml-8">
            <p className="font-bold dark:text-white text-xl">
              Buco Admin
            </p>
          </div>
          <nav className="mt-6">
            <div>
              <button
                onClick={() => {
                  setActiveMenu('dashboard');
                  navigate('/');
                }}
                className={ activeMenu === 'dashboard' ? menuActive : menuNormal}>
                <span className="text-left">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" />
                  </svg>
                </span>
                <span className="mx-2 text-sm font-normal"> Dashboard </span>
              </button>
              <button 
                onClick={() => {
                  setActiveMenu('categories');
                  navigate('/categories');
                }}
                className={ activeMenu === 'categories' ? menuActive : menuNormal}>
                <span className="text-left">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z" />
                  </svg>
                </span>
                <span className="mx-4 text-sm font-normal"> Categories </span>
              </button>
              <button 
                onClick={() => {
                  setActiveMenu('register');
                  navigate('/register-admin');
                }}
                className={ activeMenu === 'register' ? menuActive : menuNormal}>
                <span className="text-left">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M18 9v3m0 0v3m0-3h3m-3 0h-3m-2-5a4 4 0 11-8 0 4 4 0 018 0zM3 20a6 6 0 0112 0v1H3v-1z" />
                  </svg>
                </span>
                <span className="mx-4 text-sm font-normal"> Register Admin </span>
              </button>
              <button 
                onClick={() => {
                  localStorage.clear();
                  navigate('/login');
                }}
                className="w-full text-gray-400 flex items-center pl-6 p-2 my-2 transition-colors duration-200 justify-start hover:text-gray-800 border-l-4 border-transparent" href="#">
                <span className="text-left">
                  <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1" />
                  </svg>
                </span>
                <span className="mx-4 text-sm font-normal"> Logout </span>
              </button>
            </div>
          </nav>
        </div>
      </div>
    </div>
  )
}
