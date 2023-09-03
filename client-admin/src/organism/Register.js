import { useState, useRef } from 'react';
import Form from './Form';
import { toast } from 'react-toastify';

export default function Register() {
  const loading = useRef();
  const initialFormData = {
    username: '',
    email: '',
    password: '',
    phoneNumber: '',
    address: '',
  }
  const [formData, setFormData] = useState(initialFormData);

  const handleFormDataChange = (e) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
       [name]: value,
    });
  }

  const handleFormSubmit = async (e) => {
    e.preventDefault();
    loading.current = toast.info("Creating admin account, please wait...");
    try {
      const response = await fetch(`${process.env.REACT_APP_BASE_URL}/register`, {
        method: 'POST',
        headers: {
          "Content-Type": "application/json",
          access_token: localStorage.getItem('access_token'),
        },
        body: JSON.stringify(formData),
      })
      if (response.ok) {
        const result = await response.json();
        if (result) {
          toast.update(loading.current, {
            type: toast.TYPE.SUCCESS,
            render: 'new Admin account successfully created!',
          });
          setFormData(initialFormData)
        }
      } else {
        const error = await response.json();
        throw {error};
      };
    } catch (error) {
      toast.error(error.message);
    }
  }

  return (
    <div className="flex content-center items-center justify-center">
      <div className="w-full lg:w-6/12 md:w-8/12 px-4">
        <div className="relative flex flex-col min-w-0 break-words w-full mb-6 shadow-lg rounded-lg bg-white border-0">
          <div className="flex-auto px-4 lg:px-10 py-10">
            <div className="text-gray-500 text-center mb-4 font-bold">
              <small>Register New Admin</small>
            </div>
            <Form
              buttonText="Create account"
              formData={formData}
              onFormDataChange={handleFormDataChange} 
              onFormSubmit={handleFormSubmit} />
          </div>
        </div>
      </div>
    </div>
  );
}