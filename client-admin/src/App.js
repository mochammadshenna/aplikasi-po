import { Routes, Route } from 'react-router-dom';
import Dashboard from './pages/Dashboard';
import CategoryList from './organism/CategoryList';
import ItemList from './organism/ItemList';
import Register from './organism/Register';
import Login from './pages/Login';
import { RequireAuth, Authenticated } from './organism/RequireAuth';
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

function App() {
  return (
    <div className="App">
      <ToastContainer />
      <Routes>
        <Route path="/" element={
          <RequireAuth>
            <Dashboard />
          </RequireAuth>
        }>
          <Route index element={<ItemList />} />
          <Route path="categories" element={<CategoryList />} />
          <Route path="register-admin" element={<Register />} />
        </Route>
        <Route path="login" element={<Authenticated><Login /></Authenticated>} />
      </Routes>
    </div>
  );
}

export default App;
